package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type AuthUser struct {
	Id        int
	First     string
	Last      string
	Email     string `orm:"unique"`
	Password  string
	Id_key	  string
	Reg_key   string
	Reg_date  time.Time `orm:"auto_now_add;type(datetime)"`
	Reset_key string
	Block_controll int
	Group 	  int
}


/*
password
password salt
mobile pin
email
password question
is approved
is locked out
create data
last login date
last password change date
last lock out date
failed password attempt count
password attempt windows start

comment


*/
func init() {
	orm.RegisterModel(new(AuthUser))
}
