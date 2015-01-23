package main

import (
	"mime"
	"os"

	_ "canku/routers"
	"github.com/astaxie/beego"
	"canku/models"
)

func main() {
	beego.SetStaticPath("css", "static/css")
	beego.SetStaticPath("js", "static/js")

	//初始化
	initialize()

	beego.SessionOn = true
	beego.AutoRender = false

	beego.Run()
}

//初始化函数
func initialize() {
	mime.AddExtensionType(".css", "text/css")
	//判断初始化参数
	initArgs()
	models.Connect()
}

//执行命令行传入的 指定命令
func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-syncdb" {
			models.Syncdb()
			os.Exit(0)
		}
	}
}