package controllers

import (
	pk "automezzi/utilities/pbkdf2"
	"automezzi/models"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"github.com/twinj/uuid"
	"html/template"
	"strconv"
	"strings"
	//"time"
	"reflect"
)

type AdminController struct {
	beego.Controller
}

func (c *AdminController) activeAdminContent(view string) {
	c.Layout = "admin-layout.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.TplName = view + ".tpl"
	c.Data["domainname"] = "localhost:8080"

	sess := c.GetSession("automezzi")
	if sess != nil {
		c.Data["InSession"] = 1 // for login bar in header.tpl
		m := sess.(map[string]interface{})
		c.Data["First"] = m["first"]
		c.Data["Admin"] = m["admin"]
		c.Data["IDkey"] = m["idkey"]
	}
}

type compareform struct {
	Comparefield string `form:"comparefield"`
	Compareop    string `form:"compareop"`
	Compareval   string `form:"compareval" valid:"Required"`
}

func (c *AdminController) setCompare(query string) (orm.QuerySeter, bool) {
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

func max(a, b int64) int64 {
	if a < b {
		return b
	}
	return a
}

func (c *AdminController) Index() {
	c.activeAdminContent("appadmin/index")

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

	defer func(c *AdminController) {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Index", r)
			c.Redirect("/home", 302)
		}
	}(c)

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
	domainname := c.Data["domainname"]
	for x := range users {
		i := strings.Index(users[x].RegDate.String(), " ")
		rows += fmt.Sprintf("<tr><td><a href='http://%s/appadmin/update/%s!%s'>%d</a></td>"+
			"<td>%s</td><td>%s</td><td>%s</td><td>%s...</td><td>%s</td><td>%s</td><td>%s</td></tr>", domainname, users[x].Email, parms,
			users[x].ID, users[x].First, users[x].Last, users[x].Email, users[x].Password[:20],
			users[x].IDkey, users[x].RegDate.String()[:i], users[x].ResetKey)
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

func (c *AdminController) Add() {
	c.activeAdminContent("appadmin/add")

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
	parms := c.Ctx.Input.Param(":parms")
	fmt.Println(parms)
	c.Data["parms"] = parms

	if c.Ctx.Input.Method() == "POST" {

		u := authUser{}

		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			return
		}
		fmt.Println(u)
		c.Data["User"] = u
		valid := validation.Validation{}
		if b, _ := valid.Valid(&u); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return
		}
		h := pk.HashPassword(u.Password)

		//******** Save user info to database
		o := orm.NewOrm()
		o.Using("default")

		//create user and userApp models
		userAPP := models.AuthApp{Automezzi: false, Servizi: false}
		user := models.AuthUser{First: u.First, Last: u.Last, Email: u.Email, IsApproved: false, Group: 0, AuthApp: &userAPP}

		// Convert password hash to string
		user.Password = hex.EncodeToString(h.Hash) + hex.EncodeToString(h.Salt)

		// Add user to database with new uuid and send verification email
		key := uuid.NewV4()
		user.IDkey = key.String()

		_, err = o.Insert(&userAPP)
		if err != nil {
			flash.Error("Errore autorizzazioni applicazioni")
			flash.Store(&c.Controller)
			return
		}

		_, err := o.Insert(&user)
		if err != nil {
			flash.Error(u.Email + " gia' registrata")
			flash.Store(&c.Controller)
			return
		}

		flash.Notice("User added")
		flash.Store(&c.Controller)
	}

}

type authUser struct {
	ID         int    `form:"id"`
	First      string `form:"first" valid:"Required"`
	Last       string `form:"last"`
	Email      string `form:"email" valid:"Email"`
	Password   string `form:"password" valid:"MinSize(6)"`
	IDkey      string `form:"idkey"`
	IsApproved bool
	RegDate    string `form:"regdate"` // ParseForm cannot deal with time.Time in the form definition
	ResetKey   string `form:"resetkey"`
	Delete     string `form:"delete,checkbox"`
}

//Update account information
func (c *AdminController) Update() {
	c.activeAdminContent("appadmin/update")
	sess := c.GetSession("automezzi")
	//if you aren't logged redirect to home
	if sess == nil {
		c.Redirect("/home", 302)
		return
	}
	flash := beego.NewFlash()
	m := sess.(map[string]interface{})
	fmt.Println(m["admin"])
	fmt.Println(reflect.ValueOf(m["admin"]).Type())
	//check if you are admin
	if m["admin"] != 3 {
		flash.Notice("Non hai i diritti per accedere a questa pagina")
		flash.Store(&c.Controller)
		c.Redirect("/notice", 302)
	}
	fmt.Printf("hai i diritti")
	defer func(c *AdminController) {
		if r := recover(); r != nil {
			fmt.Println("Recovered in Update", r)
			c.Redirect("/home", 302)
		}
	}(c)

	str := c.Ctx.Input.Param(":username")
	i := strings.Index(str, "!")
	username := str[:i]
	c.Data["parms"] = str[i+1:]
	o := orm.NewOrm()
	o.Using("default")
	user := models.AuthUser{Email: username}
	if err := o.Read(&user, "Email"); err != nil {
		flash.Error("Internal error")
		flash.Store(&c.Controller)
		return
	}

	if c.Ctx.Input.Method() == "POST" {
		u := authUser{}
		if err := c.ParseForm(&u); err != nil {
			fmt.Println("cannot parse form")
			return
		}
		c.Data["User"] = u
		valid := validation.Validation{}
		if b, _ := valid.Valid(&u); !b {
			c.Data["Errors"] = valid.ErrorsMap
			return
		}

		if u.Delete == "on" {
			fmt.Println("about to delete record...")
			_, err := o.Delete(&user)
			if err == nil {
				flash.Notice("Record deleted")
				flash.Store(&c.Controller)
				return
			} else {
				flash.Error("Internal error")
				flash.Store(&c.Controller)
				return
			}
		}

		//******** Save user info to database
		user.First = u.First
		user.Last = u.Last
		user.Email = u.Email
		user.IDkey = u.IDkey
		user.ResetKey = u.ResetKey

		o := orm.NewOrm()
		o.Using("default")

		// Update user record
		_, err := o.Update(&user)
		if err != nil {
			flash.Error("Update failed")
			flash.Store(&c.Controller)
			return
		}

		flash.Error("User updated")
		flash.Store(&c.Controller)
	} else {
		c.Data["User"] = user
	}
}
