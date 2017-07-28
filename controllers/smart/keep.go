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
	"encoding/csv"
	"github.com/astaxie/beego/orm"
	"os"
	"strconv"
	"strings"
	//"fmt"
	"github.com/astaxie/beego"
)

type KeepController struct {
	baseController
}

func (this *KeepController) Index() {
	this.Layout = this.GetTemplate() + "/base/layout.html"
	this.TplName = this.GetTemplate() + "/back/keep.html"

}

func (this *KeepController) Query() {
	DB := orm.NewOrm()
	err := DB.Using("dbback")
	if err != nil {
		beego.Error("dbback err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	page, _ := this.GetInt("page", 1)
	rows, _ := this.GetInt("rows", 30)
	date := this.GetString("datename")
	date = strings.Replace(date, "-", "", -1)
	start := (page - 1) * rows
	num := 0
	var maps []orm.Params
	if date == "" {
		date = "20161101"
	}
	asin := this.GetString("asin")
	if asin == "" {
		dudu := "SELECT * FROM stock_info where createtime like '" + date + "%' limit " + strconv.Itoa(start) + "," + strconv.Itoa(rows) + ";"
		//fmt.Println(dudu)
		DB.Raw(dudu).Values(&maps)

		dudu1 := "SELECT count(*) as num FROM stock_info where createtime like '" + date + "%';"

		DB.Raw(dudu1).QueryRow(&num)
	} else {
		dudu := "SELECT * FROM stock_info where asin='" + asin + "' limit " + strconv.Itoa(start) + "," + strconv.Itoa(rows) + ";"

		DB.Raw(dudu).Values(&maps)

		dudu1 := "SELECT count(*) as num FROM stock_info where asin='" + asin + "';"

		DB.Raw(dudu1).QueryRow(&num)
	}

	if len(maps) == 0 {
		this.Data["json"] = &map[string]interface{}{"total": num, "rows": []interface{}{}}
	} else {
		this.Data["json"] = &map[string]interface{}{"total": num, "rows": &maps}
	}
	this.ServeJSON()
}

func (this *KeepController) Export() {
	asin := this.GetString("asin")
	if asin == "" {
		this.Rsp(false, "asin为空")
	}
	DB := orm.NewOrm()
	err := DB.Using("dbback")
	if err != nil {
		beego.Error("dbback err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	dudu := "SELECT * FROM stock_info where asin=?"
	var maps []orm.Params
	num, err := DB.Raw(dudu, asin).Values(&maps)
	if num == 0 || err != nil {
		this.Rsp(false, "asin找不到或查询结果为空")
	}
	filename := asin
	f, err := os.Create("file/data/asin-" + filename + ".csv")
	if err != nil {
		this.Rsp(false, "Excel创建出错")
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)
	//w.Write([]string{"标题", "商品链接", "价格", "小类名", "小类链接", "小类排名", "大类名", "真实大类名", "大类排名", "ReviewNum", "ReviewScore", "图片链接", "状态", "列表抓取时间", "详情抓取时间"})
	w.Write([]string{"Asin", "库存", "获取时间"})

	for _, i := range maps {
		temp := map[string]string{}
		for index, j := range i {
			innertemp := ""
			if j == nil {
				innertemp = " "
			} else {
				switch j.(type) { //多选语句switch
				case string:
					//是字符时做的事情
					innertemp = j.(string)
				case int:
					innertemp = strconv.Itoa(j.(int))
				}
			}
			temp[index] = innertemp
		}

		w.Write([]string{temp["asin"], temp["stock"], temp["createtime"]})

	}
	w.Flush()

	this.Rsp(true, "/file/data/asin-"+filename+".csv")
}
