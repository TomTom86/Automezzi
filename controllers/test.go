package controllers

import (
	"automezzi/models"
	//pk "automezzi/utilities/pbkdf2wrapper"
	//"encoding/hex"
	"fmt"
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/validation"
	//"github.com/twinj/uuid"
	//"html/template"
	"strconv"
	//"strings"
	//"time"
    //"reflect"
)

func (this *MainController) Test() {
	this.activeContent("test/test")
	//var app models.AuthApp	
	//var lists []orm.ParamsList
	
	//var lists,list1 []orm.ParamsList
	//user := models.AuthUser{Id_key: this.Ctx.Input.Param(":parms")}
	//err := o.Read(&user, "Id_key")
	
	
	parms, err := strconv.Atoi(this.Ctx.Input.Param(":parms"))
	user := models.AuthUser{Id:parms}
	o := orm.NewOrm()
	//o.QueryTable("auth_user").Filter("id",parms).RelatedSel().One(user)
	//err := o.QueryTable(app).Filter("id", parms).One(&app)
	
	err = o.Read(&user, "Id")

	if err == orm.ErrNoRows {
	    fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
	    fmt.Println("No primary key found.")
	} else {
	    fmt.Println(user.Id, user.First)
	}
		
	/*
	o.QueryTable("auth_user").Filter("Id", parms).RelatedSel().One(user)
	fmt.Println(user)
	fmt.Println(user.AuthApp)
	//fmt.Println(user.AuthApp.Automezzi)
	*/
	
	
	


}