package blog

import (
	//"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	. "github.com/hunterhug/AmazonBigSpiderWeb/lib"
	"github.com/hunterhug/AmazonBigSpiderWeb/models/blog"
)

type PaperController struct {
	baseController
}

func (this *PaperController) Index() {

	if this.IsAjax() {
		page, _ := this.GetInt64("page", 1)
		rows, _ := this.GetInt64("rows", 10)
		start := (page - 1) * rows
		paper := new(blog.Paper)
		papers := []orm.Params{}
		//默认状态为关闭
		status, _ := this.GetInt64("status", -1)
		//分类可以为空，列出所有
		cid, _ := this.GetInt64("cid", -1)
		top, _ := this.GetInt64("top", -1)
		roll, _ := this.GetInt64("roll", -1)
		q := paper.Query().Filter("Type",0)
		if status != -1 {
			q = q.Filter("Status", status)
		}else {
			q = q.Filter("Status__lt", 2)
		}
		if cid != -1 {
			q = q.Filter("Cid", cid)
		}
		if top != -1 {
			q = q.Filter("Istop", top)
		}
		if roll != -1 {
			q = q.Filter("Isroll", roll)
		}
		q.OrderBy("-Sort", "-Istop", "-View", "Createtime").Limit(rows, start).Values(&papers)
		for _, p := range papers {

			category := new(blog.Category)
			category.Id = (p["Cid"]).(int64)
			//beego.Trace(category.Id)
			err := category.Read("Id")
			if err != nil {
				p["Cid"] = "空"
			} else {
				p["Cid"] = category.Title
			}
		}
		count, _ := q.Count()

		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &papers}
		this.ServeJSON()
	} else {
		//文章列表首页
		category := new(blog.Category)
		categorys := []orm.Params{}
		//查询条件：缀美文章类型，一级
		category.Query().Filter("Pid", 0).Filter("Siteid", beautyid).Filter("Type", blogtype).OrderBy("-Sort", "Createtime").Values(&categorys, "Id", "Title")
		for _, cate := range categorys {
			//二级
			son := []orm.Params{}
			category.Query().Filter("Pid", cate["Id"]).OrderBy("-Sort", "Createtime").Values(&son, "Id", "Title")
			cate["Son"] = son
		}
		this.Data["category"] = &categorys
		this.Layout = this.GetTemplate() + "/blog/layout.html"
		this.TplName = this.GetTemplate() + "/blog/listpaper.html"
	}
}

func (this *PaperController) AddPaper() {

	if this.IsAjax() {
		title := this.GetString("title")
		content := this.GetString("content")
		descontent := this.GetString("descontent")
		author := this.GetString("author")
		cid, _ := this.GetInt64("cid", 0)
		if cid == 0 {
			this.Rsp(false, "目录选择有问题")
		}
		sort, _ := this.GetInt64("order", 0)
		view, _ := this.GetInt64("view", 0)
		status, _ := this.GetInt64("status", 0)
		top, _ := this.GetInt64("top", 0)
		roll, _ := this.GetInt64("roll", 0)
		rollpath := this.GetString("rollpath")
		photo := this.GetString("photo")
		paper := new(blog.Paper)
		paper.Title = title
		paper.Author = author
		paper.Content = content
		paper.Descontent = descontent
		paper.Cid = cid
		paper.Sort = sort
		paper.View = view
		paper.Status = status
		paper.Istop = top
		paper.Isroll = roll
		paper.Rollpath = rollpath
		paper.Photo = photo
		paper.Createtime = GetTime()
		paper.Type=0
		err := paper.Insert()
		if err != nil {
			this.Rsp(false, err.Error())
		} else {
			this.Rsp(true, "增加成功")
		}
	} else {
		//文章列表首页
		category := new(blog.Category)
		categorys := []orm.Params{}
		//查询条件：缀美文章类型，一级
		category.Query().Filter("Pid", 0).Filter("Siteid", beautyid).Filter("Type", blogtype).OrderBy("-Sort", "Createtime").Values(&categorys, "Id", "Title")
		for _, cate := range categorys {
			//二级
			son := []orm.Params{}
			category.Query().Filter("Pid", cate["Id"]).OrderBy("-Sort", "Createtime").Values(&son, "Id", "Title")
			cate["Son"] = son
		}
		this.Data["category"] = &categorys
		this.TplName = this.GetTemplate() + "/blog/addpaper.html"
	}
}

func (this *PaperController) UpdatePaper() {
	//user := this.GetSession("userinfo")
	//if user == nil {
	//	this.Rsp(false, "session失效，请重新进入后台首页")
	//}
	id, _ := this.GetInt64("id", 0)
	if id == 0 {
		this.Rsp(false, "文章id参数问题")
	}
	if this.IsAjax() {
		small := this.GetString("small")
		if small == "1" {
			status, _ := this.GetInt64("status", 0)
			paper := new(blog.Paper)
			paper.Id = id
			paper.Status = status
			paper.Updatetime = GetTime()
			err := paper.Update("Status", "Updatetime")
			if err != nil {
				this.Rsp(false, err.Error())
			}else {
				this.Rsp(true, "更改状态成功")
			}
			this.StopRun()
		}
		title := this.GetString("title")
		content := this.GetString("content")
		descontent := this.GetString("descontent")
		author := this.GetString("author")
		cid, _ := this.GetInt64("cid", 0)
		if cid == 0 {
			this.Rsp(false, "目录选择有问题")
		}
		sort, _ := this.GetInt64("order", 0)
		view, _ := this.GetInt64("view", 0)
		status, _ := this.GetInt64("status", 0)
		top, _ := this.GetInt64("top", 0)
		roll, _ := this.GetInt64("roll", 0)
		rollpath := this.GetString("rollpath")
		photo := this.GetString("photo")
		paper := new(blog.Paper)
		paper.Id = id
		paper.Title = title
		paper.Author = author
		paper.Content = content
		paper.Descontent = descontent
		paper.Cid = cid
		paper.Sort = sort
		paper.View = view
		paper.Status = status
		paper.Istop = top
		paper.Isroll = roll
		paper.Rollpath = rollpath
		paper.Updatetime = GetTime()
		var err error
		if photo != "" {
			paper.Photo = photo
			err = paper.Update("Title", "Content", "Descontent", "Author", "View", "Sort", "Istop", "Isroll", "Rollpath", "Cid", "Status", "Updatetime", "Photo")
		}else {
			err = paper.Update("Title", "Content", "Descontent", "Author", "View", "Sort", "Istop", "Isroll", "Rollpath", "Cid", "Status", "Updatetime")
		}
		if err != nil {
			this.Rsp(false, err.Error())
		} else {
			this.Rsp(true, "修改成功")
		}
	} else {

		//文章列表首页
		category := new(blog.Category)
		categorys := []orm.Params{}
		//查询条件：缀美文章类型，一级
		category.Query().Filter("Pid", 0).Filter("Siteid", beautyid).Filter("Type", blogtype).OrderBy("-Sort", "Createtime").Values(&categorys, "Id", "Title")
		for _, cate := range categorys {
			//二级
			son := []orm.Params{}
			category.Query().Filter("Pid", cate["Id"]).OrderBy("-Sort", "Createtime").Values(&son, "Id", "Title")
			cate["Son"] = son
		}
		this.Data["category"] = &categorys

		if id == 0 {
			this.Rsp(false, "没有id参数")
			this.StopRun()
		}
		//显示更改页面
		thispaper := new(blog.Paper)
		thispaper.Id = id
		err := thispaper.Read()
		if err != nil {
			this.Rsp(false, "不存在该文章或者数据库出错")
			this.StopRun()
		}
		this.Data["thispaper"] = thispaper

		this.TplName = this.GetTemplate() + "/blog/updatepaper.html"
	}
}

func (this *PaperController) DeletePaper() {
	id, _ := this.GetInt64("id", -1)
	if id != -1 {
		paper := new(blog.Paper)
		paper.Id = id
		paper.Status = 2
		paper.Updatetime = GetTime()
		err := paper.Update("Status", "Updatetime")
		if err != nil {
			this.Rsp(false, err.Error())
		}else {
			this.Rsp(true, "送到回收站")
		}
	}else {
		this.Rsp(false, "id参数问题")
	}
}

func (this *PaperController) RealDelPaper() {
	id, _ := this.GetInt64("id", -1)
	if id != -1 {
		paper := new(blog.Paper)
		paper.Id = id
		err := paper.Delete()
		if err != nil {
			this.Rsp(false, err.Error())
		}else {
			this.Rsp(true, "成功删除")
		}
	}else {
		this.Rsp(false, "id参数问题")
	}
}

func (this *PaperController)  Rubbish() {

	if this.IsAjax() {
		page, _ := this.GetInt64("page", 1)
		rows, _ := this.GetInt64("rows", 10)
		start := (page - 1) * rows
		paper := new(blog.Paper)
		papers := []orm.Params{}
		//默认状态为关闭
		status, _ := this.GetInt64("status", 2)
		//分类可以为空，列出所有
		cid, _ := this.GetInt64("cid", -1)
		top, _ := this.GetInt64("top", -1)
		roll, _ := this.GetInt64("roll", -1)
		q := paper.Query()
		q = q.Filter("Status", status).Filter("Type",0)

		if cid != -1 {
			q = q.Filter("Cid", cid)
		}
		if top != -1 {
			q = q.Filter("Istop", top)
		}
		if roll != -1 {
			q = q.Filter("Isroll", roll)
		}
		q.OrderBy("-Sort", "-Istop", "-View", "Createtime").Limit(rows, start).Values(&papers)
		for _, p := range papers {

			category := new(blog.Category)
			category.Id = (p["Cid"]).(int64)
			//beego.Trace(category.Id)
			err := category.Read("Id")
			if err != nil {
				p["Cid"] = "空"
			} else {
				p["Cid"] = category.Title
			}
		}
		count, _ := q.Count()

		this.Data["json"] = &map[string]interface{}{"total": count, "rows": &papers}
		this.ServeJSON()
	} else {
		//文章列表首页
		category := new(blog.Category)
		categorys := []orm.Params{}
		//查询条件：缀美文章类型，一级
		category.Query().Filter("Pid", 0).Filter("Siteid", beautyid).Filter("Type", blogtype).OrderBy("-Sort", "Createtime").Values(&categorys, "Id", "Title")
		for _, cate := range categorys {
			//二级
			son := []orm.Params{}
			category.Query().Filter("Pid", cate["Id"]).OrderBy("-Sort", "Createtime").Values(&son, "Id", "Title")
			cate["Son"] = son
		}
		this.Data["category"] = &categorys
		this.Layout = this.GetTemplate() + "/blog/layout.html"
		this.TplName = this.GetTemplate() + "/blog/rubbish.html"
	}
}