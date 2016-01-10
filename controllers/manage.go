package controllers

import (
	"automezzi/models"
	pk "automezzi/utilities/pbkdf2wrapper"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"html/template"
	"time"
    "strconv"
    "reflect"
)

/*var(
	
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
	Id_key string
}*/

func (this *MainController) Manage() {
// Only administrator can Manage accounts
	this.activeContent("manage/manage")

	//******** This page requires login
	sess := this.GetSession("automezzi")
	if sess == nil {
		this.Redirect("/home", 302)
		return
	} 
	flash := beego.NewFlash()
	m := sess.(map[string]interface{})
    fmt.Println(m["admin"])
    fmt.Println(reflect.ValueOf(m["admin"]).Type())  
	if m["admin"] == 3 {
        fmt.Printf("hai i diritti")
		//******** Read users from database
		o := orm.NewOrm()
		o.Using("default")
		var users []User
		num, err := o.Raw("SELECT id, first, last, email, id_key FROM auth_user",).QueryRows(&users)
		if err == nil {
			fmt.Println("user nums: ", num)
			for i := range users { 
				fmt.Println(users[i])
			}
			rows := "<tr><center><td>ID</td><td>NOME</td><td>COGNOME</td><td>EMAIL</td><td>MODIFICA</td></center></tr>"
			for i := range users {
				rows += fmt.Sprintf("<tr><td>%d</td>"+
					"<td>%s</td><td>%s</td><td>%s</td><td><center><a href='http://%s/manage/user/%s'>+</a></center></td></tr>", users[i].Id, users[i].First, users[i].Last, users[i].Email,"localhost:8080", users[i].Id_key)
			}
			this.Data["Rows"] = template.HTML(rows)		
		} else {
			flash.Notice("Errore, contattare l'amministratore del sito")
			flash.Store(&this.Controller)
			this.Redirect("/notice", 302)
		}
	} else {
			flash.Notice("Non hai i diritti per accedere a questa pagina")
			flash.Store(&this.Controller)
			this.Redirect("/notice", 302)	
	}
}

func (this *MainController) UsersManage() {
	this.activeContent("manage/user")

	//******** This page requires login
	sess := this.GetSession("automezzi")
	if sess != nil {
		m := sess.(map[string]interface{})
		flash := beego.NewFlash()
		if m["admin"] == 3 {
			var x pk.PasswordHash
	
			x.Hash = make([]byte, 32)
			x.Salt = make([]byte, 16)
		
			o := orm.NewOrm()
			o.Using("default")
			var id_key string
			id_key = this.Ctx.Input.Param(":parms")
			fmt.Println("key: ", id_key)
			user := models.AuthUser{Id_key: this.Ctx.Input.Param(":parms")}
			err := o.Read(&user, "Id_key")
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
			var userlvllist string
			switch user.Group {
				case 0:
				    userlvllist += fmt.Sprintf("<option value=\"0\" selected=\"selected\">Utente</option>"+
          "<option value=\"1\">Utente Speciale</option>"+
          "<option value=\"2\">Agente</option>"+
          "<option value=\"3\">Amministratore</option>")
				case 1:
				    userlvllist += fmt.Sprintf("<option value=\"0\">Utente</option>"+
          "<option value=\"1\" selected=\"selected\">Utente Speciale</option>"+
          "<option value=\"2\">Agente</option>"+
          "<option value=\"3\">Amministratore</option>")
				case 2:
				    userlvllist += fmt.Sprintf("<option value=\"0\">Utente</option>"+
          "<option value=\"1\">Utente Speciale</option>"+
          "<option value=\"2\" selected=\"selected\">Agente</option>"+
          "<option value=\"3\">Amministratore</option>")
				case 3:
				    userlvllist += fmt.Sprintf("<option value=\"0\">Utente</option>"+
          "<option value=\"1\">Utente Speciale</option>"+
          "<option value=\"2\">Agente</option>"+
          "<option value=\"3\" selected=\"selected\">Amministratore</option>")
				default:
				    panic("unrecognized escape character")
				}


	
			
    	
			// this deferred function ensures that the correct fields from the database are displayed
			defer func(this *MainController, user *models.AuthUser) {
				this.Data["UFirst"] = user.First
				this.Data["ULast"] = user.Last
				this.Data["UEmail"] = user.Email
				this.Data["Userlvllist"] = template.HTML(userlvllist)
			}(this, &user)
		
			if this.Ctx.Input.Method() == "POST" {
				first := this.GetString("first")
				last := this.GetString("last")
				email := this.GetString("email")
				//current := this.GetString("current")
				password := this.GetString("password")
				password2 := this.GetString("password2")
                userlvl := this.GetString("userlvl")
				valid := validation.Validation{}
				valid.Required(first, "first")
				valid.Email(email, "email")
				//valid.Required(current, "current")
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
		
				/******** Compare submitted password with database
				if !pk.MatchPassword(current, &x) {
					flash.Error("Password attuale errata")
					flash.Store(&this.Controller)
					return
				}*/
		
				//******** Save user info to database
				user.First = first
				user.Last = last
				user.Email = email
				user.Last_edit_date = time.Now()
                user.Group = ConvertInt(userlvl)	
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
		} else {
            //if user lvl isn't admin he is redirect
            this.Redirect("/home", 302)
			flash.Error("Non disponi dei privilegi necessari")
			flash.Store(&this.Controller)
            return
		}  
		
		
	} else {
		this.Redirect("/user/login/home", 302)
		return
	}
	
}

// this function convert string in int
func ConvertInt(s string) int {
    //convert string in int
    i, err := strconv.Atoi(s)
	if err != nil {
 	   panic(err)
	}
	return i
}
