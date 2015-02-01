package controllers

import (
	//"fmt"
	"strings"
	//"strconv"
	"github.com/astaxie/beego"
	"github.com/dreamzml/canku/lib"

	//"os"
	"github.com/beego/i18n"
)

type BaseController struct {
	beego.Controller
}

var(
	langTypes []langType
)

//导航链接结构体
type Link struct {
	Name string
	Url  string
	Cur  string
	Items []Link
}

//语言设定
type langType struct {
	Lang string
	Name  string
}

/**
 * 页面载入之前
 * @param  {[type]} c *BaseController) Prepare( [description]
 * @return {[type]}   [description]
 */
func (c *BaseController) Prepare() {
	//预设layout布局模板信息
	c.PrepareLayout()

}

/**
* 预设布局模板信息
*/
func (c *BaseController) PrepareLayout() {
	//初始化语言包
	c.settingLocales()
	c.setLangVer()

	c.Layout = "layout/main.tpl"

	c.Data["title"] = "canku点餐系统R"
	c.Data["cur"] = ""

	nav := []Link{
		Link{Name: "主页", Url: "HomeController.Get", Cur: "home"},
		Link{Name: "今日订单", Url: "HomeController.Today", Cur: "today"},
	}

	sess_nickname := c.GetSession("nickname")
	sess_isadmin := c.GetSession("isadmin")

	var adminCode int8 = 1
	if sess_isadmin == adminCode {
		//商家管理展开选择
		shopManageItems := []Link{
			Link{Name: "商家管理", Url: "admin/ShopController.Index", Cur: "shopmanage"},
			Link{Name: "用户管理", Url: "admin/UserController.Index", Cur: "todayshop"},
		}
		nav = append(nav, Link{Name: "商家管理", Url: "dropdown", Cur: "shop", Items: shopManageItems})
	}

	if sess_nickname != nil {
		nav = append(nav, Link{Name: "历史订单", Url: "UserController.History", Cur: "history"})
		nav = append(nav, Link{Name: strings.Join([]string{"退出(", sess_nickname.(string), ")"}, ""), Url: "UserController.Logout", Cur: "logout"})
	} else {
		nav = append(nav, Link{Name: "注册", Url: "UserController.Register", Cur: "register"})
		nav = append(nav, Link{Name: "登录", Url: "UserController.Login", Cur: "login"})
	}

	c.Data["nav"] = nav
}

/**
 * 显示错误提示
 * @param  {[type]} c *BaseController) showmsg(title string, content string [description]
 * @return {[type]}   [description]
 */
func (c *BaseController) showmsg(title string, content string) {
	c.Data["title"] = title
	c.Data["content"] = content
	c.Layout = "layout/main.tpl"
	c.TplNames = "error/user_error.tpl"
	c.Render()
	c.StopRun()
}

/**
 * 初始化分页
 */
func (this *BaseController) setPager(per int, nums int64) *lib.Pager {
	p := lib.NewPager(this.Ctx.Request, per, nums)
	this.Data["pager"] = p
	return p
}

/**
 * 设置语言包
 */
func (this *BaseController) settingLocales() {
	// load locales with locale_LANG.ini files
	langTypes = []langType{
		langType{Lang:"en-US", Name:"English"},
		langType{Lang:"zh-CN", Name:"中文"},
	}

	for _, langT := range langTypes {
		lang := langT.Lang

	    if i18n.IsExist(lang) {
	    	continue
	    }
	    if err := i18n.SetMessage(lang, "conf/i18n/"+lang+".ini"); err != nil {
	        beego.Error("Fail to set message file: " + err.Error())
	        return
	    }
	}
}

/**
 * 设定当前语言
 */
func (this *BaseController) setLangVer() bool {

    hasCookie := false

    // 参数传入语言类型
    lang := this.Input().Get("lang")

    // 丛cokkie获取语言类型
    if len(lang) == 0 {
        lang = this.Ctx.GetCookie("lang")
        hasCookie = true
    }

    // 如果语言不在设定内，重置语言类型
    if !i18n.IsExist(lang) {
        lang = ""
        hasCookie = false
    }

    // 从 'Accept-Language' 响应头部获取已设定的语言类型
    if len(lang) == 0 {
        al := this.Ctx.Request.Header.Get("Accept-Language")
        if len(al) > 4 {
            al = al[:5] // Only compare first 5 letters.
            if i18n.IsExist(al) {
                lang = al
            }
        }
    }

    // 无法获取最配置文件默认
    if len(lang) == 0 {
        lang = beego.AppConfig.String("lang")
    }

    // 设置cookie
    if !hasCookie {
        this.Ctx.SetCookie("lang", lang, 1<<31-1, "/")
    }

    this.Data["Lang"] = lang
    return true
}