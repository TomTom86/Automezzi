package controllers

import (
	"automezzi/models"
	pk "automezzi/utilities/pbkdf2"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/go-gomail/gomail"
	"github.com/twinj/uuid"
)

var (
	appcfgdomainname        string = beego.AppConfig.String("appcfgdomainname")
	appcfgMailAccount       string = beego.AppConfig.String("appcfgMailAccount")
	appcfgMailAccountPsw    string = beego.AppConfig.String("appcfgMailAccountPsw")
	appcfgMailHost          string = beego.AppConfig.String("appcfgMailHost")
	appcfgMailHostPort, err        = beego.AppConfig.Int("appcfgMailHostPort")
)

//TODO la gestione dei permessi utente non è molto sicura, forse è meglio dividere i permessi in una tabella a parte
// BUG** LE MODIFICHE EFFETTUATE ALLE APPLICAZIONI E QUINDI AL MENU SONO VALIDE SOLO DOPO AVER RILOGGATO
//Login func manage User's login
func (c *MainController) Login() {
	c.activeContent("user/login")
	sess := c.GetSession("automezzi")
	if sess != nil {
		c.Redirect("/", 302)
		return
	}
	back := strings.Replace(c.Ctx.Input.Param(":back"), ">", "/", -1) // allow for deeper URL such as l1/l2/l3 represented by l1>l2>l3
	fmt.Println("back is", back)
	if c.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		email := strings.ToLower(c.GetString("email"))
		password := c.GetString("password")
		valid := validation.Validation{}
		valid.Email(email, "email")
		valid.Required(password, "password")
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
			}
			c.Data["Errors"] = errormap
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
			flash.Store(&c.Controller)
			return
		} else if err == orm.ErrMissPK {
			fmt.Println("Errore - Contattare l'amministratore del sito")
		}
		//check if the account is verified
		if user.IsApproved != true {
			flash.Error("Account non verificato")
			flash.Store(&c.Controller)
			return
		}
		//if the account is blocked
		if user.BlockControll > 2 || user.BlockControll < 0 {
			flash.Error("Account bloccato, contattare l'amministratore del sito")
			flash.Store(&c.Controller)
			return
		}
		// scan in the password hash/salt
		fmt.Println("Password to scan:", user.Password)
		if x.Hash, err = hex.DecodeString(user.Password[:64]); err != nil {
			fmt.Println("ERROR:", err)
		}
		if x.Salt, err = hex.DecodeString(user.Password[64:]); err != nil {
			fmt.Println("ERROR:", err)
		}
		fmt.Println("decoded password is", x)
		// Reset blockControll if user login correctly
		_, err := o.Update(&user)
		if err != nil {
			flash.Error("Internal error")
			flash.Store(&c.Controller)
			return
		}

		//******** Compare submitted password with database and increment Block_controll
		if !pk.MatchPassword(password, &x) {
			flash.Error("Bad password")
			flash.Store(&c.Controller)
			fmt.Println(user.BlockControll)
			user.BlockControll++
			fmt.Println(user.BlockControll)
			_, err := o.Update(&user)
			if err == nil {
				return
			}
			flash.Error("Internal error")
			flash.Store(&c.Controller)
			return
		}
		// blockControll = 0 because i logged
		user.BlockControll = 0

		//******** Create session and go back to previous page

		fmt.Println("user group: ", user.Group)
		m := make(map[string]interface{})
		m["first"] = user.First
		m["username"] = user.Email
		m["timestamp"] = time.Now()
		m["idkey"] = user.IDkey
		// check if userlvl is Administrator
		if user.Group == 3 {
			m["admin"] = user.Group
		} else {
			m["admin"] = 0
		}
		m["automezzi"] = user.AuthApp.Automezzi
		c.SetSession("automezzi", m)
		c.Redirect("/"+back, 302)

		//******** Update last login date
		user.LastLoginDate = time.Now()
		_, err1 := o.Update(&user)
		if err1 != nil {
			flash.Error("Errore interno")
			flash.Store(&c.Controller)
			return
		}
		fmt.Println("Aggiornato ultimo login")
	}
}

//Logout fun delete session and logout user
func (c *MainController) Logout() {
	c.activeContent("logout")
	c.DelSession("automezzi")
	c.Redirect("/", 302)
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

//Register func register user in the db
//TODO: migliorare errore validazione campo email
// BUG: se l'account esiste già crea comunque la tabella app
func (c *MainController) Register() {
	c.activeContent("user/register")

	if c.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		u := userForm{}
		m := message1{}
		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			return
		}
		u.Email = strings.ToLower(u.Email)
		c.Data["User"] = u
		valid := validation.Validation{}
		if b, _ := valid.Valid(&u); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return
		}
		if u.Password != u.Confirm {
			flash.Error("Le password non combaciano")
			flash.Store(&c.Controller)
			return
		}
		h := pk.HashPassword(u.Password)

		//******** Save user info to database
		o := orm.NewOrm()
		o.Using("default")

		//Create and set user and userApp models
		userAPP := models.AuthApp{Automezzi: false, Servizi: false}
		user := models.AuthUser{First: u.First, Last: u.Last, Email: u.Email, IsApproved: false, Group: 0, AuthApp: &userAPP}

		// Convert password hash to string
		user.Password = hex.EncodeToString(h.Hash) + hex.EncodeToString(h.Salt)

		// Add user to database with new uuid and send verification email
		key := uuid.NewV4()
		user.IDkey = key.String()

		//Check if e-mail is used yet
		var maps []orm.Params
		num, err := o.QueryTable("auth_user").Filter("Email", u.Email).Values(&maps, "Email")
		if err != nil {
			flash.Error("Errore interno - Contattare l'Amministratore del sito")
			flash.Store(&c.Controller)
			return
		}
		//fmt.Println(num)
		//fmt.Println(u.Email)
		if num == 0 {
			//if any account use email create account
			fmt.Println("Indirizzo Email non utilizzato")
			_, err = o.Insert(&userAPP)
			if err != nil {
				flash.Error("Errore autorizzazioni applicazioni")
				flash.Store(&c.Controller)
				return
			}
			_, err := o.Insert(&user)
			if err != nil {
				flash.Error(u.Email + " gia' registrata - Contattare l'amministratore")
				flash.Store(&c.Controller)
				return
			}

		} else {
			//else say "Account exist
			flash.Error(u.Email + " gia' registrata")
			flash.Store(&c.Controller)
			return
		}

		//Set verify message
		link := "http://" + appcfgdomainname + "/user/check/" + user.IDkey
		m.Email = u.Email
		m.Subject = "Verifica account portale automezzi"
		m.Body = "Per verificare l'account premere sul link: <a href=\"" + link + "\">" + link + "</a><br><br>Grazie,<br>E' Cosi'"
		if !sendComunication(m) {
			flash.Error("Impossibile inviare email di verifica")
			flash.Store(&c.Controller)
			return
		}
		flash.Notice("L'account e' stato creato. Ti abbiamo inviato una e-mail per verificare l'account.")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
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
	//fmt.Println(appcfgMailHost)
	//fmt.Println(appcfgMailHostPort)
	//fmt.Println(appcfgMailAccount)
	//fmt.Println(appcfgMailAccountPsw)
	msg := gomail.NewMessage()
	msg.SetHeader("From", appcfgMailAccount, "E' Cosi'")
	msg.SetHeader("To", email.Email)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody("text/html", email.Body)
	m := gomail.NewPlainDialer(appcfgMailHost, appcfgMailHostPort, appcfgMailAccount, appcfgMailAccountPsw)
	if err := m.DialAndSend(msg); err != nil {
		return false
	}
	return true
}

//Verify func verifing user by id key
//TODO tradurre messaggio di conferma verifica
func (c *MainController) Verify() {
	c.activeContent("user/check")
	flash := beego.NewFlash()
	u := c.Ctx.Input.Param(":uuid")
	o := orm.NewOrm()
	o.Using("default")
	user := models.AuthUser{IDkey: u}
	err := o.Read(&user, "IDkey")
	if err != nil {
		flash.Error("Chiave di verifica errata - Riprovare o contattare l'Amministratore")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	c.Data["Verified"] = 1
	user.IsApproved = true
	if _, err := o.Update(&user); err != nil {
		delete(c.Data, "Verified")
	}

}

// Profile func: User's can manage their account information
func (c *MainController) Profile() {
	c.activeContent("user/profile")

	//******** c page requires login
	sess := c.GetSession("automezzi")
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
			flash.Store(&c.Controller)
			return
		}

		// scan in the password hash/salt
		if x.Hash, err = hex.DecodeString(user.Password[:64]); err != nil {
			fmt.Println("ERROR:", err)
		}
		if x.Salt, err = hex.DecodeString(user.Password[64:]); err != nil {
			fmt.Println("ERROR:", err)
		}

		// c deferred function ensures that the correct fields from the database are displayed
		defer func(c *MainController, user *models.AuthUser) {
			c.Data["First"] = user.First
			c.Data["Last"] = user.Last
			c.Data["Email"] = user.Email
		}(c, &user)

		if c.Ctx.Input.Method() == "POST" {
			first := c.GetString("first")
			last := c.GetString("last")
			email := c.GetString("email")
			current := c.GetString("current")
			password := c.GetString("password")
			password2 := c.GetString("password2")
			valid := validation.Validation{}
			valid.Required(first, "first")
			valid.Email(email, "email")
			valid.Required(current, "current")
			if valid.HasErrors() {
				errormap := []string{}
				for _, err := range valid.Errors {
					errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
				}
				c.Data["Errors"] = errormap
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
					c.Data["Errors"] = errormap
					return
				}

				if password != password2 {
					flash.Error("Le password non corrispondono")
					flash.Store(&c.Controller)
					return
				}
				h := pk.HashPassword(password)

				// Convert password hash to string
				user.Password = hex.EncodeToString(h.Hash) + hex.EncodeToString(h.Salt)
			}

			//******** Compare submitted password with database
			if !pk.MatchPassword(current, &x) {
				flash.Error("Password attuale errata")
				flash.Store(&c.Controller)
				return
			}

			//******** Save user info to database
			user.First = first
			user.Last = last
			user.Email = email
			user.LastEditDate = time.Now()

			_, err := o.Update(&user)
			if err != nil {
				flash.Error("Errore interno")
				flash.Store(&c.Controller)
				return
			}

			flash.Notice("Profilo aggiornato")
			flash.Store(&c.Controller)
			//update sessin email
			m["username"] = email
		}
	} else {
		//if user isn't logged redirect in the ompage
		c.Redirect("/user/login/", 302)
		return
	}

}

//Remove func delete user from DB
func (c *MainController) Remove() {
	c.activeContent("user/remove")

	//******** c page requires login
	sess := c.GetSession("automezzi")
	if sess == nil {
		c.Redirect("/user/login/", 302)
		return
	}
	m := sess.(map[string]interface{})

	if c.Ctx.Input.Method() == "POST" {
		current := c.GetString("current")
		valid := validation.Validation{}
		valid.Required(current, "current")
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
			}
			c.Data["Errors"] = errormap
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
			flash.Store(&c.Controller)
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
			flash.Store(&c.Controller)
			return
		}

		//******** Delete user record
		_, err = o.Delete(&user)
		if err != nil {
			flash.Error("Errore Interno - Contattare l'Amministratore del sito")
			flash.Store(&c.Controller)
			return
		}
		flash.Notice("Il tuo account e' stato cancellato.")
		flash.Store(&c.Controller)
		c.DelSession("automezzi")
		c.Redirect("/notice", 302)

	}
}

//Forgot func help user to restore password if they forgot it
func (c *MainController) Forgot() {
	c.activeContent("user/forgot")

	if c.Ctx.Input.Method() == "POST" {
		email := c.GetString("email")
		valid := validation.Validation{}
		valid.Email(email, "email")
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
			}
			c.Data["Errors"] = errormap
			return
		}

		flash := beego.NewFlash()

		o := orm.NewOrm()
		o.Using("default")
		user := models.AuthUser{Email: email}
		err := o.Read(&user, "Email")
		if err != nil {
			flash.Error("Non esiste un utente con questo indirizzo e-mail")
			flash.Store(&c.Controller)
			return
		}

		u := uuid.NewV4()
		user.ResetKey = u.String()
		_, err = o.Update(&user)
		if err != nil {
			flash.Error("Errore interno")
			flash.Store(&c.Controller)
			return
		}

		m := message1{}
		link := "http://" + appcfgdomainname + "/user/reset/" + u.String()
		m.Email = email
		m.Subject = "Richiesta di azzeramento password Portale E' Così"
		m.Body = "Per resettare la tua password, premi sul seguente link: <a href=\"" + link + "\">" + link + "</a><br><br>Grazie,<br>E' Cosi'"
		sendComunication(m)
		flash.Notice("Ti abbiamo inviato un link per resettare la password. Controlla la tua email.")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
}

//Reset func reset password if user forgot login credentials
func (c *MainController) Reset() {
	c.activeContent("user/reset")

	flash := beego.NewFlash()

	u := c.Ctx.Input.Param(":uuid")
	o := orm.NewOrm()
	o.Using("default")
	user := models.AuthUser{ResetKey: u}
	err := o.Read(&user, "ResetKey")
	if err != nil {
		flash.Error("Chiave invalida.")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	if c.Ctx.Input.Method() == "POST" {
		password := c.GetString("password")
		password2 := c.GetString("password2")
		valid := validation.Validation{}
		valid.MinSize(password, 6, "password")
		valid.Required(password2, "password2")
		if valid.HasErrors() {
			errormap := []string{}
			for _, err := range valid.Errors {
				errormap = append(errormap, "Validation failed on "+err.Key+": "+err.Message+"\n")
			}
			c.Data["Errors"] = errormap
			return
		}

		if password != password2 {
			flash.Error("Le password non corrispondono")
			flash.Store(&c.Controller)
			return
		}
		h := pk.HashPassword(password)

		// Convert password hash to string
		user.Password = hex.EncodeToString(h.Hash) + hex.EncodeToString(h.Salt)
		// Reset ResetKey flag and update lastEditDate
		user.ResetKey = ""
		user.LastEditDate = time.Now()
		if _, err := o.Update(&user); err != nil {
			flash.Error("Errore interno")
			flash.Store(&c.Controller)
			return
		}
		flash.Notice("Password aggiornata.")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
}
