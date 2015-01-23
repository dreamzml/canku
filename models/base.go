package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)


//数据库连接
func Connect() {
	var dns string
	db_type := beego.AppConfig.String("dbdriver")
	db_host := beego.AppConfig.String("dbhost")
	db_port := beego.AppConfig.String("dbport")
	db_user := beego.AppConfig.String("dbuser")
	db_pass := beego.AppConfig.String("dbpass")
	db_name := beego.AppConfig.String("dbname")

	switch db_type {
	case "mysql":
		orm.RegisterDriver("mysql", orm.DR_MySQL)
		dns = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", db_user, db_pass, db_host, db_port, db_name)
		break
	default:
		beego.Critical("Database driver is not allowed:", db_type)
	}
	orm.RegisterDataBase("default", db_type, dns)
}

func Md5(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
