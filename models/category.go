package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	//"time"
)


type Category struct {
	Id        int64       //`orm:"column(id);pk"`
	Title     string    `orm:"column(name);size(40)"`
	ShopId    int64    `orm:"column(ahop_id);unique;size(20)"`
}

func (u *Category) TableName() string {
	return "category"
}

func init() {
	orm.RegisterModel(new(Category))
}