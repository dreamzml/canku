package controllers

import (
	//"canku/models"
	//"fmt"
	"github.com/astaxie/beego/validation"
	"log"
)

type UserController struct {
	BaseController
}

type user struct {
	Email    string
	Nickname string
	Password string
}

/**
* 登录入口
 */

func (this *UserController) Login() {
	//去掉布局
	this.Layout = ""

	this.TplNames = "user/login.tpl"
}

/**
* 登录接收
 */
func (this *UserController) Signup() {

	beego.AutoRender = false
	requestEmail := this.GetString("email")
	requestPassword := this.GetString("password")
	//this.Ctx.WriteString(requestEmail)
	u := user{requestEmail, requestPassword}
	valid := validation.Validation{}
	valid.Required(u.Email, "email")
	valid.MaxSize(u.Email, 40, "emailMax")
	valid.Required(u.Password, "password")
	if valid.HasErrors() {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}
	//this.Ctx.WriteString(email)
	//this.UserModel.insert(email, nickname, password)

}

/**
* 注册
 */
func (this *UserController) Register() {
	//去掉布局
	this.Layout = ""

	this.TplNames = "user/register.tpl"
}

func (this *UserController) Join() {
	beego.AutoRender = false
	//this.user.Insert(user)
}

/**
* 退出
 */
func (this *UserController) Logout() {

}

/**
* 忘记密码
 */
func (this *UserController) Forget() {
	this.TplNames = "user/forget.tpl"
}
