package controllers

import (
	"html/template"
	"time"
    "strconv"
    "reflect"
	"encoding/hex"
	"fmt"
	"strings"
	
	"automezzi/models"
	pk "automezzi/utilities/pbkdf2wrapper"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"

)

//TODO: SE MODIFICO UN UTENTE AUTOMATICAMENTE ANDANDO SU PROFILE VIENE CARICATO L'ULTIMO UTENTE AGGIORANTO E NON IL MIO


func (this *MainController) setCompare(query string) (orm.QuerySeter, bool) {
	o := orm.NewOrm()
	qs := o.QueryTable("auth_user")
	if this.Ctx.Input.Method() == "POST" {
		f := compareform{}
		if err := this.ParseForm(&f); err != nil {
			fmt.Println("cannot parse form")
			return qs, false
		}
		valid := validation.Validation{}
		if b, _ := valid.Valid(&f); !b {
			this.Data["Errors"] = valid.ErrorsMap
			return qs, false
		}
		if len(f.Compareop) >= 5 && f.Compareop[:5] == "__not" {
			qs = qs.Exclude(f.Comparefield+f.Compareop[5:], f.Compareval)
		} else {
			qs = qs.Filter(f.Comparefield+f.Compareop, f.Compareval)
		}
		this.Data["query"] = f.Comparefield + f.Compareop + "," + f.Compareval
	} else {
		str := strings.Split(query, ",")
		i := strings.Index(str[0], "__")
		if len(str[0][i:]) >= 5 && str[0][i:i+5] == "__not" {
			qs = qs.Exclude(str[0][:i]+str[0][i+5:], str[1])
		} else {
			qs = qs.Filter(str[0], str[1])
		}
		this.Data["query"] = query
	}
	return qs, true
}

//TODO ordinare i nomi maiuscolo e minuscolo assieme
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
	if m["admin"] != 3 {
			flash.Notice("Non hai i diritti per accedere a questa pagina")
			flash.Store(&this.Controller)
			this.Redirect("/notice", 302)	
	}
	
    fmt.Printf("hai i diritti")
	
	//in caso di panic reindirizza alla home
	defer func(this *MainController) {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Index", r)
			this.Redirect("/home", 302)
		}
	}(this)
	//NON VA SENZA PARAMETRI
	//******** Read users from database
	if this.Ctx.Input.Param(":parms") == ""{
		o := orm.NewOrm()
		o.Using("default")
		var users []models.AuthUser
		
		o.QueryTable("auth_user")
		//num, err := o.Raw("SELECT id, first, last, email, id_key FROM auth_user",).QueryRows(&users)
		if err != nil {
			flash.Notice("Errore, contattare l'amministratore del sito")
			flash.Store(&this.Controller)
			this.Redirect("/notice", 302)		
		}
		
		//fmt.Println("user nums: ", num)
		for i := range users { 
			fmt.Println(users[i])
		}
		rows := "<tr><center><td>ID</td><td>NOME</td><td>COGNOME</td><td>EMAIL</td><td>MODIFICA</td></center></tr>"
		for i := range users {
			rows += fmt.Sprintf("<tr><td>%d</td>"+
				"<td>%s</td><td>%s</td><td>%s</td><td><center><a href='http://%s/manage/user/%s' class=\"user\"> </a></center></td></tr>", users[i].Id, users[i].First, users[i].Last, users[i].Email, appcfg_domainname, users[i].Id_key)
		}
		this.Data["Rows"] = template.HTML(rows)		
	}
	const pagesize = 10
	parms := this.Ctx.Input.Param(":parms")
	this.Data["parms"] = parms
	str := strings.Split(parms, "!")
	fmt.Println("parms is", str)
	order := str[0]
	off, _ := strconv.Atoi(str[1])
	offset := int64(off)
	if offset < 0 {
		offset = 0
	}
	query := str[2]

	var users []*models.AuthUser
	rows := ""

	qs, ok := this.setCompare(query)
	if !ok {
		fmt.Println("cannot set QuerySeter")
		o := orm.NewOrm()
		qs := o.QueryTable("auth_user")
		qs = qs.Filter("id__gte", 0)
		this.Data["query"] = "id__gte,0"
	}

	count, _ := qs.Count()
	this.Data["count"] = count
	if offset >= count {
		offset = 0
	}
	num, err := qs.Limit(pagesize, offset).OrderBy(order).All(&users)
	if err != nil {
		fmt.Println("Query table failed:", err)
	}
	//domainname := this.Data["domainname"]
	//rows := "<tr><center><td>ID</td><td>NOME</td><td>COGNOME</td><td>EMAIL</td><td>MODIFICA</td></center></tr>"
	for i := range users {
			rows += fmt.Sprintf("<tr><td>%d</td>"+
				"<td>%s</td><td>%s</td><td>%s</td><td><center><a href='http://%s/manage/user/%s' class=\"user\"> </a></center></td></tr>", users[i].Id, users[i].First, users[i].Last, users[i].Email, appcfg_domainname, users[i].Id_key)
	}
	this.Data["Rows"] = template.HTML(rows)

	this.Data["order"] = order
	this.Data["offset"] = offset
	this.Data["end"] = max(0, (count/pagesize)*pagesize)

	
	if num+offset < count {
		this.Data["next"] = num + offset
	}
	if offset-pagesize >= 0 {
		this.Data["prev"] = offset - pagesize
		this.Data["showprev"] = true
	} else if offset > 0 && offset < pagesize {
		this.Data["prev"] = 0
		this.Data["showprev"] = true
	}

	if count > pagesize {
		this.Data["ShowNav"] = true
	}
	this.Data["progress"] = float64(offset*100) / float64(max(count, 1))
}

func (this *MainController) UsersManage() {
	this.activeContent("manage/user")

	//******** This page requires login
	sess := this.GetSession("automezzi")
	if sess == nil {
		this.Redirect("/user/login/home", 302)
		return
	}
		
	m := sess.(map[string]interface{})
	flash := beego.NewFlash()
	if m["admin"] != 3 {
	    this.Redirect("/home", 302)
		flash.Error("Non disponi dei privilegi necessari")
		flash.Store(&this.Controller)
        return	
	}
	
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
	if err != nil {
		flash.Error("Internal error")
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


	userAPP := models.AuthApp{Id: user.Id}
	err = o.Read(&userAPP, "Id")
	if err != nil {
		flash.Error("Internal error")
		flash.Store(&this.Controller)
		return
	}

	// this deferred function ensures that the correct fields from the database are displayed
	defer func(this *MainController, user *models.AuthUser, userAPP *models.AuthApp) {
		//check the user lvl
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
			//check the app authorization
		var checkautomezzi, checkservizi string
		if userAPP.Automezzi {
			checkautomezzi += fmt.Sprintf("<input type=\"checkbox\" name=\"apps\" value=\"automezzi\" checked=\"checked\"> Automezzi<br>")
		} else {
			checkautomezzi += fmt.Sprintf("<input type=\"checkbox\" name=\"apps\" value=\"automezzi\"> Automezzi<br>")
		}
		if userAPP.Servizi {
			checkservizi += fmt.Sprintf("<input type=\"checkbox\" name=\"apps\" value=\"servizi\" checked=\"checked\"> Servizi<br>")
		} else {
			checkservizi += fmt.Sprintf("<input type=\"checkbox\" name=\"apps\" value=\"servizi\"> Servizi<br>")
		}
		
		
		this.Data["UFirst"] = user.First
		this.Data["ULast"] = user.Last
		this.Data["UEmail"] = user.Email
		this.Data["Userlvllist"] = template.HTML(userlvllist)
		this.Data["Checkautomezzi"] = template.HTML(checkautomezzi)
		this.Data["Checkservizi"] = template.HTML(checkservizi)
	}(this, &user, &userAPP)

	if this.Ctx.Input.Method() == "POST" {
		first := this.GetString("first")
		last := this.GetString("last")
		email := this.GetString("email")
		password := this.GetString("password")
		password2 := this.GetString("password2")
        userlvl := this.GetString("userlvl")
		apps := this.GetStrings("apps")

		valid := validation.Validation{}
		valid.Required(first, "first")
		valid.Email(email, "email")
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

	
		if stringInSlice("automezzi", apps) {
			userAPP.Automezzi = true
		} else{
			userAPP.Automezzi = false
		}
		if stringInSlice("servizi", apps) {
			userAPP.Servizi = true
		} else{
			userAPP.Servizi = false
		}
		
		_, err := o.Update(&user)
		if err != nil {
			flash.Error("Errore interno")
			flash.Store(&this.Controller)
			return
		}
		
		_, err = o.Update(&userAPP)
		if err != nil {
			flash.Error("Errore interno")
			flash.Store(&this.Controller)
			return
		}

		flash.Notice("Profilo aggiornato")
		flash.Store(&this.Controller)
		
	}		

}

// this funcion check if string is in slice
func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
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
