package admin

import (
	//"fmt"
	//"strings"
	//"strconv"
	//"github.com/astaxie/beego"
	"github.com/dreamzml/canku/controllers"
)

type AdminbaseController struct {
	controllers.BaseController
}


/**
 * 页面载入之前
 * @param  {[type]} c *BaseController) Prepare( [description]
 * @return {[type]}   [description]
 */
func (this *AdminbaseController) Prepare() {
	//预设layout布局模板信息
	this.PrepareLayout()

	this.Data["subCur"] = "--";
}

