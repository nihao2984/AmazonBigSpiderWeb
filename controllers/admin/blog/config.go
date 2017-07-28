package blog

import (
	// "fmt"
	//"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	//. "github.com/hunterhug/AmazonBigSpiderWeb/lib"
	//"github.com/hunterhug/AmazonBigSpiderWeb/models/admin"
	//"github.com/hunterhug/AmazonBigSpiderWeb/models/blog"
)
import "github.com/hunterhug/AmazonBigSpiderWeb/models/blog"

type ConfigController struct {
	baseController
}

func (this *ConfigController) IndexUser() {
	this.Rsp(true,"暂时不提供修改个人信息功能，请联系管理员")
}

func (this *ConfigController) IndexOption() {
	config:=new(blog.Config)
	config.Id=1
	config.Read()
	this.Data["config"]=config
	this.TplName=this.GetTemplate()+"/config/weboption.html"
}

func (this *ConfigController) UpdateUser() {
}
func (this *ConfigController) UpdateOption() {
	config:=new(blog.Config)
	config.Id=1
	config.Read()
	title:=this.GetString("title")
	content:=this.GetString("content")
	slogan:=this.GetString("slogan")
	address:=this.GetString("address")
	phone:=this.GetString("phone")
	webinfo:=this.GetString("webinfo")
	code1:=this.GetString("code1")
	code2:=this.GetString("code2")
	code3:=this.GetString("code3")
	photo:=this.GetString("photo")
	if photo==""{

	}else{
		config.Photo=photo
	}
	config.Title=title
	config.Content=content
	config.Slogan=slogan
	config.Address=address
	config.Phone=phone
	config.Webinfo=webinfo
	config.Code1=code1
	config.Code2=code2
	config.Code3=code3
	err:=config.Update()
	if err!=nil{
		this.Rsp(false,err.Error())
	}
	this.Rsp(true,"修改成功")
}