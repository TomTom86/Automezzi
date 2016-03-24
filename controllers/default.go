package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) activeContent(view string) {
	c.Layout = "basic-layout.tpl"
	c.Data["domainname"] = "localhost:8080"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "header.tpl"
	c.LayoutSections["Sidebar"] = "sidebar.tpl"
	c.LayoutSections["Footer"] = "footer.tpl"
	c.TplName = view + ".tpl"
	//c.Data["domainname"] = "localhost:8080"

	sess := c.GetSession("automezzi")
	if sess != nil {
		c.Data["InSession"] = 1 // for login bar in header.tpl
		m := sess.(map[string]interface{})
		c.Data["First"] = m["first"]
		c.Data["Admin"] = m["admin"]
		c.Data["IDkey"] = m["idkey"]
        fmt.Println(m["automezzi"])
		c.Data["Automezzi"] = m["automezzi"]
	}
}

func (c *MainController) Get() {
	c.activeContent("index")

	//******** c page requires login
	sess := c.GetSession("automezzi")
	if sess == nil {
		c.Redirect("/user/login/home", 302)
		return
	}
  
	m := sess.(map[string]interface{})
	fmt.Println("username is", m["username"])
	fmt.Println("logged in at", m["timestamp"])
    
    flash := beego.ReadFromRequest(&c.Controller)
    if _, ok := flash.Data["notice"]; ok {
        // Display settings successful
        c.TplName = "notice.tpl"
    } else if _, ok = flash.Data["error"]; ok {
        // Display error messages
        c.TplName = "error.tpl"
    } else {
        // Display default settings page
        c.Data["list"] = GetInfo()
        c.TplName = "setting_list.tpl"
    }
}



//Notice show flash message
func (c *MainController) Notice() {
	c.activeContent("notice")

	flash := beego.ReadFromRequest(&c.Controller)
	if n, ok := flash.Data["notice"]; ok {
		c.Data["notice"] = n
	}
}


func GetInfo() string{
	return "ciao"
}