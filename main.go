package main

import (
	_ "canku/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.SetStaticPath("css", "static/css")
	beego.SetStaticPath("js", "static/js")
	beego.Run()
}
