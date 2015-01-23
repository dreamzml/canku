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

func (this *UserController) Err() {
	this.showmsg("User Error code 001", "User Error code 001")
}

/**
 * 登录入口
 */
func (this *UserController) Login() {
	//去掉布局
	this.Layout = ""
	this.TplNames = "user/login.tpl"
	this.Render()
}

/**
 * 登录接收
 */
func (this *UserController) Signup() {
	//errmsg := make(map[string]string)
	requestEmail := this.GetString("email")
	requestPassword := this.GetString("password")

	svalid := validation.Validation{}

	svalid.Required(requestEmail, "email")
	svalid.MaxSize(requestEmail, 40, "email")
	svalid.Required(requestPassword, "password")

	if svalid.HasErrors() {
		for _, err := range svalid.Errors {
			//fmt.Println(err.Key, err.Message)
			this.showmsg("Error Message", "["+err.Key+"]"+err.Message)
		}
	}

	var user models.User
	user.Email = requestEmail
	user.Password = models.Md5([]byte(requestPassword))
fmt.Printf("user:%s\n\n",user)
	var Ru models.ReturnUser
	Ru = user.Select()
fmt.Printf("Ru:%s\n\n",user)
	if Ru.Id > 0 {
		this.SetSession("nickname", Ru.Nickname)
		this.SetSession("email", Ru.Email)
		this.SetSession("isadmin", Ru.Isadmin)
		this.Redirect("/", 302)
	} else {
		this.showmsg("Error", "Email OR Password IS WRONG")
	}
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

	jvalid := validation.Validation{}

	jvalid.Required(requestEmail, "email")
	jvalid.MaxSize(requestEmail, 40, "email")
	jvalid.Required(requestNickname, "nickname")
	jvalid.MaxSize(requestNickname, 20, "nickname")
	jvalid.Required(requestPassword, "password")

	fmt.Println(jvalid.HasErrors())

	if jvalid.HasErrors() {
		for _, err := range jvalid.Errors {
			//fmt.Println(err.Key, err.Message)
			this.showmsg("Error Message", "["+err.Key+"]"+err.Message)
		}
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
	beego.AutoRender = false
	this.DelSession("nickname")
	this.DelSession("email")
	this.DelSession("isadmin")
	this.Redirect("/", 302)
	this.StopRun()
}

/**
 * 忘记密码
 */
func (this *UserController) Forget() {
	this.Layout = ""
	this.TplNames = "user/forget.tpl"
	this.Render()
}

func (this *UserController) Sendmail() {
	beego.AutoRender = false
}

/**
 * 用户自己的历史订单
 * @param  {[type]} this *UserController) History( [description]
 * @return {[type]}      [description]
 */
func (this *UserController) History() {

}
