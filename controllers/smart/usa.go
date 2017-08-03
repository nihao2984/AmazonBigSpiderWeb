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
	"encoding/csv"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hunterhug/GoSpider/util"
	"os"
	"strconv"
	"strings"
)

type UsaController struct {
	baseController
}

func (this *UsaController) Index() {
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
	this.TplName = this.GetTemplate() + "/usa/list.html"

}

func (this *UsaController) Query() {
	DB := orm.NewOrm()
	err := DB.Using("usadatadb")
	if err != nil {
		beego.Error("usadatadb err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	asin := this.GetString("asin")
	num := 0
	var maps []orm.Params
	date := this.GetString("datename")
	date = strings.Replace(date, "-", "", -1)
	name := this.GetString("name")
	name = strings.TrimSpace(name)
	page, _ := this.GetInt("page", 1)
	rows, _ := this.GetInt("rows", 30)
	start := (page - 1) * rows

	if name != "" {
		dudu := "SELECT * FROM `" + date + "`where name='" + name + "' limit " + strconv.Itoa(start) + "," + strconv.Itoa(rows) + ";"
		//fmt.Println(dudu)
		DB.Raw(dudu).Values(&maps)
		dudu1 := "SELECT count(*) FROM `" + date + "`where name='" + name + "'"
		DB.Raw(dudu1).QueryRow(&num)
		if len(maps) == 0 {
			this.Data["json"] = &map[string]interface{}{"total": num, "rows": []interface{}{}}
		} else {
			this.Data["json"] = &map[string]interface{}{"total": num, "rows": &maps}
		}
		this.ServeJSON()
		return
	}

	if date == "" {
		date = "20161101"
	}
	if asin == "" {

		bigname := this.GetString("bigname")
		iscatchi, _ := this.GetInt("iscatch", 2)
		if iscatchi > 2 || iscatchi < 0 {
			this.Rsp(false, "没毛病吧")
		}
		iscatch := util.IS(iscatchi)

		if bigname == "" {
			bigname = "all"
			//什么都没有
		}
		dudu := ""
		if bigname == "all" {
			if iscatchi == 2 {
				dudu = "SELECT * FROM `" + date + "` order by name,smallrank limit " + strconv.Itoa(start) + "," + strconv.Itoa(rows) + ";"
			} else {
				dudu = "SELECT * FROM `" + date + "`where iscatch=" + iscatch + " order by name,smallrank limit " + strconv.Itoa(start) + "," + strconv.Itoa(rows) + ";"
			}
		} else {
			if iscatchi == 2 {
				dudu = "SELECT * FROM `" + date + "` where bigname like ? order by rank limit " + strconv.Itoa(start) + "," + strconv.Itoa(rows) + ";"
			} else {
				dudu = "SELECT * FROM `" + date + "` where bigname like ? and iscatch=" + iscatch + " order by rank limit " + strconv.Itoa(start) + "," + strconv.Itoa(rows) + ";"

			}
		}
		if bigname != "all" {
			DB.Raw(dudu, bigname).Values(&maps)
		} else {
			DB.Raw(dudu).Values(&maps)
		}
		dudu1 := ""
		if bigname != "all" {
			if iscatchi == 2 {
				dudu1 = "SELECT count(*) as num FROM `" + date + "` where bigname like ?;"
			} else {
				dudu1 = "SELECT count(*) as num FROM `" + date + "` where bigname like ? and iscatch=" + iscatch + ";"
			}
			DB.Raw(dudu1, bigname).QueryRow(&num)
		} else {
			if iscatchi == 2 {
				dudu1 = "SELECT count(*) as num FROM `" + date + "`;"
			} else {
				dudu1 = "SELECT count(*) as num FROM `" + date + "` where iscatch=" + iscatch + ";"
			}
			DB.Raw(dudu1).QueryRow(&num)

		}
	} else {
		dudu := "SELECT * FROM `" + date + "`where id like '" + asin + "|%' limit " + strconv.Itoa(start) + "," + strconv.Itoa(rows) + ";"
		//fmt.Println(dudu)
		DB.Raw(dudu).Values(&maps)
		dudu1 := "SELECT count(*) as num FROM `" + date + "`where id like '" + asin + "|%'"
		DB.Raw(dudu1).QueryRow(&num)
	}
	if len(maps) == 0 {
		this.Data["json"] = &map[string]interface{}{"total": num, "rows": []interface{}{}}
	} else {
		this.Data["json"] = &map[string]interface{}{"total": num, "rows": &maps}
	}
	this.ServeJSON()
}

// no use
func (this *UsaController) Export() {
	DB := orm.NewOrm()
	err := DB.Using("usadatadb")
	if err != nil {
		beego.Error("usadatadb err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	date := this.GetString("datename")
	date = strings.Replace(date, "-", "", -1)
	bigname := this.GetString("bigname")
	dudu := ""
	if bigname == "all" {
		this.Rsp(false, "不能导出全部分类")
	} else {
		dudu = "SELECT * FROM `" + date + "` where bigname like ? and rank<15000 order by rank limit 15000;"
	}
	var maps []orm.Params
	num, err := DB.Raw(dudu, bigname).Values(&maps)
	if num == 0 || err != nil {
		this.Rsp(false, "日期找不到或查询结果为空")
	}
	filename := strings.Replace(strings.Replace(bigname, "&", "", -1), " ", "", -1)
	f, err := os.Create("file/data/" + filename + "-" + date + ".csv")
	if err != nil {
		this.Rsp(false, "Excel创建出错")
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)
	w.Write([]string{"标题", "商品链接", "价格", "小类名", "小类链接", "小类排名", "大类名", "真实大类名", "大类排名", "ReviewNum", "ReviewScore", "图片链接", "状态", "列表抓取时间", "详情抓取时间"})
	//w.Write([]string{"标题", "Asin", "商品链接", "大类名", "大类链接", "大类排名", "抓取时间"})

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
		if temp["iscatch"] == "1" {
			temp["iscatch"] = "已抓"
		} else {
			temp["iscatch"] = "待抓"
		}
		w.Write([]string{temp["title"], temp["url"], temp["price"], temp["name"], temp["purl"], temp["smallrank"], temp["bigname"], temp["rbigname"], temp["rank"], temp["reviews"], temp["score"], temp["img"], temp["iscatch"], temp["createtime"], temp["updatetime"]})
		//w.Write([]string{temp["title"], temp["id"], "https://www.amazon.com/dp/" + temp["id"], temp["bigname"], temp["purl"], temp["rank"], temp["createtime"]})

		// iscatch:1
		// asin:B000BVXDZM
		// dbname:19
	}
	w.Flush()

	this.Rsp(true, "/file/data/"+filename+"-"+date+".csv")
}
