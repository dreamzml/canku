package routers

import (
	"canku/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.HomeController{})
	// 首页
	beego.Router("/login", &controllers.UserController{}, "get:Login")       // 登陆页面
	beego.Router("/signup", &controllers.UserController{}, "post:Signup")    // 登录逻辑判断
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")     // 退出
	beego.Router("/register", &controllers.UserController{}, "get:Register") // 注册页面
	beego.Router("/join", &controllers.UserController{}, "post:Join")        // 注册逻辑处理
	beego.Router("/forget", &controllers.UserController{}, "get:Forget")     // 忘记密码

}
