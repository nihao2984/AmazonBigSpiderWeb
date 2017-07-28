package routers

import "github.com/astaxie/beego"

 import (
// 	"github.com/astaxie/beego"
// 	"github.com/hunterhug/AmazonBigSpiderWeb/controllers/admin"
 	"github.com/hunterhug/AmazonBigSpiderWeb/controllers/home"

	 //"github.com/hunterhug/AmazonBigSpiderWeb/controllers/admin/rbac"
 )

func homerouter() {

// 	//前台路由
	beego.Router("/home", &home.MainController{}, "*:Index")
 	//beego.Router("/", &home.MainController{}, "*:Index")
 	beego.Router("/home/:id/", &home.MainController{}, "*:Category")
	beego.Router("/home/:cid/:id/", &home.MainController{}, "*:Paper")
// 	beego.Router("/404.html", &home.MainController{}, "*:Go404")
// 	beego.Router("/index:page:int.html", &home.MainController{}, "*:Index")

// 	beego.Router("/article/:id:int", &home.MainController{}, "*:Show")      //ID访问
// 	beego.Router("/article/:urlname(.+)", &home.MainController{}, "*:Show") //别名访问文章

// 	beego.Router("/category/:name(.+?)", &home.MainController{}, "*:Category")
// 	beego.Router("/category/:name(.+?)/page/:page:int", &home.MainController{}, "*:Category")

// 	beego.Router("/life:page:int.html", &home.MainController{}, "*:homeList")
// 	beego.Router("/life.html", &home.MainController{}, "*:homeList")

// 	beego.Router("/mood.html", &home.MainController{}, "*:Mood")
// 	beego.Router("/mood:page:int.html", &home.MainController{}, "*:Mood")

// 	//照片展示
// 	beego.Router("/photo.html", &home.MainController{}, "*:Photo")
// 	beego.Router("/photo:page:int.html", &home.MainController{}, "*:Photo")

// 	//相册展示
// 	beego.Router("/album.html", &home.MainController{}, "*:Album")
// 	beego.Router("/album:page:int.html", &home.MainController{}, "*:Album")

// 	beego.Router("/book.html", &home.MainController{}, "*:Book")
// 	beego.Router("/:urlname(.+)", &home.MainController{}, "*:Show") //别名访问

 }
