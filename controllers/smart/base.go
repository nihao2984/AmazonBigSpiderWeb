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
package smart

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/orm"
	"github.com/beego/i18n"
)

type baseController struct {
	beego.Controller
	i18n.Locale
}


func (this *baseController) Prepare() {
	this.Lang = ""

	al := this.Ctx.Request.Header.Get("Accept-Language")
	if len(al) > 4 {
		al = al[:5]
		if i18n.IsExist(al) {
			this.Lang = al
		}
	}

	if len(this.Lang) == 0 {
		this.Lang = "zh-CN"
	}

	this.Data["Lang"] = this.Lang


}

// 获取模板位置
func (this *baseController) GetTemplate() string {
	templatetype := beego.AppConfig.String("smart_temlate")
	if templatetype == "" {
		templatetype = "default"
	}
	return templatetype
}

func (this *baseController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
	this.StopRun()
}
