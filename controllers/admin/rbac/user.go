package rbac

import (
	m "github.com/hunterhug/AmazonBigSpiderWeb/models/admin"
)

type UserController struct {
	CommonController
}

func (this *UserController) Index() {
	if this.IsAjax() {
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

		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &users}
		this.ServeJSON()
		return
	} else {
		this.Layout = this.GetTemplate() + "/public/layout.html"
		this.TplName = this.GetTemplate() + "/rbac/user.html"
	}

}

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

//UpdateUserPasswd
func (this *UserController) UpdateUserPasswd() {
	u := m.User{}
	if err := this.ParseForm(&u); err != nil {
		//handle error
		this.Rsp(false, err.Error())
		return
	}
	id, err := m.UpdateUserPasswd(&u)
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}

}

func (this *UserController) DelUser() {
	// Id, _ := this.GetInt64("Id")
	// status, err := m.DelUserById(Id)
	// if err == nil && status > 0 {
	// 	this.Rsp(true, "Success")
	// 	return
	// } else {
	// 	this.Rsp(false, err.Error())
	// 	return
	// }

	this.Rsp(false, "本系统用户一添加则不允许删除")
	return
}
