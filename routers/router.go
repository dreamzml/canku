package routers

import (
	"canku/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})             // 首页
	beego.Router("/login", &controllers.LoginController{})       // 登陆
	beego.Router("/logout", &controllers.LogoutController{})     // 退出
	beego.Router("/register", &controllers.RegisterController{}) // 注册
	beego.Router("/forget", &controllers.ForgetController{})     // 忘记密码
}
