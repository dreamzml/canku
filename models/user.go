package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	//"time"
)

type User struct {
	Id        int64       //`orm:"column(id);pk"`
	Email     string    `orm:"column(email);size(40)"`
	Nickname  string    `orm:"column(nickname);unique;size(20)"`
	Password  string    `orm:"column(password);size(32)"`
	Lastlogin int64		`orm:"column(lastlogin);type(int)"`
	Isadmin   int8      `orm:"column(isadmin);default(0)"`
}


func (u *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

type ReturnUser struct {
	Id       int
	Email    string
	Nickname string
	Isadmin  int8
}

func (m *User) Select() ReturnUser {
	var u ReturnUser
	orm.NewOrm().Raw("SELECT * FROM user WHERE (email = ? OR nickname = ?) AND password = ? ", m.Email, m.Email, m.Password).QueryRow(&u)
	return u
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

//get user list
func Getuserlist(page int64, page_size int64, sort string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&users)
	count, _ = qs.Count()
	return users, count
}