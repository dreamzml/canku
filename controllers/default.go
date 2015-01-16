package controllers

import (
	"github.com/astaxie/beego"
)

type HomeController struct {
	beego.Controller
}
type LoginController struct {
	beego.Controller
}
type LogoutController struct {
	beego.Controller
}
type RegisterController struct {
	beego.Controller
}
type ForgetController struct {
	beego.Controller
}

func (c *HomeController) Get() {

}

func (c *LoginController) Get() {
	c.TplNames = "login.tpl"
}

func (c *LogoutController) Get() {

}

func (c *RegisterController) Get() {
	c.TplNames = "register.tpl"
}

func (c *ForgetController) Get() {
	c.TplNames = "forget.tpl"
}
