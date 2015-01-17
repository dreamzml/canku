package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql/" // import mysql driver
)

type User struct {
	id       int
	nickname string `orm:"size(20)"`
	password string `orm:"size(32)"`
	regtime  string `orm:"size(11)"`
	isadmin  int
}

func init() {
	orm.RegisterModel(new(User))

	orm.RegisterDatabase("default", databasedriver, mysqldbname+":"+mysqldbpass+"@/"+mysqldbname+"?charset="+mysqldbchar, 30)
}

func main() {
	o := orm.NewOrm()

}
