package blog

import (
	"github.com/astaxie/beego"
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
	templatetype := beego.AppConfig.String("admin_template")
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
