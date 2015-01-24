package models

import (
	//"fmt"
	//"float"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import mysql driver
	//"time"
)

type OrderFoods struct {
	Id        		int64       //`orm:"column(id);pk"`
	OrderId    		int64    	`orm:"column(order_id)"`
	FoodId	     	int64    	`orm:"column(shop_name);size(500)"`
	FootName    	string    	`orm:"column(desp)"`
	ShopId     		int64    	`orm:"column(food_count)"`
	Price     		float32   	`orm:"column(price)"`
	UserId	  		int64   	`orm:"column(user_id)"`
	UserName	  	int64   	`orm:"column(user_name)"`
	CreateTime		int64		`orm:"column(create_time)"`
}

func (u *OrderFoods) TableName() string {
	return "order_foods"
}

func init() {
	orm.RegisterModel(new(OrderFoods))
}