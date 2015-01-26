package admin

import (
	//"fmt"
	"time"
	//"reflect"
	m "github.com/dreamzml/canku/models"
)

type UserController struct {
	AdminbaseController
}

/**
* 会员管理
*/
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
	 	
	//输出数据类型
	//fmt.Printf("time format type is : %s \n\n", reflect.TypeOf(users))

	for _,d := range users {
	 	d["Data"] = time.Unix(d["Lastlogin"].(int64), 0).Format("2006-01-02 15:04:05")
	 	//d["Id"] = d["Id"].(string)
	}

	this.Data["count"] = count
	this.Data["rows"] = &users

	this.Data["subCur"] = "user_index"
	this.TplNames = "admin/user/index.tpl"
	this.Render()
}


func (this *UserController) Update() {
	// u := m.User{}
	// if err := this.ParseForm(&u); err != nil {
	// 	//handle error
	// 	this.Rsp(false, err.Error())
	// 	return
	// }
	// id, err := m.UpdateUser(&u)
	// if err == nil && id > 0 {
	// 	this.Rsp(true, "Success")
	// 	return
	// } else {
	// 	this.Rsp(false, err.Error())
	// 	return
	// }

}

func (this *UserController) Delete() {

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
