package controllers

type UserController struct {
	BaseController
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
	this.TplNames = "user/login.tpl"
}

/**
* 注册
*/
func (this *UserController) Register() {
	//去掉布局
	this.Layout = "" 

	this.TplNames = "user/register.tpl"
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
