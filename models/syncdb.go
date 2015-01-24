package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//. "github.com/beego/admin/src/lib"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	//_ "github.com/mattn/go-sqlite3"
)

var o orm.Ormer

/**
* 同步数据库
*/
func Syncdb() {

	createdb()
	Connect()
	o = orm.NewOrm()

	// 数据库别名
	name := "default"
	// drop table 后再建表
	force := true
	// 打印执行过程
	verbose := true
	// 遇到错误立即返回

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
	insertUser()
	insertCategory()
	insertShop()
	insertFood()
	//insertNodes()
	fmt.Println("database init is complete.\nPlease restart the application")

}


/**
* 创建数据库
*/
func createdb() {
	fmt.Printf("run dbport \n")
	db_type := beego.AppConfig.String("dbdriver")
	db_host := beego.AppConfig.String("dbhost")
	db_port := beego.AppConfig.String("dbport")
	db_user := beego.AppConfig.String("dbuser")
	db_pass := beego.AppConfig.String("dbpass")
	db_name := beego.AppConfig.String("dbname")


	var dns string
	var sqlstring string
	switch db_type {
	case "mysql":
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8", db_user, db_pass, db_host, db_port)
		sqlstring = fmt.Sprintf("CREATE DATABASE  if not exists `%s` CHARSET utf8 COLLATE utf8_general_ci", db_name)
		break
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}
	db, err := sql.Open(db_type, dns)
	if err != nil {
		panic(err.Error())
	}
	r, err := db.Exec(sqlstring)
	if err != nil {
		log.Println(err)
		log.Println(r)
	} else {
		log.Println("Database ", db_name, " created")
	}
	defer db.Close()

}

/**
* 创建会员表数据
*/
func insertUser() {
	fmt.Println("insert user ...")
	
	users := []User{
		{Nickname: "admin", Password: "admin", Email: "admin@admin.com", Isadmin: 1, Lastlogin:time.Now().Unix()},
		{Nickname: "dome", Password: "dome", Email: "dome@dome.com", Isadmin: 1, Lastlogin:time.Now().Unix()},
	}
	for _, v := range users {
		u := new(User)
		u.Nickname = v.Nickname
		u.Password = Md5([]byte(v.Password))
		u.Email = v.Email
		u.Isadmin = v.Isadmin
		u.Lastlogin = v.Lastlogin
		o = orm.NewOrm()
		o.Insert(u)
	}

	fmt.Println("insert user end")
}

/**
* 创建分类表数据
*/
func insertCategory() {
	fmt.Println("insert category ...")

	categorys := []Category{
		{Title: "湘菜", ShopId: 1},
		{Title: "川菜", ShopId: 2},
	}
	for _, v := range categorys {
		g := new(Category)
		g.Title = v.Title
		g.ShopId = v.ShopId
		o.Insert(g)
	}

	fmt.Println("insert category end")
}

/**
* 创建分类表数据
*/
func insertShop() {
	fmt.Println("insert shop ...")

	lists := []Shop{
		{Title: "美味小厨港式烧腊", Address: "灵石路697号", MinPrice: 30.00, Freight: 5.00, Tel: "4000557117"},
		{Title: "悦宾饭店", Address: "上海市沪太路546号", MinPrice: 25.00, Freight: 0.00, Tel: "13564853119"},
	}
	for _, v := range lists {
		g := new(Shop)
		g.Title = v.Title
		g.Address = v.Address
		g.MinPrice = v.MinPrice
		g.Freight = v.Freight
		g.Tel = v.Tel
		o.Insert(g)
	}

	fmt.Println("insert shop end")
}


/**
* 创建食物表数据
*/
func insertFood() {
	fmt.Println("insert food ...")

	lists := []Food{
		{CategoryId: 1, Title: "广东豉油鸡", Price: 16.00, ShopId: 2},
		{CategoryId: 2, Title: "麻辣豆腐饭", Price: 14.00, ShopId: 1},
	}
	for _, v := range lists {
		g := new(Food)
		g.CategoryId = v.CategoryId
		g.Title = v.Title
		g.Price = v.Price
		g.ShopId = v.ShopId
		o.Insert(g)
	}

	fmt.Println("insert food end")
}

