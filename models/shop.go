package models

import (
	//"fmt"
	//"float"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	//"time"
)


type Shop struct {
	Id        int64       //`orm:"column(id);pk"`
	Title     string    `orm:"column(title);unique;size(255)"`
	Address   string    `orm:"column(address);size(500)"`
	MinPrice  float32   `orm:"column(minprice);type(float)"`
	Freight	  float32   `orm:"column(freight);type(float)"`
	Tel		  string    `orm:"column(tel);unique;size(20)"`
}

func (u *Shop) TableName() string {
	return "shop"
}

func init() {
	orm.RegisterModel(new(Shop))
}