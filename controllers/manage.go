package controllers

import (
	"encoding/hex"
	"fmt"
	"html/template"
	"reflect"
	"strconv"
	"strings"
	"time"

	pk "acme3/utilities/pbkdf2"
	"automezzi/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

func (c *MainController) setCompare(query string) (orm.QuerySeter, bool) {
	o := orm.NewOrm()
	qs := o.QueryTable("auth_user")
	if c.Ctx.Input.Method() == "POST" {
		f := compareform{}
		if err := c.ParseForm(&f); err != nil {
			fmt.Println("cannot parse form")
			return qs, false
		}
		valid := validation.Validation{}
		if b, _ := valid.Valid(&f); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return qs, false
		}
		if len(f.Compareop) >= 5 && f.Compareop[:5] == "__not" {
			qs = qs.Exclude(f.Comparefield+f.Compareop[5:], f.Compareval)
		} else {
			qs = qs.Filter(f.Comparefield+f.Compareop, f.Compareval)
		}
		c.Data["query"] = f.Comparefield + f.Compareop + "," + f.Compareval
	} else {
		str := strings.Split(query, ",")
		i := strings.Index(str[0], "__")
		if len(str[0][i:]) >= 5 && str[0][i:i+5] == "__not" {
			qs = qs.Exclude(str[0][:i]+str[0][i+5:], str[1])
		} else {
			qs = qs.Filter(str[0], str[1])
		}
		c.Data["query"] = query
	}
	return qs, true
}

//Manage help administrator to manage all accounts
//TODO ordinare i nomi maiuscolo e minuscolo assieme
func (c *MainController) Manage() {
	// Only administrator can Manage accounts
	c.activeContent("manage/manage")

	//******** c page requires login
	sess := c.GetSession("automezzi")
	if sess == nil {
		c.Redirect("/home", 302)
		return
	}
	flash := beego.NewFlash()
	m := sess.(map[string]interface{})
	fmt.Println(m["admin"])
	fmt.Println(reflect.ValueOf(m["admin"]).Type())
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}

	fmt.Printf("hai i diritti")

	//in caso di panic reindirizza alla home
	defer func(c *MainController) {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Index", r)
			c.Redirect("/home", 302)
		}
	}(c)
	//NON VA SENZA PARAMETRI
	//******** Read users from database
	if c.Ctx.Input.Param(":parms") == "" {
		o := orm.NewOrm()
		o.Using("default")
		var users []models.AuthUser

		o.QueryTable("auth_user")
		//num, err := o.Raw("SELECT id, first, last, email, idkey FROM auth_user",).QueryRows(&users)
		if err != nil {
			flash.Notice("Errore, contattare l'amministratore del sito")
			flash.Store(&c.Controller)
			c.Redirect("/notice", 302)
		}

		//fmt.Println("user nums: ", num)
		for i := range users {
			fmt.Println(users[i])
		}
		rows := "<tr><center><td>ID</td><td>NOME</td><td>COGNOME</td><td>EMAIL</td><td>MODIFICA</td></center></tr>"
		for i := range users {
			rows += fmt.Sprintf("<tr><td>%d</td>"+
				"<td>%s</td><td>%s</td><td>%s</td><td><center><a href='http://%s/manage/user/%s' class=\"user\"> </a></center></td></tr>", users[i].ID, users[i].First, users[i].Last, users[i].Email, appcfgdomainname, users[i].IDkey)
		}
		c.Data["Rows"] = template.HTML(rows)
	}
	const pagesize = 10
	parms := c.Ctx.Input.Param(":parms")
	c.Data["parms"] = parms
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

	qs, ok := c.setCompare(query)
	if !ok {
		fmt.Println("cannot set QuerySeter")
		o := orm.NewOrm()
		qs := o.QueryTable("auth_user")
		qs = qs.Filter("id__gte", 0)
		c.Data["query"] = "id__gte,0"
	}

	count, _ := qs.Count()
	c.Data["count"] = count
	if offset >= count {
		offset = 0
	}
	num, err := qs.Limit(pagesize, offset).OrderBy(order).All(&users)
	if err != nil {
		fmt.Println("Query table failed:", err)
	}
	for i := range users {
		rows += fmt.Sprintf("<tr><td>%d</td>"+
			"<td>%s</td><td>%s</td><td>%s</td><td><center><a href='http://%s/manage/user/%s' class=\"user\"> </a></center></td></tr>", users[i].ID, users[i].First, users[i].Last, users[i].Email, appcfgdomainname, users[i].IDkey)
	}
	c.Data["Rows"] = template.HTML(rows)

	c.Data["order"] = order
	c.Data["offset"] = offset
	c.Data["end"] = max(0, (count/pagesize)*pagesize)

	if num+offset < count {
		c.Data["next"] = num + offset
	}
	if offset-pagesize >= 0 {
		c.Data["prev"] = offset - pagesize
		c.Data["showprev"] = true
	} else if offset > 0 && offset < pagesize {
		c.Data["prev"] = 0
		c.Data["showprev"] = true
	}

	if count > pagesize {
		c.Data["ShowNav"] = true
	}
	c.Data["progress"] = float64(offset*100) / float64(max(count, 1))
}

//UsersManage is for edit accounts by administrator
//TODO quando ritorna al manage lo fa nella pagina 1 e non in quella in cui si trovava l'utente
func (c *MainController) UsersManage() {
	c.activeContent("manage/user")

	//******** c page requires login
	sess := c.GetSession("automezzi")
	if sess == nil {
		c.Redirect("/user/login/home", 302)
		return
	}

	m := sess.(map[string]interface{})
	flash := beego.NewFlash()
	if m["admin"] != 3 {
		c.Redirect("/home", 302)
		flash.Error("Non disponi dei privilegi necessari")
		flash.Store(&c.Controller)
		return
	}

	var x pk.PasswordHash

	x.Hash = make([]byte, 32)
	x.Salt = make([]byte, 16)

	o := orm.NewOrm()
	o.Using("default")
	var IDkey string
	IDkey = c.Ctx.Input.Param(":parms")
	fmt.Println("key: ", IDkey)
	user := models.AuthUser{IDkey: c.Ctx.Input.Param(":parms")}
	err := o.Read(&user, "IDKey")
	if err != nil {
		flash.Error("Internal error")
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

	userAPP := models.AuthApp{ID: user.ID}
	err = o.Read(&userAPP, "ID")
	if err != nil {
		flash.Error("Internal error")
		flash.Store(&c.Controller)
		return
	}

	// c deferred function ensures that the correct fields from the database are displayed
	defer func(c *MainController, user *models.AuthUser, userAPP *models.AuthApp) {
		//check the user lvl
		var userlvllist string
		switch user.Group {
		case 0:
			userlvllist += fmt.Sprintf("<option value=\"0\" selected=\"selected\">Utente</option>" +
				"<option value=\"1\">Utente Speciale</option>" +
				"<option value=\"2\">Agente</option>" +
				"<option value=\"3\">Amministratore</option>")
		case 1:
			userlvllist += fmt.Sprintf("<option value=\"0\">Utente</option>" +
				"<option value=\"1\" selected=\"selected\">Utente Speciale</option>" +
				"<option value=\"2\">Agente</option>" +
				"<option value=\"3\">Amministratore</option>")
		case 2:
			userlvllist += fmt.Sprintf("<option value=\"0\">Utente</option>" +
				"<option value=\"1\">Utente Speciale</option>" +
				"<option value=\"2\" selected=\"selected\">Agente</option>" +
				"<option value=\"3\">Amministratore</option>")
		case 3:
			userlvllist += fmt.Sprintf("<option value=\"0\">Utente</option>" +
				"<option value=\"1\">Utente Speciale</option>" +
				"<option value=\"2\">Agente</option>" +
				"<option value=\"3\" selected=\"selected\">Amministratore</option>")
		default:
			panic("unrecognized escape character")
		}

		fmt.Println(user.BlockControll)
		var checkbloccato string
		if user.BlockControll >= 3 {
			checkbloccato += fmt.Sprintf("<td><input type=\"checkbox\" name=\"blocco\" value=\"bloccato\" checked=\"checked\"> BLOCCATO<br></td>")
			//<td><input type="checkbox" name="apps" value="bloccato"> BLOCCATO<br></td>
		} else {
			checkbloccato += fmt.Sprintf("<td><input type=\"checkbox\" name=\"blocco\" value=\"bloccato\"> BLOCCATO<br></td>")

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

		c.Data["UFirst"] = user.First
		c.Data["ULast"] = user.Last
		c.Data["UEmail"] = user.Email
		c.Data["Userlvllist"] = template.HTML(userlvllist)
		c.Data["Checkbloccato"] = template.HTML(checkbloccato)
		c.Data["Checkautomezzi"] = template.HTML(checkautomezzi)
		c.Data["Checkservizi"] = template.HTML(checkservizi)
	}(c, &user, &userAPP)

	if c.Ctx.Input.Method() == "POST" {
		first := c.GetString("first")
		last := c.GetString("last")
		email := c.GetString("email")
		password := c.GetString("password")
		password2 := c.GetString("password2")
		userlvl := c.GetString("userlvl")
		apps := c.GetStrings("apps")
		blocco := c.GetStrings("blocco")

		valid := validation.Validation{}
		valid.Required(first, "first")
		valid.Email(email, "email")
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

		/******** Compare submitted password with database
		if !pk.MatchPassword(current, &x) {
			flash.Error("Password attuale errata")
			flash.Store(&c.Controller)
			return
		}*/

		//******** Save user info to database
		user.First = first
		user.Last = last
		user.Email = email
		user.LastEditDate = time.Now()
		user.Group = ConvertInt(userlvl)

		if stringInSlice("bloccato", blocco) {
			user.BlockControll = 3
		} else {
			user.BlockControll = 0
		}
		if stringInSlice("automezzi", apps) {
			userAPP.Automezzi = true
		} else {
			userAPP.Automezzi = false
		}
		if stringInSlice("servizi", apps) {
			userAPP.Servizi = true
		} else {
			userAPP.Servizi = false
		}

		_, err := o.Update(&user)
		if err != nil {
			flash.Error("Errore interno")
			flash.Store(&c.Controller)
			return
		}

		_, err = o.Update(&userAPP)
		if err != nil {
			flash.Error("Errore interno")
			flash.Store(&c.Controller)
			return
		}

		flash.Notice("Profilo aggiornato")
		flash.Store(&c.Controller)

	}

}

//stringInSlice funcion check if string is in slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//ConvertInt function convert string in int
func ConvertInt(s string) int {
	//convert string in int
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
