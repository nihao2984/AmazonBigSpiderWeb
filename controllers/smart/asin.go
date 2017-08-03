/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:459527502

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:459527502

	2017.7 by hunterhug
*/
package smart

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hunterhug/GoSpider/util"
	"strconv"
	"strings"
)

type AsinController struct {
	baseController
}

func (this *AsinController) Index() {
	DB := orm.NewOrm()
	err := DB.Using("usabasicdb")
	if err != nil {
		beego.Error("usabasicdb err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	var categorys []orm.Params
	DB.Raw("SELECT bigpname as Bigpname,id FROM smart_category where pid=0 group by bigpname,id").Values(&categorys)
	this.Data["category"] = &categorys
	this.Layout = this.GetTemplate() + "/base/layout.html"
	this.TplName = this.GetTemplate() + "/asin/list.html"

}

func (this *AsinController) Query() {
	DB := orm.NewOrm()
	err := DB.Using("usabasicdb")
	if err != nil {
		beego.Error("usabasicdb err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	asin := this.GetString("asin")
	num := 0
	var maps []orm.Params
	if asin == "" {
		page, _ := this.GetInt("page", 1)
		rows, _ := this.GetInt("rows", 30)
		date := this.GetString("datename")
		date = strings.Replace(date, "-", "", -1)
		isvalid, _ := this.GetInt("isvalid", 2)
		bigname := this.GetString("bigname")
		start := (page - 1) * rows
		where := []string{}
		wheresql := ""
		if date == "" {
		} else {
			where = append(where, `updatetime like "`+date+`%"`)
		}
		if bigname == "" {
		} else {
			where = append(where, `category="`+bigname+`"`)
		}

		if isvalid == 1 || isvalid == 0 {
			where = append(where, `isvalid=`+util.IS(isvalid))
		}
		if len(where) == 0 {

		} else {
			wheresql = strings.Join(where, " and ")
			wheresql = "where " + wheresql
		}
		dudu := "SELECT * FROM smart_asin " + wheresql + " order by updatetime limit " + strconv.Itoa(start) + "," + strconv.Itoa(rows) + ";"
		DB.Raw(dudu).Values(&maps)

		dudu1 := "SELECT count(*) as num FROM smart_asin " + wheresql + ";"

		DB.Raw(dudu1).QueryRow(&num)
	} else {
		dudu := "SELECT * FROM smart_asin where id=?;"
		DB.Raw(dudu, asin).Values(&maps)
		dudu1 := "SELECT count(*) as num FROM smart_asin where id=?;"
		DB.Raw(dudu1, asin).QueryRow(&num)
		num = 1
	}
	if len(maps) == 0 {
		this.Data["json"] = &map[string]interface{}{"total": num, "rows": []interface{}{}}
	} else {
		this.Data["json"] = &map[string]interface{}{"total": num, "rows": &maps}
	}
	this.ServeJSON()
}
