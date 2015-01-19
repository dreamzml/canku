package controllers

import (
	//"canku/models"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"log"
)

type UserController struct {
	beego.Controller
}

type user struct {
	Email    string
	Password string
}

func (this *UserController) Login() {
	this.TplNames = "login.tpl"
}

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

func (this *UserController) Register() {
	this.TplNames = "register.tpl"
}

func (this *UserController) Join() {
	beego.AutoRender = false
	//this.user.Insert(user)
}

func (this *UserController) Logout() {

}

func (this *UserController) Forget() {
	this.TplNames = "forget.tpl"
}
