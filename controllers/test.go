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
	
	
	parms,  err:= strconv.Atoi(this.Ctx.Input.Param(":parms"))
	if err != nil {
		fmt.Println(err)
	}
	user := models.AuthUser{Id:parms}
	o := orm.NewOrm()
	
/*	var maps []orm.Params
	num, err := o.QueryTable("auth_user").Values(&maps, "id", "First", "id_key", "auth_app__automezzi")
	if err == nil {
	    fmt.Printf("Result Nums: %d\n", num)
	    for _, m := range maps {
	        fmt.Println(m["Id"], m["First"], m["id_key"], m["Auth_App__Automezzi"])
	    // There is no complicated nesting data in the map
	    }
	}
	
	*/
	
	err = o.Read(&user, "Id")

	if err == orm.ErrNoRows {
	    fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
	    fmt.Println("No primary key found.")
	} else {
	    fmt.Println(user.Id, user.First)
	}
			

	err = o.QueryTable("auth_user").Filter("Id", parms).RelatedSel().One(&user)
	if err == orm.ErrNoRows {
	    fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
	    fmt.Println("No primary key found.")
	} else {
	    fmt.Println(user)
		fmt.Println(user.AuthApp.Automezzi)
	}

	
	
	


}