package admin

import (
	//"fmt"
	"time"
	"strconv"

	//"reflect"
	m "github.com/dreamzml/canku/models"
	. "github.com/dreamzml/canku/lib"
)

type UserController struct {
	AdminbaseController
}

/**
* 会员管理
*/
func (this *UserController) Index() {

	users, count := m.Getuserlist(this.Page, this.Page_size, this.Sort)

	//输出数据类型
	//fmt.Printf("time format type is : %s \n\n", reflect.TypeOf(users))
	for _,d := range users {
	 	d["Lastlogin_data"] = time.Unix(d["Lastlogin"].(int64), 0).Format("2006-01-02 15:04:05")
	 	d["Id_str"] = strconv.Itoa(int(d["Id"].(int64)))
	}

	this.Data["count"] = count
	this.Data["rows"] = &users


	//this.showmsg("500","test")
	this.Data["pager"] = this.Buildpager(count)

    this.Data["iamgeSource"] = TexToImg("杨")
    this.Data["iamgeSource2"] = TexToImg("明")
    this.Data["iamgeSource3"] = TexToImg("傲")
    this.Data["iamgeSource4"] = TexToImg(" D")
    this.Data["iamgeSource5"] = TexToImg(" 8")
    this.Data["iamgeSource6"] = TexToImg(" K")



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
 