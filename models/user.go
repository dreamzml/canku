package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	"time"
)

type User struct {
	Id        int64     `orm:"auto"`
	Email     string    `orm:"size(40)"`
	Nickname  string    `orm:"unique;size(20)"`
	Password  string    `orm:"size(32)"`
	Lastlogin time.Time `orm:"auto_now_add;type(datetime)"`
	Isadmin   int8
}

func (m *User) Insert() {
	o := orm.NewOrm()
	id, err := o.Insert(m)
	if err == nil {
		fmt.Println(id)
	}
}

func (m *User) Read() {
	var user User
	o := orm.NewOrm()
	err := o.Raw("SELECT id, nickname FROM user WHERE email = ? AND password = ?", m.Email, m.Password).QueryRow(&user)
	if err == nil {
		fmt.Println(user.Email)
	}
}
