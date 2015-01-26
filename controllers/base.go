package controllers

import (
	//"fmt"
	"strings"
	//"strconv"
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
	Items []Link
}

/**
 * 页面载入之前
 * @param  {[type]} c *BaseController) Prepare( [description]
 * @return {[type]}   [description]
 */
func (c *BaseController) Prepare() {
	//预设layout布局模板信息
	c.PrepareLayout()
}

/**
* 预设布局模板信息
*/
func (c *BaseController) PrepareLayout() {
	c.Layout = "layout/main.tpl"

	c.Data["title"] = "canku点餐系统R"
	c.Data["cur"] = ""

	nav := []Link{
		Link{Name: "主页", Url: "HomeController.Get", Cur: "home"},
		Link{Name: "今日订单", Url: "HomeController.Today", Cur: "today"},
	}

	sess_nickname := c.GetSession("nickname")
	sess_isadmin := c.GetSession("isadmin")

	var adminCode int8 = 1
	if sess_isadmin == adminCode {
		//商家管理展开选择
		shopManageItems := []Link{
			Link{Name: "商家管理", Url: "admin/ShopController.Index", Cur: "shopmanage"},
			Link{Name: "用户管理", Url: "admin/UserController.Index", Cur: "todayshop"},
		}
		nav = append(nav, Link{Name: "商家管理", Url: "dropdown", Cur: "shop", Items: shopManageItems})
	}

	if sess_nickname != nil {
		nav = append(nav, Link{Name: "历史订单", Url: "UserController.History", Cur: "history"})
		nav = append(nav, Link{Name: strings.Join([]string{"退出(", sess_nickname.(string), ")"}, ""), Url: "UserController.Logout", Cur: "logout"})
	} else {
		nav = append(nav, Link{Name: "注册", Url: "UserController.Register", Cur: "register"})
		nav = append(nav, Link{Name: "登录", Url: "UserController.Login", Cur: "login"})
	}

	c.Data["nav"] = nav
}

/**
 * 显示错误提示
 * @param  {[type]} c *BaseController) showmsg(title string, content string [description]
 * @return {[type]}   [description]
 */
func (c *BaseController) showmsg(title string, content string) {
	c.Data["title"] = title
	c.Data["content"] = content
	c.Layout = "layout/main.tpl"
	c.TplNames = "error/user_error.tpl"
	c.Render()
	c.StopRun()
}
