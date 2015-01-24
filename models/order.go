package models

import (
	//"fmt"
	//"float"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	//"time"
)

type Order struct {
	Id        		int64       //`orm:"column(id);pk"`
	ShopId    		int64    	`orm:"column(shop_id)"`
	ShopName     	string    	`orm:"column(shop_name);size(500)"`
	Desp    		int64    	`orm:"column(desp)"`
	FoodCount     	int    		`orm:"column(food_count)"`
	Price     		float32   	`orm:"column(price)"`
	CreateTime		int64		`orm:"column(create_time)"`
	UserId	  		int64   	`orm:"column(user_id)"`
}

func (u *Order) TableName() string {
	return "order"
}

func init() {
	orm.RegisterModel(new(Order))
}