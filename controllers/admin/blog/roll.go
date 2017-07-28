package blog

import (
	// "fmt"
	//"github.com/astaxie/beego"
	. "github.com/hunterhug/AmazonBigSpiderWeb/lib"
	//"github.com/hunterhug/AmazonBigSpiderWeb/models/admin"
	"github.com/astaxie/beego/orm"
	"github.com/hunterhug/AmazonBigSpiderWeb/models/blog"
)

type RollController struct {
	baseController
}

func (this *RollController) Index() {
	roll := new(blog.Roll)
	rolls := []orm.Params{}
	roll.Query().OrderBy("-Status", "-Sort", "Createtime").Values(&rolls)
	this.Data["roll"] = rolls
	this.TplName = this.GetTemplate() + "/blog/rollindex.html"
}

func (this *RollController) AddRoll() {
	roll := new(blog.Roll)
	num, _ := roll.Query().Count()
	if num >= 6 {
		this.Rsp(false, "图片不能超过六张")
	}

	title := this.GetString("title", "")
	content := this.GetString("content", "")
	url := this.GetString("url", "")
	sort, _ := this.GetInt64("sort", 0)
	photo := this.GetString("photo")
	roll.Content = content
	roll.Createtime = GetTime()
	roll.Photo = photo
	roll.Status = 0
	roll.Title = title
	roll.Url = url
	roll.Sort = sort
	roll.View = 0
	err := roll.Insert()
	if err != nil {
		this.Rsp(false, err.Error())
	} else {
		this.Rsp(true, "增加成功")
	}
}

func (this *RollController) UpdateRoll() {
	id, _ := this.GetInt64("id", -1)
	if id == -1 {
		this.Rsp(false, "id参数问题")
	}
	roll := new(blog.Roll)
	roll.Id = id
	err := roll.Read()
	if err != nil {
		this.Rsp(false, "找不到")
	}
	if this.IsAjax() {
		status, _ := this.GetInt64("status")
		small := this.GetString("small", "")
		if small == "open" {
			roll.Status=1
		} else if small == "close" {
			roll.Status = 0
		} else {
			content := this.GetString("content")
			url := this.GetString("url")
			sort, _ := this.GetInt64("sort")
			photo := this.GetString("photo", "")
			title := this.GetString("title")
			roll.Title = title
			roll.Status = status
			roll.Content = content
			roll.Updatetime = GetTime()
			roll.Url = url
			roll.Sort = sort
			if photo == "" {

			} else {
				roll.Photo = photo
			}
		}
		err := roll.Update()
		if err != nil {
			this.Rsp(false, err.Error())

		} else {
			this.Rsp(true, "修改成功")
		}
	} else {
		this.Data["roll"] = roll
		this.TplName = this.GetTemplate() + "/blog/rollupdate.html"

	}
}
func (this *RollController) DeleteRoll() {
	id, _ := this.GetInt64("id", -1)
	if id == -1 {
		this.Rsp(false, "id参数问题")
	}
	roll := new(blog.Roll)
	roll.Id = id
	err := roll.Delete()
	if err != nil {
		this.Rsp(false, err.Error())
	} else {
		this.Rsp(true, "删除成功")
	}
}
