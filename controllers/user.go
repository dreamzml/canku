package controllers

import (
	"github.com/astaxie/beego"
	//"github.com/beego/canku/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Login() {
	this.TplNames = "login.tpl"
}

func (this *UserController) Signup() {
	this.TplNames = "login.tpl"
}

func (this *UserController) Register() {
	this.TplNames = "register.tpl"
}

func (this *UserController) Logout() {

}

func (this *UserController) Forget() {
	this.TplNames = "forget.tpl"
}
