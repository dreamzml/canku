package controllers

import (
	"canku/models"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	//"log"
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
	errmsg := make(map[string]string)
	requestEmail := this.GetString("email")
	requestPassword := this.GetString("password")

	valid := validation.Validation{}
	if v := valid.Required(requestEmail, "email"); !v.Ok {
		errmsg["email"] = "请输入邮箱地址"
	} else if v := valid.MaxSize(requestEmail, 40, "email"); !v.Ok {
		errmsg["email"] = "邮箱地址长度不能大于40个字符"
	}
	var user models.User
	user.Email = requestEmail
	user.Password = models.Md5([]byte(requestPassword))

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
	errmsg := make(map[string]string)
	requestEmail := this.GetString("email")
	requestNickname := this.GetString("nickname")
	requestPassword := this.GetString("password")

	valid := validation.Validation{}
	if v := valid.Required(requestEmail, "email"); !v.Ok {
		errmsg["email"] = "请输入邮箱地址"
	} else if v := valid.MaxSize(requestEmail, 40, "email"); !v.Ok {
		errmsg["email"] = "邮箱地址长度不能大于40个字符"
	}

	if v := valid.Required(requestNickname, "nickname"); !v.Ok {
		errmsg["nickname"] = "请输入昵称"
	} else if v := valid.MaxSize(requestNickname, 20, "nickname"); !v.Ok {
		errmsg["nickname"] = "昵称长度不能大于20个字符"
	}

	var user models.User
	user.Email = requestEmail
	user.Nickname = requestNickname
	user.Password = models.Md5([]byte(requestPassword))
	models.Init()
	user.Insert()
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
