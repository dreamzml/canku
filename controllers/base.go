package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

//导航链接结构体
type Link struct {
	Name string
	Url  string
	Cur  string
}

func (c *BaseController) Prepare() {
	// page start time
	c.Layout = "layout/main.tpl"

	c.Data["title"] = "canku点餐系统R"
	c.Data["cur"] = ""

	c.Data["nav"] = []Link{
		Link{Name: "主页", Url: "HomeController.Get", Cur: "home"},
		Link{Name: "关于我们", Url: "/about", Cur: "about"},
		Link{Name: "注册", Url: "UserController.Register", Cur: "register"},
		Link{Name: "登录", Url: "UserController.Login", Cur: "login"},
	}

	// Setting properties.
	//c.Data["AppDescription"] = utils.AppDescription

	// if app, ok := c.AppController.(NestPreparer); ok {
	//         app.NestPrepare()
	// }
}
