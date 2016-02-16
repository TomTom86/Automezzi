package controllers

import (
	"automezzi/models"
	pk "automezzi/utilities/pbkdf2wrapper"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/go-gomail/gomail"
	"github.com/twinj/uuid"
	"strings"
	"time"
)

var (
	appcfg_domainname        string = beego.AppConfig.String("appcfg_domainname")
	appcfg_MailAccount       string = beego.AppConfig.String("appcfg_MailAccount")
	appcfg_MailAccountPsw    string = beego.AppConfig.String("appcfg_MailAccountPsw")
	appcfg_MailHost          string = beego.AppConfig.String("appcfg_MailHost")
	appcfg_MailHostPort, err        = beego.AppConfig.Int("appcfg_MailHostPort")
)

/*
type User struct {
    Id   int32
    First string
	Last string
	Email string
	Group int
	Id_key string
	Is_approved bool
	Last_login_date	time.Time
	Last_edit_date	time.Time
}
*/
//TODO la gestione dei permessi utente non è molto sicura, forse è meglio dividere i permessi in una tabella a parte

//Login func manage User's login
func (this *MainController) Login() {
	this.activeContent("user/login")
	sess := this.GetSession("automezzi")
	if sess != nil {
		this.Redirect("/home", 302)
		return
	}
	back := strings.Replace(this.Ctx.Input.Param(":back"), ">", "/", -1) // allow for deeper URL such as l1/l2/l3 represented by l1>l2>l3
	fmt.Println("back is", back)
	if this.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		email := this.GetString("email")
		password := this.GetString("password")
		valid := validation.Validation{}
		valid.Email(email, "email")
		valid.Required(password, "password")
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
			}
			this.Data["Errors"] = errormap
			return
		}
		fmt.Println("Authorization is", email, ":", password)

		//******** Read password hash from database
		var x pk.PasswordHash

		x.Hash = make([]byte, 32)
		x.Salt = make([]byte, 16)

		o := orm.NewOrm()
		o.Using("default")
		user := models.AuthUser{Email: email}
		err = o.QueryTable("auth_user").Filter("Email", email).RelatedSel().One(&user)
		if err == orm.ErrNoRows {
			//check if the account exist
			flash.Error("Account non esiste")
			flash.Store(&this.Controller)
			return
		} else if err == orm.ErrMissPK {
			fmt.Println("Errore - Contattare l'amministratore del sito")
		}
		//check if the account is verified
		if user.Is_approved != true {
			flash.Error("Account non verificato")
			flash.Store(&this.Controller)
			return
		}
		//if the account is blocked
		if user.Block_controll > 2 || user.Block_controll < 0 {
			flash.Error("Account bloccato, contattare l'amministratore del sito")
			flash.Store(&this.Controller)
			return
		} else {
			// scan in the password hash/salt
			fmt.Println("Password to scan:", user.Password)
			if x.Hash, err = hex.DecodeString(user.Password[:64]); err != nil {
				fmt.Println("ERROR:", err)
			}
			if x.Salt, err = hex.DecodeString(user.Password[64:]); err != nil {
				fmt.Println("ERROR:", err)
			}
			fmt.Println("decoded password is", x)
			// Reset block_controll if user login correctly
			_, err := o.Update(&user)
			if err != nil {
				flash.Error("Internal error")
				flash.Store(&this.Controller)
				return
			}
		}

		//******** Compare submitted password with database and increment Block_controll
		if !pk.MatchPassword(password, &x) {
			flash.Error("Bad password")
			flash.Store(&this.Controller)
			fmt.Println(user.Block_controll)
			user.Block_controll++
			fmt.Println(user.Block_controll)
			_, err := o.Update(&user)
			if err == nil {
				return
			} else {
				flash.Error("Internal error")
				flash.Store(&this.Controller)
				return
			}
		}
		// block_controll = 0 because i logged
		user.Block_controll = 0

		//******** Create session and go back to previous page

		fmt.Println("user group: ", user.Group)
		m := make(map[string]interface{})
		m["first"] = user.First
		m["username"] = user.Email
		m["timestamp"] = time.Now()
		m["id_key"] = user.Id_key
		// check if userlvl is Administrator
		if user.Group == 3 {
			m["admin"] = user.Group
		} else {
			m["admin"] = 0
		}
		m["automezzi"] = user.AuthApp.Automezzi
		this.SetSession("automezzi", m)
		this.Redirect("/"+back, 302)

		//******** Update last login date
		user.Last_login_date = time.Now()
		_, err1 := o.Update(&user)
		if err1 != nil {
			flash.Error("Errore interno")
			flash.Store(&this.Controller)
			return
		}
		fmt.Println("Aggiornato ultimo login")
	}
}

//Logout fun delete session and logout user
func (this *MainController) Logout() {
	this.activeContent("logout")
	this.DelSession("automezzi")
	this.Redirect("/home", 302)
}

//Type userForm is for get information by form
type userForm struct {
	First     string `form:"first" valid:"Required"`
	Last      string `form:"last"`
	Email     string `form:"email" valid:"Email"`
	Password  string `form:"password" valid:"MinSize(6)"`
	Confirm   string `form:"password2" valid:"Required"`
	Automezzi bool
	Servizi   bool
}

//TODO: migliorare errore validazione campo email
//Registration func register user in the db
func (this *MainController) Register() {
	this.activeContent("user/register")

	if this.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		u := userForm{}
		m := message1{}
		if err := this.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			return
		}
		this.Data["User"] = u
		valid := validation.Validation{}
		if b, _ := valid.Valid(&u); !b {
			this.Data["Errors"] = valid.ErrorsMap
			return
		}
		if u.Password != u.Confirm {
			flash.Error("Le password non combaciano")
			flash.Store(&this.Controller)
			return
		}
		h := pk.HashPassword(u.Password)

		//******** Save user info to database
		o := orm.NewOrm()
		o.Using("default")

		//Create and set user and userApp models
		userAPP := models.AuthApp{Automezzi: false, Servizi: false}
		user := models.AuthUser{First: u.First, Last: u.Last, Email: u.Email, Is_approved: false, Group: 0, AuthApp: &userAPP}

		// Convert password hash to string
		user.Password = hex.EncodeToString(h.Hash) + hex.EncodeToString(h.Salt)

		// Add user to database with new uuid and send verification email
		key := uuid.NewV4()
		user.Id_key = key.String()

		_, err = o.Insert(&userAPP)
		if err != nil {
			flash.Error("Errore autorizzazioni applicazioni")
			flash.Store(&this.Controller)
			return
		}

		_, err := o.Insert(&user)
		if err != nil {
			flash.Error(u.Email + " gia' registrata")
			flash.Store(&this.Controller)
			return
		}

		//Set verify message
		link := "http://" + appcfg_domainname + "/user/check/" + user.Id_key
		m.Email = u.Email
		m.Subject = "Verifica account portale automezzi"
		m.Body = "Per verificare l'account premere sul link: <a href=\"" + link + "\">" + link + "</a><br><br>Grazie,<br>E' Cosi'"
		if !sendComunication(m) {
			flash.Error("Impossibile inviare email di verifica")
			flash.Store(&this.Controller)
			return
		}
		flash.Notice("L'account e' stato creato. Ti abbiamo inviato una e-mail per verificare l'account.")
		flash.Store(&this.Controller)
		this.Redirect("/notice", 302)
	}
}

//  Type used for send email. It contain mail adress, subject and Body
type message1 struct {
	Email   string
	Subject string
	Body    string
}

//S endComunication func get smtp setting from app.conf and send e-mail
func sendComunication(email message1) bool {
	fmt.Println(appcfg_MailHost)
	fmt.Println(appcfg_MailHostPort)
	fmt.Println(appcfg_MailAccount)
	fmt.Println(appcfg_MailAccountPsw)
	msg := gomail.NewMessage()
	msg.SetHeader("From", appcfg_MailAccount, "E' Cosi'")
	msg.SetHeader("To", email.Email)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody("text/html", email.Body)
	m := gomail.NewPlainDialer(appcfg_MailHost, appcfg_MailHostPort, appcfg_MailAccount, appcfg_MailAccountPsw)
	if err := m.DialAndSend(msg); err != nil {
		return false
	}
	return true
}

//TODO tradurre messaggio di conferma verifica
//Verify func verifing user by id key
func (this *MainController) Verify() {
	this.activeContent("user/check")
	flash := beego.NewFlash()
	u := this.Ctx.Input.Param(":uuid")
	o := orm.NewOrm()
	o.Using("default")
	user := models.AuthUser{Id_key: u}
	err := o.Read(&user, "Id_key")
	if err != nil {
		flash.Notice("Chiave di verifica errata - Riprovare o contattare l'Amministratore")
		flash.Store(&this.Controller)
		this.Redirect("/notice", 302)
	}
	this.Data["Verified"] = 1
	user.Is_approved = true
	if _, err := o.Update(&user); err != nil {
		delete(this.Data, "Verified")
	}

}

// Profile func: User's can manage their account information
func (this *MainController) Profile() {
	this.activeContent("user/profile")

	//******** This page requires login
	sess := this.GetSession("automezzi")
	if sess != nil {
		m := sess.(map[string]interface{})
		flash := beego.NewFlash()

		//******** Read password hash from database
		var x pk.PasswordHash

		x.Hash = make([]byte, 32)
		x.Salt = make([]byte, 16)

		o := orm.NewOrm()
		o.Using("default")
		user := models.AuthUser{Email: m["username"].(string)}
		err := o.Read(&user, "Email")
		if err != nil {
			flash.Error("Errore interno - Contattare l'Amministratore del sito")
			flash.Store(&this.Controller)
			return
		}

		// scan in the password hash/salt
		if x.Hash, err = hex.DecodeString(user.Password[:64]); err != nil {
			fmt.Println("ERROR:", err)
		}
		if x.Salt, err = hex.DecodeString(user.Password[64:]); err != nil {
			fmt.Println("ERROR:", err)
		}

		// this deferred function ensures that the correct fields from the database are displayed
		defer func(this *MainController, user *models.AuthUser) {
			this.Data["First"] = user.First
			this.Data["Last"] = user.Last
			this.Data["Email"] = user.Email
		}(this, &user)

		if this.Ctx.Input.Method() == "POST" {
			first := this.GetString("first")
			last := this.GetString("last")
			email := this.GetString("email")
			current := this.GetString("current")
			password := this.GetString("password")
			password2 := this.GetString("password2")
			valid := validation.Validation{}
			valid.Required(first, "first")
			valid.Email(email, "email")
			valid.Required(current, "current")
			if valid.HasErrors() {
				errormap := []string{}
				for _, err := range valid.Errors {
					errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
				}
				this.Data["Errors"] = errormap
				return
			}

			if password != "" {
				valid.MinSize(password, 6, "password")
				valid.Required(password2, "password2")
				if valid.HasErrors() {
					errormap := []string{}
					for _, err := range valid.Errors {
						errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
					}
					this.Data["Errors"] = errormap
					return
				}

				if password != password2 {
					flash.Error("Le password non corrispondono")
					flash.Store(&this.Controller)
					return
				}
				h := pk.HashPassword(password)

				// Convert password hash to string
				user.Password = hex.EncodeToString(h.Hash) + hex.EncodeToString(h.Salt)
			}

			//******** Compare submitted password with database
			if !pk.MatchPassword(current, &x) {
				flash.Error("Password attuale errata")
				flash.Store(&this.Controller)
				return
			}

			//******** Save user info to database
			user.First = first
			user.Last = last
			user.Email = email
			user.Last_edit_date = time.Now()

			_, err := o.Update(&user)
			if err != nil {
				flash.Error("Errore interno")
				flash.Store(&this.Controller)
				return
			}

			flash.Notice("Profilo aggiornato")
			flash.Store(&this.Controller)
			//update sessin email
			m["username"] = email
		}
	} else {
		//if user isn't logged redirect in the ompage
		this.Redirect("/user/login/home", 302)
		return
	}

}

//Remove func delete user from DB
func (this *MainController) Remove() {
	this.activeContent("user/remove")

	//******** This page requires login
	sess := this.GetSession("automezzi")
	if sess == nil {
		this.Redirect("/user/login/home", 302)
		return
	}
	m := sess.(map[string]interface{})

	if this.Ctx.Input.Method() == "POST" {
		current := this.GetString("current")
		valid := validation.Validation{}
		valid.Required(current, "current")
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
			}
			this.Data["Errors"] = errormap
			return
		}

		flash := beego.NewFlash()

		//******** Read password hash from database
		var x pk.PasswordHash

		x.Hash = make([]byte, 32)
		x.Salt = make([]byte, 16)

		o := orm.NewOrm()
		o.Using("default")
		user := models.AuthUser{Email: m["username"].(string)}
		err := o.Read(&user, "Email")
		if err != nil {
			flash.Error("Errore interno")
			flash.Store(&this.Controller)
			return
		}
		// scan in the password hash/salt
		if x.Hash, err = hex.DecodeString(user.Password[:64]); err != nil {
			fmt.Println("ERROR:", err)
		}
		if x.Salt, err = hex.DecodeString(user.Password[64:]); err != nil {
			fmt.Println("ERROR:", err)
		}

		//******** Compare submitted password with database
		if !pk.MatchPassword(current, &x) {
			flash.Error("Password corrente sbagliata")
			flash.Store(&this.Controller)
			return
		}

		//******** Delete user record
		_, err = o.Delete(&user)
		if err != nil {
			flash.Error("Errore Interno - Contattare l'Amministratore del sito")
			flash.Store(&this.Controller)
			return
		}
		flash.Notice("Il tuo account e' stato cancellato.")
		flash.Store(&this.Controller)
		this.DelSession("automezzi")
		this.Redirect("/notice", 302)

	}
}

//Forgot func help user to restore password if they forgot it
func (this *MainController) Forgot() {
	this.activeContent("user/forgot")

	if this.Ctx.Input.Method() == "POST" {
		email := this.GetString("email")
		valid := validation.Validation{}
		valid.Email(email, "email")
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
			}
			this.Data["Errors"] = errormap
			return
		}

		flash := beego.NewFlash()

		o := orm.NewOrm()
		o.Using("default")
		user := models.AuthUser{Email: email}
		err := o.Read(&user, "Email")
		if err != nil {
			flash.Error("Non esiste un utente con questo indirizzo e-mail")
			flash.Store(&this.Controller)
			return
		}

		u := uuid.NewV4()
		user.Reset_key = u.String()
		_, err = o.Update(&user)
		if err != nil {
			flash.Error("Errore interno")
			flash.Store(&this.Controller)
			return
		}

		m := message1{}
		link := "http://" + appcfg_domainname + "/user/reset/" + u.String()
		m.Email = email
		m.Subject = "Richiesta di azzeramento password Portale E' Così"
		m.Body = "Per resettare la tua password, premi sul seguente link: <a href=\"" + link + "\">" + link + "</a><br><br>Grazie,<br>E' Cosi'"
		sendComunication(m)
		flash.Notice("Ti abbiamo inviato un link per resettare la password. Controlla la tua email.")
		flash.Store(&this.Controller)
		this.Redirect("/notice", 302)
	}
}

//Reset func reset password if user forgot login credentials
func (this *MainController) Reset() {
	this.activeContent("user/reset")

	flash := beego.NewFlash()

	u := this.Ctx.Input.Param(":uuid")
	o := orm.NewOrm()
	o.Using("default")
	user := models.AuthUser{Reset_key: u}
	err := o.Read(&user, "Reset_key")
	if err != nil {
		flash.Notice("Chiave invalida.")
		flash.Store(&this.Controller)
		this.Redirect("/notice", 302)
	}
	if this.Ctx.Input.Method() == "POST" {
		password := this.GetString("password")
		password2 := this.GetString("password2")
		valid := validation.Validation{}
		valid.MinSize(password, 6, "password")
		valid.Required(password2, "password2")
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
			}
			this.Data["Errors"] = errormap
			return
		}

		if password != password2 {
			flash.Error("Le password non corrispondono")
			flash.Store(&this.Controller)
			return
		}
		h := pk.HashPassword(password)

		// Convert password hash to string
		user.Password = hex.EncodeToString(h.Hash) + hex.EncodeToString(h.Salt)
		// Reset Reset_key flag and update last_edit_date
		user.Reset_key = ""
		user.Last_edit_date = time.Now()
		if _, err := o.Update(&user); err != nil {
			flash.Error("Errore interno")
			flash.Store(&this.Controller)
			return
		}
		flash.Notice("Password aggiornata.")
		flash.Store(&this.Controller)
		this.Redirect("/notice", 302)
	}
}
