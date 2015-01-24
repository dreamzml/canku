package models

import (
	//"fmt"
	//"float"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	//"time"
)


type Food struct {
	Id        		int64       //`orm:"column(id);pk"`
	CategoryId    	int64    	`orm:"column(category_id)"`
	Title     		string    	`orm:"column(title);size(500)"`
	Price     		float32   	`orm:"column(price);type(float)"`
	ShopId	  		float32   	`orm:"column(shop_id);type(float)"`
}

func (u *Food) TableName() string {
	return "food"
}

func init() {
	orm.RegisterModel(new(Food))
}