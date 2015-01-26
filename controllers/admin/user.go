package admin

import (
	//"fmt"
	m "github.com/dreamzml/canku/models"
)

type UserController struct {
	AdminbaseController
}

func (this *UserController) Index() {
	page, _ := this.GetInt64("page")
	page_size, _ := this.GetInt64("rows")
	sort := this.GetString("sort")
	order := this.GetString("order")
	if len(order) > 0 {
		if order == "desc" {
			sort = "-" + sort
		}
	} else {
		sort = "Id"
	}
	users, count := m.Getuserlist(page, page_size, sort)
	// if this.IsAjax() {
	// 	this.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
	// 	this.ServeJson()
	// 	return
	// } else {
	// 	tree := this.GetTree()
	// 	this.Data["tree"] = &tree
	// 	this.Data["users"] = &users
	// 	if this.GetTemplatetype() != "easyui" {
	// 		this.Layout = this.GetTemplatetype() + "/public/layout.tpl"
	// 	}
	this.Data["rows"] = &users
	this.Data["count"] = count
	 	
	// }
	this.Data["subCur"] = "user_index"
	this.TplNames = "admin/user/index.tpl"
	this.Render()
}

/*
func (this *UserController) AddUser() {
	u := m.User{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.AddUser(&u)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *UserController) UpdateUser() {
	u := m.User{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.UpdateUser(&u)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *UserController) DelUser() {
	Id, _ := this.GetInt64("Id")
	status, err := m.DelUserById(Id)
	if err == nil && status > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}
*/
