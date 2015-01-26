package routers

import (
	"github.com/dreamzml/canku/controllers"
	"github.com/dreamzml/canku/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {

	beego.Router("/", &controllers.HomeController{})
	beego.Router("/today", &controllers.HomeController{}, "get:Today")        // 今日订单
	beego.Router("/login", &controllers.UserController{}, "get:Login")        // 登陆页面
	beego.Router("/signup", &controllers.UserController{}, "post:Signup")     // 登录逻辑判断
	beego.Router("/logout", &controllers.UserController{}, "get:Logout")      // 退出
	beego.Router("/register", &controllers.UserController{}, "get:Register")  // 注册页面
	beego.Router("/join", &controllers.UserController{}, "post:Join")         // 注册逻辑处理
	beego.Router("/forget", &controllers.UserController{}, "get:Forget")      // 忘记密码
	beego.Router("/sendmail", &controllers.UserController{}, "post:Sendmail") // 忘记密码 发送邮件
	beego.Router("/history", &controllers.UserController{}, "get:History")    // 历史订单

	beego.Router("/err", &controllers.UserController{}, "get:Err")

	beego.Router("/admin/user", &admin.UserController{}, "*:Index")
	beego.Router("/admin/shop", &admin.ShopController{}, "*:Index")
}
