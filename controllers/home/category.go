package home

import (
	"github.com/astaxie/beego/orm"
	"github.com/hunterhug/AmazonBigSpiderWeb/models/blog"
	"github.com/hunterhug/AmazonBigSpiderWeb/lib"
	//"github.com/astaxie/beego"
)

func (this *MainController) Category() {
	id := this.Ctx.Input.Param(":id")
	types := 0
	err, category := getcategory(id)
	if err != nil {
		err, category = getalbum(id)
		if err != nil {
			this.Rsp(false, "没有这个分类啊，哥哥")
		} else {
			types = 1
		}
	}

	//本大爷
	this.Data["thiscategory"] = category

	//附录爸爸
	cid := category["Pid"]
	father := new(blog.Category)
	father.Id = cid.(int64)
	err1:=father.Read()
	if err1!=nil{

	}else {
		this.Data["father"] = father
	}
	//附录儿子
	son := []orm.Params{}
	new(blog.Category).Query().Filter("Pid", category["Id"].(int64)).Filter("Status", 1).OrderBy("-Sort", "Createtime").Values(&son, "Title")
	this.Data["son"] = son

	//文章儿子
	var limit int64=8
	papers := []orm.Params{}
	query1:=new(blog.Paper).Query().Filter("Cid",category["Id"].(int64)).Filter("Status", 1).OrderBy("-Istop", "Createtime")
	page,_:=this.GetInt64("page",1)
	if page<=0{
		page=1
	}
	//总数
	count,_:=query1.Count()

	temp:=new(lib.Pager)
	temp.Page=page
	temp.Pagesize=limit
	temp.Totalnum=count
	temp.Urlpath="/"+category["Title"].(string)
	//beego.Trace("Dddd"+temp.ToString())
	this.Data["nums"]=temp.ToString()
	query1.Limit(limit,(page-1)*limit).Values(&papers)

	this.Data["papers"] = papers

	////图片轮转
	//roll := new(blog.Roll)
	//rolls := []orm.Params{}
	//roll.Query().Filter("Status", 1).OrderBy("-Sort", "Createtime").Values(&rolls)
	//this.Data["roll"] = rolls

	if types == 0 {
		this.TplName = this.GetTemplate() + "/category.html"
	} else {
		this.TplName = this.GetTemplate() + "/album.html"
	}
}
