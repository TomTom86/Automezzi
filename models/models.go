package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type AuthUser struct {
	Id	int  `orm:"auto"`
	First	string `orm:"size(20)"`
	Last	string `orm:"size(20)"`
	Email	string `orm:"unique"`
	Password	string `orm:"size(60)"`
	Is_approved	bool
	Id_key	string `orm:"size(20)"`
	Reg_date	time.Time `orm:"auto_now_add;type(datetime)"`
	Last_login_date	time.Time `orm:"auto_now_add;type(datetime)"`
	Last_edit_date	time.Time `orm:"auto_now_add;type(datetime)"` 
	Reset_key	string `orm:"size(20)"`
	Block_controll	int
	Group	int
	Note	string `orm:"size(100)"`
	AuthApp *AuthApp `orm:"rel(one)"`
}




type AuthApp struct{
	Id int
	Automezzi bool
	Servizi bool 
	AuthUser *AuthUser `orm:"reverse(one)"`
	
	
}

/*
password salt
mobile pin
is locked out
last lock out date
failed password attempt count
password attempt windows start

comment


*/
func init() {
	orm.RegisterModel(new(AuthUser),new(AuthApp))
}
