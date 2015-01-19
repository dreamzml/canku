package controllers

type HomeController struct {
	BaseController
}

/**
* 首页
*/
func (this *HomeController) Get() {
	//声明导航焦点
	this.Data["cur"] = "home";
	this.TplNames = "home/index.tpl" 

}
