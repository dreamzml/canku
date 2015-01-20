package controllers

type HomeController struct {
	BaseController
}

/**
* 首页
 */
func (this *HomeController) Get() {
	nickname := this.GetSession("nickname")
	if nickname != nil {
		this.Data["nickname"] = nickname
	}
	//声明导航焦点
	this.Data["cur"] = "home"
	this.TplNames = "home/index.tpl"
	this.Render()
}

/**
 * 今日订单
 * @param  {[type]} this *HomeController) Today( [description]
 * @return {[type]}      [description]
 */
func (this *HomeController) Today() {

}
