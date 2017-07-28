/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:569929309

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:569929309

	2017.7 by hunterhug
*/
// 应用主函数包
package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/hunterhug/AmazonBigSpiderWeb/controllers"
	. "github.com/hunterhug/AmazonBigSpiderWeb/lib"
	"github.com/hunterhug/AmazonBigSpiderWeb/models"
	"github.com/hunterhug/AmazonBigSpiderWeb/routers"
	"mime"
	"strings"
)

// 国际化语言数组
var langTypes []string

// 加载、初始化国际化
func init() {
	langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

	for _, lang := range langTypes {
		beego.Trace("加载语言: " + lang)
		if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
			beego.Error("加载语言文件失败:", err)
			return
		}
	}

	// 添加映射
	beego.Trace("添加i18n函数映射")
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Trace("添加json格式化函数映射")
	beego.AddFuncMap("stringsToJson", StringsToJson)
	mime.AddExtensionType(".css", "text/css")

	// 模型初始化
	beego.Trace("模型初始化")
	models.Run()

	beego.Trace("路由开始")
	routers.Run()

	beego.Trace("错误模板开启")
	beego.ErrorController(&controllers.ErrorController{})
}

func main() {
	beego.Trace("监听开始")
	beego.Run()
}
