package controllers

import (
	"canku/models"
	"fmt"
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/session"
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
	// v := this.GetSession("asta")
	// if v == nil {
	// 	this.SetSession("asta", int(1))
	// } else {
	// 	this.SetSession("asta", v.(int)+1)
	// }

	// fmt.Println(this.GetSession("asta"))
	//去掉布局
	this.Layout = ""
	this.TplNames = "user/login.tpl"
	this.Render()
}

/**
* 登录接收
 */
func (this *UserController) Signup() {
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

	var Ru models.ReturnUser
	Ru = user.Select()

	if Ru.Id > 0 {
		this.SetSession("nickname", Ru.Nickname)
		this.SetSession("email", Ru.Email)
		this.SetSession("isadmin", Ru.Isadmin)
		this.Redirect("/", 302)
	}
	this.Redirect("/login", 302)
}

/**
* 注册
 */
func (this *UserController) Register() {
	//去掉布局
	this.Layout = ""
	this.TplNames = "user/register.tpl"
	this.Render()
}

/**
 * 注册逻辑处理
 * @param  {[type]} this *UserController) Join( [description]
 * @return {[type]}      [description]
 */
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

	if len(errmsg) == 0 {
		var user models.User
		user.Email = requestEmail
		user.Nickname = requestNickname
		user.Password = models.Md5([]byte(requestPassword))
		if err := user.Insert(); err != nil {
			this.Ctx.WriteString(err.Error())
		}
	}
	this.SetSession("nickname", requestNickname)
	this.SetSession("email", requestEmail)
	this.SetSession("isadmin", 0)
	this.Redirect("/", 302)
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
	this.Render()
}

/**
 * 用户自己的历史订单
 * @param  {[type]} this *UserController) History( [description]
 * @return {[type]}      [description]
 */
func (this *UserController) History() {

}
