package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

type Link struct{
    Name 	string
    Url  	string
    Cur 	string
}


func (c *BaseController) Prepare() {


    // page start time
	c.Layout = "layout/main.tpl"

    c.Data["title"] = "canku点餐系统R"
    c.Data["cur"] = "";

	c.Data["nav"] = []Link{
						Link{Name:"主页",Url:"HomeController.Get",Cur:"home"},
						Link{Name:"关于我们",Url:"/about",Cur:"about"},
						Link{Name:"联系我",Url:"/contact",Cur:"contact"},
						Link{Name:"登录",Url:"/login",Cur:"login"},
					}

    // Setting properties.
    //c.Data["AppDescription"] = utils.AppDescription


    // if app, ok := c.AppController.(NestPreparer); ok {
    //         app.NestPrepare()
    // }
}