package controllers

import (
	//"fmt"
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

type ErrorData struct {
	Title   string
	Content string
}

func (c *BaseController) Prepare() {
	// page start time
	c.Layout = "layout/main.tpl"

	c.Data["title"] = "canku点餐系统R"
	c.Data["cur"] = ""

	nav := []Link{
		Link{Name: "主页", Url: "HomeController.Get", Cur: "home"},
		Link{Name: "今日订单", Url: "HomeController.Today", Cur: "today"},
	}

	sess_nickname := c.GetSession("nickname")
	sess_isadmin := c.GetSession("isadmin")
	//fmt.Println(sess_nickname)

	if sess_nickname != nil {
		nav = append(nav, Link{Name: "历史订单", Url: "UserController.History", Cur: "history"})
		nav = append(nav, Link{Name: "退出", Url: "UserController.Logout", Cur: "logout"})
	} else {
		nav = append(nav, Link{Name: "注册", Url: "UserController.Register", Cur: "register"})
		nav = append(nav, Link{Name: "登录", Url: "UserController.Login", Cur: "login"})
	}


	if sess_isadmin != 0 && sess_isadmin != nil {
		nav = append(nav, Link{Name: "商家管理", Url: "ShopController.Shop", Cur: "shop"})
	}

	c.Data["nav"] = nav

	//fmt.Println(c.Data["nav"])

	// Setting properties.
	//c.Data["AppDescription"] = utils.AppDescription

	// if app, ok := c.AppController.(NestPreparer); ok {
	//         app.NestPrepare()
	// }
}

//显示错误提示
func (c *BaseController) showmsg(title string, content string) {
	c.Data["title"] = title
	c.Data["content"] = content
	c.Layout = "layout/main.tpl"
	c.TplNames = "error/user_error.tpl"
	c.Render()
	c.StopRun()
}
