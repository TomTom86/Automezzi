package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type AuthUser struct {
	Id	int
	First	string
	Last	string
	Email	string `orm:"unique"`
	Password	string
	Is_approved	bool
	Id_key	string
	//Reg_key   string
	Reg_date	time.Time `orm:"auto_now_add;type(datetime)"`
	Last_login_date	time.Time `orm:"auto_now_add;type(datetime)"`
	Last_edit_date	time.Time `orm:"auto_now_add;type(datetime)"` 
	Reset_key	string
	Block_controll	int
	Group	int
	Note	string
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
	orm.RegisterModel(new(AuthUser))
}
