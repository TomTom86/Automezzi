package routers

import (
	"automezzi/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/home", &controllers.MainController{})
	beego.Router("/user/login/:back", &controllers.MainController{}, "get,post:Login")
	beego.Router("/user/logout", &controllers.MainController{}, "get:Logout")
	beego.Router("/user/register", &controllers.MainController{}, "get,post:Register")
	beego.Router("/user/profile/", &controllers.MainController{}, "get,post:Profile")
	beego.Router("/manage/user/:parms", &controllers.MainController{}, "get,post:UsersManage")
	beego.Router("/user/check/:uuid", &controllers.MainController{}, "get:Verify")
	beego.Router("/user/remove", &controllers.MainController{}, "get,post:Remove")
	beego.Router("/user/forgot", &controllers.MainController{}, "get,post:Forgot")
	beego.Router("/user/reset/:uuid", &controllers.MainController{}, "get,post:Reset")
	beego.Router("/notice", &controllers.MainController{}, "get:Notice")
	beego.Router("/manage/:parms", &controllers.MainController{}, "get,post:Manage")
	beego.Router("/test/:parms", &controllers.MainController{}, "get,post:Test")
	//beego.Router("/manage/", &controllers.MainController{}, "get,post:Manage")
	beego.Router("/appadmin/index/:parms", &controllers.AdminController{}, "get,post:Index")
	beego.Router("/appadmin/add/:parms", &controllers.AdminController{}, "get,post:Add")
	beego.Router("/appadmin/update/:username", &controllers.AdminController{}, "get,post:Update")
	//beego.Router("/appadmin/manage/:parms", &controllers.AdminController{}, "get,post:Manage")
    beego.Router("/automezzi/acquisto/", &controllers.MainController{}, "get,post:Acquisto")
}
