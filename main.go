package main

import (
	_ "canku/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("css", "static/css")
	beego.SetStaticPath("js", "static/js")

	beego.SessionOn = true
	//beego.AutoRender = false

	beego.Run()
}
