package controllers

type HomeController struct {
	BaseController
}
type LoginController struct {
	BaseController
}
type LogoutController struct {
	BaseController
}
type RegisterController struct {
	BaseController
}
type ForgetController struct {
	BaseController
}

func (c *HomeController) Get(){
	//c.Data["title"] = "canku点餐系统"
	c.Data["cur"] = "home";
	c.TplNames = "home/index.tpl" 
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
