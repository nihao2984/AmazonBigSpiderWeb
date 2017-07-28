package blog

import (
	// "fmt"
	"github.com/astaxie/beego/orm"
	. "github.com/hunterhug/AmazonBigSpiderWeb/lib"
	"github.com/hunterhug/AmazonBigSpiderWeb/models/admin"
	"github.com/hunterhug/AmazonBigSpiderWeb/models/blog"
	//"github.com/astaxie/beego"
)

type AlbumController struct {
	baseController
}


func (this *AlbumController) Index() {
	/*
	   status: status,
	     mulu: mulu
	*/
	category := new(blog.Category)
	categorys := []orm.Params{}
	if this.IsAjax() {
		status, _ := this.GetInt64("status", 0)
		mulu, _ := this.GetInt64("mulu", 0)
		// beego.Trace(status, mulu)
		if status == 0 {
			category.Query().Filter("Pid", mulu).Filter("Siteid", beautyid).Filter("Type", albumtype).OrderBy("-Sort", "Createtime").Values(&categorys)
		} else {
			category.Query().Filter("Pid", mulu).Filter("Status", status).Filter("Siteid", beautyid).Filter("Type", albumtype).OrderBy("-Sort", "Createtime").Values(&categorys)
		}
		count := len(categorys)
		// beego.Trace("%v", categorys)
		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &categorys}
		this.ServeJSON()
		return
	} else {
		category.Query().Filter("Pid", 0).Filter("Siteid", beautyid).Filter("Type", albumtype).OrderBy("-Sort", "Createtime").Values(&categorys)
		this.Data["category"] = &categorys
		this.Layout = this.GetTemplate() + "/album/layout.html"
		this.TplName = this.GetTemplate() + "/album/listcate.html"
	}
}

func (this *AlbumController) AddCategory() {
	user := this.GetSession("userinfo")
	isajax, _ := this.GetInt("isajax", 0)
	if isajax == 1 {
		status := false
		message := ""
		if user == nil {
			message = "session失效，请重新进入后台首页"
		} else {
			category := new(blog.Category)
			category.Createtime = GetTime()
			category.Username = user.(admin.User).Username
			category.Title = this.GetString("title", "")
			category.Pid, _ = this.GetInt64("mulu", 0)
			category.Sort, _ = this.GetInt64("order", 0)
			category.Status, _ = this.GetInt64("status", 2)
			category.Content = this.GetString("content", "")
			category.Image = this.GetString("photo", "")
			category.Siteid = beautyid
			category.Type = albumtype
			err := category.Insert()
			if err != nil {
				message = err.Error()
			} else {
				status = true
				message = "增加成功"
			}
		}
		this.Rsp(status, message)
	} else {
		category := new(blog.Category)
		categorys := []orm.Params{}
		category.Query().Filter("Pid", 0).Filter("Siteid", beautyid).Filter("Type", albumtype).OrderBy("-Sort", "Createtime").Values(&categorys)
		this.Data["category"] = &categorys
		this.TplName = this.GetTemplate() + "/album/addcate.html"
	}
}

func (this *AlbumController) UpdateCategory() {
	user := this.GetSession("userinfo")
	if user == nil {
		this.Rsp(false, "session失效，请重新进入后台首页")
	}
	small, _ := this.GetInt64("small", 0)
	id, _ := this.GetInt64("id", 0)
	//小更改
	if small == 1 {
		status, _ := this.GetInt64("status", 0)
		if id == 0 || status == 0 {
			this.Rsp(false, "有问题")
		} else {
			category := new(blog.Category)
			category.Id = id
			category.Status = status
			category.Updatetime = GetTime()
			err := category.Update("Status", "Updatetime")
			if err != nil {
				this.Rsp(false, "更新失败")
			} else {
				this.Rsp(true, "更新成功")
			}
		}
	} else {
		if this.IsAjax() {
			//大更改
			thiscategory := new(blog.Category)
			thiscategory.Id = id
			thiscategory.Username = user.(admin.User).Username
			thiscategory.Title = this.GetString("title", "")
			thiscategory.Pid, _ = this.GetInt64("mulu", 0)
			thiscategory.Sort, _ = this.GetInt64("order", 0)
			thiscategory.Status, _ = this.GetInt64("status", 2)
			thiscategory.Content = this.GetString("content", "")
			thiscategory.Updatetime = GetTime()
			//不存在则不改图片
			photo := this.GetString("photo", "")
			//beego.Trace("图片：" + photo)
			var err error
			if photo != "" {
				thiscategory.Image = photo
				err = thiscategory.Update("Username", "Title", "Pid", "Sort", "Status", "Content", "Updatetime", "Image")
			}else {
				err = thiscategory.Update("Username", "Title", "Pid", "Sort", "Status", "Content", "Updatetime")
				//beego.Trace("空图片：" + photo)
			}
			if err != nil {
				this.Rsp(false, err.Error())
			}else {
				this.Rsp(true, "更改成功")
			}
		} else {
			if id == 0 {
				this.Rsp(false, "没有id参数")
			}
			//显示更改页面
			thiscategory := new(blog.Category)
			thiscategory.Id = id
			err := thiscategory.Read()
			if err != nil {
				this.Rsp(false, "不存在该目录或者数据库出错")
			}
			this.Data["thiscategory"] = thiscategory

			category := new(blog.Category)
			categorys := []orm.Params{}
			category.Query().Exclude("Id", id).Filter("Pid", 0).Filter("Siteid", beautyid).Filter("Type", albumtype).OrderBy("-Sort", "Createtime").Values(&categorys)
			this.Data["category"] = &categorys

			this.TplName = this.GetTemplate() + "/album/updatecate.html"
		}

	}
}

func (this *AlbumController) DeleteCategory() {
	category := new(blog.Category)
	id, err := this.GetInt64("id", 0)
	if err != nil || id == 0 {
		this.Rsp(false, "出现错误")
	}
	num, err := category.Query().Filter("Id", id).Filter("Siteid", beautyid).Filter("Type", albumtype).Count()
	if err != nil {
		this.Rsp(false, err.Error())
	}else if num == 0 {
		this.Rsp(false, "找不到该目录")
	}else {
		paper := new(blog.Paper)
		num1, err1 := paper.Query().Filter("Cid", id).Count()
		if num1 != 0 {
			this.Rsp(false, "目录下有小东西")
		}else if err1 != nil {
			this.Rsp(false, err1.Error())
		}else {
			num2, err2 := category.Query().Filter("Pid", id).Count()
			if err2 != nil {
				this.Rsp(false, err2.Error())
			}else if num2 != 0 {
				this.Rsp(false, "目录下有目录")
			}else {
				category.Id = id
				err3 := category.Delete()
				if err3 != nil {
					this.Rsp(false, err2.Error())
				}else {
					this.Rsp(true, "删除成功")
				}
			}
		}
	}
}
