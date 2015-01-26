package admin

import (
	//"fmt"
	//m "github.com/dreamzml/canku/models"
)

type ShopController struct {
	AdminbaseController
}

func (this *ShopController) Index() {
	this.Data["subCur"] = "shop_index";
//fmt.Printf(this.Data["cur"])
	this.TplNames = "admin/shop/index.tpl"
	this.Render()
}