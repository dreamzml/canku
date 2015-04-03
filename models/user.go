package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	"errors"
	"github.com/astaxie/beego/validation"
	//"time"
)

type User struct {
	Id        int64       //`orm:"column(id);pk"`
	Email     string    `orm:"column(email);unique;null;size(40)" form:"Email" valid:"Required;Email"`
	Nickname  string    `orm:"column(nickname);unique;null;size(20)" form:"Nickname" valid:"Required;MaxSize(20);MinSize(6)"`
	Password  string    `orm:"column(password);null;size(32)" form:"Nickname" valid:"Required;MaxSize(20);MinSize(6)"`
	Lastlogin int64		`orm:"column(lastlogin);type(int)"`
	Isadmin   int8      `orm:"column(isadmin);default(0)" form:"Isadmin"`
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

//验证用户信息
func (m *User) checkUser(u *User) (err error) {

	//return errors.New("err.Message")
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}


func (m *User) Select() ReturnUser {
	var u ReturnUser
	orm.NewOrm().Raw("SELECT * FROM user WHERE (email = ? OR nickname = ?) AND password = ? ", m.Email, m.Email, m.Password).QueryRow(&u)
	return u
}

func (m *User) Insert() error {
	if err := m.checkUser(m); err != nil {
		return err
	}

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