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
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hunterhug/GoSpider/util"
	"os"
	"strconv"
	"strings"
)

type BigController struct {
	baseController
}

func (this *BigController) Index() {
	DB := orm.NewOrm()
	err := DB.Using("usabasicdb")
	if err != nil {
		beego.Error("usabasicdb err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	var categorys []orm.Params
	DB.Raw("SELECT bigpname as Bigpname,id FROM smart_category where pid=0 group by bigpname,id").Values(&categorys)
	this.Data["category"] = &categorys
	//fmt.Printf("----%#v----", categorys)
	this.Layout = this.GetTemplate() + "/base/layout.html"
	this.TplName = this.GetTemplate() + "/big/list.html"

}
func hashcode(asin string) string {
	dd := []byte(util.Md5("iloveyou"+asin+"hunterhug") + util.Md5(asin))
	sum := 0
	for _, i := range dd {
		sum = sum + int(i)
	}
	hashcode := sum % 81
	return util.IS(hashcode)
}

func (this *BigController) Asin() {
	asin := this.GetString("asin", "")
	if asin == "" {
		this.Rsp(false, "找不到该Asin历史趋势")
	}
	tablename := hashcode(asin)
	sql := fmt.Sprintf(`SELECT * FROM %sA%s%s where id="%s" order by day;`, "`", tablename, "`", asin)
	//beego.Error(tablename)
	DB := orm.NewOrm()
	err := DB.Using("usahashdb")
	if err != nil {
		beego.Error("usahashdb err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	var maps []orm.Params
	DB.Raw(sql).Values(&maps)
	if len(maps) == 0 {
		this.Rsp(false, "结果为空")
	}
	data := map[string][]interface{}{}
	data["day"] = []interface{}{}
	data["price"] = []interface{}{}
	data["score"] = []interface{}{}
	data["rank"] = []interface{}{}
	data["reviews"] = []interface{}{}
	data["bigname"] = []interface{}{}
	for _, v := range maps {
		data["bigname"] = append(data["bigname"], v["bigname"])
		data["day"] = append(data["day"], v["day"])
		data["price"] = append(data["price"], v["price"])
		data["score"] = append(data["score"], v["score"])
		data["rank"] = append(data["rank"], v["rank"])
		data["reviews"] = append(data["reviews"], v["reviews"])
	}
	d, _ := json.Marshal(maps[0])
	this.Data["json"] = string(d)

	dd, _ := json.Marshal(data)
	this.Data["ddjson"] = string(dd)

	this.Data["title"] = "Amazon.com:" + asin
	this.TplName = this.GetTemplate() + "/big/asin.html"

}

func (this *BigController) Query() {
	DB := orm.NewOrm()
	err := DB.Using("usadatadb")
	if err != nil {
		beego.Error("usadatadb err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	asin := this.GetString("asin")
	page, _ := this.GetInt("page", 1)
	rows, _ := this.GetInt("rows", 30)
	date := this.GetString("datename")
	date = strings.Replace(date, "-", "", -1)
	bigname := this.GetString("bigname")
	start := (page - 1) * rows
	num := 0
	var maps []orm.Params
	if date == "" {
		date = "20161101"
	}
	if asin == "" {
		bprice, _ := this.GetFloat("bprice", -1)
		eprice, _ := this.GetFloat("eprice", -1)
		breview, _ := this.GetFloat("breview", -1)
		ereview, _ := this.GetFloat("ereview", -1)
		bscore, _ := this.GetFloat("bscore", -1)
		escore, _ := this.GetFloat("escore", -1)
		brate, _ := this.GetFloat("brate", -1)
		erate, _ := this.GetFloat("erate", -1)
		fba := this.GetString("fba")
		sold := this.GetString("sold")
		where := []string{}
		wheresql := ""
		if fba != "" {
			where = append(where, `ship ="`+fba+`"`)
		}
		if sold != "" {
			where = append(where, `sold ="`+sold+`"`)
		}
		if bigname != "" {
			where = append(where, `bigname like "`+bigname+`"`)
		}
		if bprice >= 0 {
			where = append(where, "price >="+this.GetString("bprice"))
		}
		if eprice >= 0 {
			where = append(where, "price <="+this.GetString("eprice"))
		}
		if breview >= 0 {
			where = append(where, "reviews >="+this.GetString("breview"))
		}
		if ereview >= 0 {
			where = append(where, "reviews <="+this.GetString("ereview"))
		}
		if bscore >= 0 {
			where = append(where, "score >="+this.GetString("bscore"))
		}
		if escore >= 0 {
			where = append(where, "score <="+this.GetString("escore"))
		}
		if brate >= 0 {
			where = append(where, "rank >="+this.GetString("brate"))
		}
		if erate >= 0 {
			where = append(where, "rank <="+this.GetString("erate"))
		}
		if len(where) != 0 {
			wheresql = " where " + strings.Join(where, " and ")
		}
		dudu := "SELECT * FROM `Asin" + date + "`" + wheresql + " order by bigname,rank limit " + strconv.Itoa(start) + "," + strconv.Itoa(rows) + ";"

		DB.Raw(dudu).Values(&maps)

		dudu1 := "SELECT count(*) as num FROM `Asin" + date + "`" + wheresql + ";"

		DB.Raw(dudu1).QueryRow(&num)

	} else {
		dudu := "SELECT * FROM `Asin" + date + "` where id=?"
		//fmt.Println(dudu)
		DB.Raw(dudu, asin).Values(&maps)
		dudu1 := "SELECT count(*) as num FROM `Asin" + date + "` where id=?"
		DB.Raw(dudu1, asin).QueryRow(&num)

	}
	if len(maps) == 0 {
		this.Data["json"] = &map[string]interface{}{"total": num, "rows": []interface{}{}}
	} else {
		this.Data["json"] = &map[string]interface{}{"total": num, "rows": &maps}
	}
	this.ServeJSON()
}

func (this *BigController) Export() {
	DB := orm.NewOrm()
	err := DB.Using("usadatadb")
	if err != nil {
		beego.Error("usadatadb err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	date := this.GetString("datename")
	date = strings.Replace(date, "-", "", -1)
	bigname := this.GetString("bigname")
	var maps []orm.Params
	if date == "" {
		this.Rsp(false, "日期不能为空")
	}
	isall := this.GetString("isall", "0")
	if isall == "1" {
		if bigname == "" {
			this.Rsp(false, "不能导出全部大类")
		}
		dudu := "SELECT * FROM `Asin" + date + "` where bigname like '" + bigname + "' and rank<=15000 order by rank limit 20000;"

		num, err := DB.Raw(dudu).Values(&maps)

		if num == 0 || err != nil {
			this.Rsp(false, "查询结果为空")
		}
	} else {
		page, _ := this.GetInt("page", 1)
		rows, _ := this.GetInt("rows", 30)
		start := (page - 1) * rows
		bprice, _ := this.GetFloat("bprice", -1)
		eprice, _ := this.GetFloat("eprice", -1)
		breview, _ := this.GetFloat("breview", -1)
		ereview, _ := this.GetFloat("ereview", -1)
		bscore, _ := this.GetFloat("bscore", -1)
		escore, _ := this.GetFloat("escore", -1)
		brate, _ := this.GetFloat("brate", -1)
		erate, _ := this.GetFloat("erate", -1)
		fba := this.GetString("fba")
		sold := this.GetString("sold")
		where := []string{}
		wheresql := ""
		if fba != "" {
			where = append(where, `ship ="`+fba+`"`)
		}
		if sold != "" {
			where = append(where, `sold ="`+sold+`"`)
		}
		if bigname != "" {
			where = append(where, `bigname like "`+bigname+`"`)
		}
		if bprice >= 0 {
			where = append(where, "price >="+this.GetString("bprice"))
		}
		if eprice >= 0 {
			where = append(where, "price <="+this.GetString("eprice"))
		}
		if breview >= 0 {
			where = append(where, "reviews >="+this.GetString("breview"))
		}
		if ereview >= 0 {
			where = append(where, "reviews <="+this.GetString("ereview"))
		}
		if bscore >= 0 {
			where = append(where, "score >="+this.GetString("bscore"))
		}
		if escore >= 0 {
			where = append(where, "score <="+this.GetString("escore"))
		}
		if brate >= 0 {
			where = append(where, "rank >="+this.GetString("brate"))
		}
		if erate >= 0 {
			where = append(where, "rank <="+this.GetString("erate"))
		}
		if len(where) != 0 {
			wheresql = " where " + strings.Join(where, " and ")
		}
		dudu := "SELECT * FROM `Asin" + date + "`" + wheresql + " order by bigname,rank limit " + strconv.Itoa(start) + "," + strconv.Itoa(rows) + ";"

		num, err := DB.Raw(dudu).Values(&maps)

		if num == 0 || err != nil {
			this.Rsp(false, "查询结果为空")
		}
	}
	filename := strings.Replace(strings.Replace(bigname, "&", "", -1), " ", "", -1)
	f, err := os.Create("file/data/big-" + filename + "-" + date + ".csv")
	if err != nil {
		this.Rsp(false, "Excel创建出错")
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)
	//w.Write([]string{"标题", "商品链接", "价格", "小类名", "小类链接", "小类排名", "大类名", "真实大类名", "大类排名", "ReviewNum", "ReviewScore", "图片链接", "状态", "列表抓取时间", "详情抓取时间"})
	w.Write([]string{"标题", "Asin", "商品链接", "大类名", "大类排名", "price", "reviews", "score", "ship", "sold", "img", "抓取时间"})

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
		if _, ok := temp["price"]; ok == false {
			temp["price"] = ""
		}
		if _, ok := temp["reviews"]; ok == false {
			temp["reviews"] = ""
		}
		if _, ok := temp["score"]; ok == false {
			temp["score"] = ""
		}
		if _, ok := temp["ship"]; ok == false {
			temp["ship"] = ""
		}
		if _, ok := temp["sold"]; ok == false {
			temp["sold"] = ""
		}
		if _, ok := temp["img"]; ok == false {
			temp["img"] = ""
		}

		w.Write([]string{temp["title"], temp["id"], "https://www.amazon.com/dp/" + temp["id"], temp["bigname"], temp["rank"], temp["price"], temp["reviews"], temp["score"], temp["ship"], temp["sold"], temp["img"], temp["createtime"]})

		// iscatch:1
		// asin:B000BVXDZM
		// dbname:19
	}
	w.Flush()

	this.Rsp(true, "/file/data/big-"+filename+"-"+date+".csv")
}
