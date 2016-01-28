package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) activeContent(view string) {
	this.Layout = "basic-layout.tpl"
	this.Data["domainname"] = "localhost:8080"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Header"] = "header.tpl"
	this.LayoutSections["Sidebar"] = "sidebar.tpl"
	this.LayoutSections["Footer"] = "footer.tpl"
	this.TplNames = view + ".tpl"
	this.Data["domainname"] = "localhost:8080"

	sess := this.GetSession("automezzi")
	if sess != nil {
		this.Data["InSession"] = 1 // for login bar in header.tpl
		m := sess.(map[string]interface{})
		this.Data["First"] = m["first"]
		this.Data["Admin"] = m["admin"]
		this.Data["ID_key"] = m["id_key"]
		this.Data["Automezzi"] = m["automezzi"]
	}
}

func (this *MainController) Get() {
	this.activeContent("index")

	//******** This page requires login
	sess := this.GetSession("automezzi")
	if sess == nil {
		this.Redirect("/user/login/home", 302)
		return
	}
	m := sess.(map[string]interface{})
	fmt.Println("username is", m["username"])
	fmt.Println("logged in at", m["timestamp"])
}

func (this *MainController) Notice() {
	this.activeContent("notice")

	flash := beego.ReadFromRequest(&this.Controller)
	if n, ok := flash.Data["notice"]; ok {
		this.Data["notice"] = n
	}
}
