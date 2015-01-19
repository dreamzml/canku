package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	"time"
)

type User struct {
	Id        int64     `orm:"pk"`
	Email     string    `orm:"size(40)"`
	Nickname  string    `orm:"unique;size(20)"`
	Password  string    `orm:"size(32)"`
	Lastlogin time.Time `orm:"auto_now_add;type(datetime)"`
	Isadmin   int8
}

func (m *User) Insert() {
	o := orm.NewOrm()
	fmt.Println(o.Insert(m))
}
