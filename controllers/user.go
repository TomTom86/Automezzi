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
	"html/template"
	"strings"
	"time"
)

var(
	
	appcfg_GmailAccount string = beego.AppConfig.String("appcfg_GmailAccount")
	appcfg_GmailAccountPsw string = beego.AppConfig.String("appcfg_GmailAccountPsw")
	appcfg_MailHost string = beego.AppConfig.String("appcfg_MailHost")
	appcfg_MailHostPort, err = beego.AppConfig.Int("appcfg_MailHostPort")
)

type User struct {
    Id   int
    First string
	Last string
	Email string
	Group int
}

func (this *MainController) Login() {
	this.activeContent("user/login")

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
		err := o.Read(&user, "Email")
		//controll if the account is blocked
		if err == nil && user.Block_controll < 3 && user.Block_controll >= 0 {
			if user.Reg_key != "" {
				flash.Error("Account not verified")
				flash.Store(&this.Controller)
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
			// Reset block_controll if user login correctly 
			user.Block_controll = 0
			_, err := o.Update(&user)
			if err != nil {
				flash.Error("Internal error")
				flash.Store(&this.Controller)
				return
			}
			
		
		} else {
			if user.Block_controll > 2 {
				fmt.Println(user.Block_controll)
				flash.Error("Account bloccato, contattare l'amministratore del sito")
				flash.Store(&this.Controller)
				return
			} else {
			
			flash.Error("No such user/email")
			flash.Store(&this.Controller)
			return					
			}

		}

		//******** Compare submitted password with database and increment Block_controll
		if !pk.MatchPassword(password, &x) {
			flash.Error("Bad password")
			flash.Store(&this.Controller)
			user.Block_controll++
			_, err := o.Update(&user)
			if err == nil {
				return
			} else {
				flash.Error("Internal error")
				flash.Store(&this.Controller)
				return
			}


		}
	
		//******** Create session and go back to previous page
		var users []User
		num, err := o.Raw("SELECT auth_user.'group' FROM auth_user WHERE email = ?",email).QueryRows(&users)
		if err == nil {
			fmt.Println("Group: ", num)
			fmt.Println("user group: ", users[0].Group)
			m := make(map[string]interface{})
		  	m["first"] = user.First
		 	m["username"] = email
		 	m["timestamp"] = time.Now()
		 	if users[0].Group == 1 {
				 m["group"] = users[0].Group
			 } else {
				 m["group"] = 0
			 }
		 	this.SetSession("automezzi", m)
		 	this.Redirect("/"+back, 302)
		} else {
			flash.Error("Errore - Contattare l'amministratore del sito")
			flash.Store(&this.Controller)
			return
		}
		
	}
}

func (this *MainController) Logout() {
	this.activeContent("logout")
	this.DelSession("automezzi")
	this.Redirect("/home", 302)
}

type user1 struct {
	First    string `form:"first" valid:"Required"`
	Last     string `form:"last"`
	Email    string `form:"email" valid:"Email"`
	Password string `form:"password" valid:"MinSize(6)"`
	Confirm  string `form:"password2" valid:"Required"`
}

func (this *MainController) Register() {
	this.activeContent("user/register")

	if this.Ctx.Input.Method() == "POST" {
		flash := beego.NewFlash()
		u := user1{}
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

		user := models.AuthUser{First: u.First, Last: u.Last, Email: u.Email}

		// Convert password hash to string
		user.Password = hex.EncodeToString(h.Hash) + hex.EncodeToString(h.Salt)

		// Add user to database with new uuid and send verification email
		key := uuid.NewV4()
		user.Reg_key = key.String()
		_, err := o.Insert(&user)
		if err != nil {
			flash.Error(u.Email + " gia' registrata")
			flash.Store(&this.Controller)
			return
		}
		link := "http://localhost:8080/user/verif/"+ key.String()
		m.Email = u.Email
		m.Subject = "Verifica account portale automezzi"
		m.Body = "Per verificare l'account premere sul link: <a href=\""+link+"\">"+link+"</a><br><br>Grazie,<br>E' Cosi'"
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

type message1 struct {
  Email string
  Subject string
  Body string
}

func sendComunication(email message1) bool {
	msg := gomail.NewMessage()
	msg.SetHeader("From", appcfg_GmailAccount, "E' Cosi'")
	msg.SetHeader("To", email.Email)
	msg.SetHeader("Subject", email.Subject)
	msg.SetBody("text/html", email.Body)
	m := gomail.NewPlainDialer(appcfg_MailHost, appcfg_MailHostPort, appcfg_GmailAccount, appcfg_GmailAccountPsw)
	if err := m.DialAndSend(msg); err != nil {
		return false
	}
	return true
}


// da implementare meglio il sistema di verifica in quanto attualmente non funziona il get uuid.
// per ora disabilito la verifica dando per assodato che chi riceve l'account e' inserito volontariamente.
//di conseguenza in realta' si potrebbe disabilitare anche il moduro di registrazione per come e' stato predisposto o meglio
//renderlo accessibile soltanto agli utenti amministratori
func (this *MainController) Verify() {
	this.activeContent("user/verif")

	u := this.Ctx.Input.Param(":uuid")
	o := orm.NewOrm()
	o.Using("default")
	user := models.AuthUser{Reg_key: u}
	err := o.Read(&user, "Reg_key")
	if err == nil {
		this.Data["Verified"] = 1
		user.Reg_key = ""
		if _, err := o.Update(&user); err != nil {
			delete(this.Data, "Verified")
		}
	}
}

func (this *MainController) Profile() {
	this.activeContent("user/profile")

	//******** This page requires login
	sess := this.GetSession("automezzi")
	if sess == nil {
		this.Redirect("/user/login/home", 302)
		return
	}
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
	if err == nil {
		// scan in the password hash/salt
		if x.Hash, err = hex.DecodeString(user.Password[:64]); err != nil {
			fmt.Println("ERROR:", err)
		}
		if x.Salt, err = hex.DecodeString(user.Password[64:]); err != nil {
			fmt.Println("ERROR:", err)
		}
	} else {
		flash.Error("Internal error")
		flash.Store(&this.Controller)
		return
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

		_, err := o.Update(&user)
		if err == nil {
			flash.Notice("Profilo aggiornato")
			flash.Store(&this.Controller)
			m["username"] = email
		} else {
			flash.Error("Errore interno")
			flash.Store(&this.Controller)
			return
		}
	}
}

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
		if err == nil {
			// scan in the password hash/salt
			if x.Hash, err = hex.DecodeString(user.Password[:64]); err != nil {
				fmt.Println("ERROR:", err)
			}
			if x.Salt, err = hex.DecodeString(user.Password[64:]); err != nil {
				fmt.Println("ERROR:", err)
			}
		} else {
			flash.Error("Errore interno")
			flash.Store(&this.Controller)
			return
		}

		//******** Compare submitted password with database
		if !pk.MatchPassword(current, &x) {
			flash.Error("Password corrente sbagliata")
			flash.Store(&this.Controller)
			return
		}

		//******** Delete user record
		_, err = o.Delete(&user)
		if err == nil {
			flash.Notice("Il tuo account e' stato cancellato.")
			flash.Store(&this.Controller)
			this.DelSession("automezzi")
			this.Redirect("/notice", 302)
		} else {
			flash.Error("Internal error")
			flash.Store(&this.Controller)
			return
		}
	}
}

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
		link := "http://localhost:8080/user/reset/"+ u.String()
		m.Email = email
		m.Subject = "Richiesta di azzeramento password Portale E' Così"
		m.Body = "Per resettare la tua password, premi sul seguente link: <a href=\""+link+"\">"+link+"</a><br><br>Grazie,<br>E' Cosi'"
		sendComunication(m)
		flash.Notice("Ti abbiamo inviato un link per resettare la password. Controlla la tua email.")
		flash.Store(&this.Controller)
		this.Redirect("/notice", 302)
	}
}

func (this *MainController) Reset() {
	this.activeContent("user/reset")

	flash := beego.NewFlash()

	u := this.Ctx.Input.Param(":uuid")
	o := orm.NewOrm()
	o.Using("default")
	user := models.AuthUser{Reset_key: u}
	err := o.Read(&user, "Reset_key")
	if err == nil {
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

			user.Reset_key = ""
			if _, err := o.Update(&user); err != nil {
				flash.Error("Errore interno")
				flash.Store(&this.Controller)
				return
			}
			flash.Notice("Password aggiornata.")
			flash.Store(&this.Controller)
			this.Redirect("/notice", 302)
		}
	} else {
		flash.Notice("Chiave invalida.")
		flash.Store(&this.Controller)
		this.Redirect("/notice", 302)
	}
}


func (this *MainController) Manage() {
	this.activeContent("user/manage")

	//******** This page requires login
	sess := this.GetSession("automezzi")
	if sess == nil {
		this.Redirect("/home", 302)
		return
	}
	//m := sess.(map[string]interface{})

	//flash := beego.NewFlash()

	//******** Read password hash from database
	var x pk.PasswordHash

	x.Hash = make([]byte, 32)
	x.Salt = make([]byte, 16)

	o := orm.NewOrm()
	o.Using("default")
	var users []User
	num, err := o.Raw("SELECT id, first, last, email FROM auth_user",).QueryRows(&users)
	if err == nil {
		fmt.Println("user nums: ", num)
		for i := range users { 
			fmt.Println(users[i])
		}
		rows := "<tr><td>ID</td><td>NOME</td><td>COGNOME</td><td>EMAIL</td></tr>"
		for i := range users {
			rows += fmt.Sprintf("<tr><td>%d</td>"+
				"<td>%s</td><td>%s</td><td>%s</td></tr>", users[i].Id, users[i].First, users[i].Last, users[i].Email)
	}
	this.Data["Rows"] = template.HTML(rows)
		
	}

	if this.Ctx.Input.Method() == "POST" {
		fmt.Println(this)		
	}
}