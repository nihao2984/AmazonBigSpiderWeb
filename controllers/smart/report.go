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
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/hunterhug/GoSpider/util"
	"github.com/hunterhug/AmazonBigSpiderWeb/lib"
	"os"
	"strconv"
	"strings"
)

type ReportController struct {
	baseController
}

func (this *ReportController) Index() {
	DB := orm.NewOrm()
	err := DB.Using("dbback")
	if err != nil {
		beego.Error("dbback err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	var categorys []orm.Params
	DB.Raw("SELECT aws as AWS FROM `report` group by aws").Values(&categorys)
	this.Data["category"] = &categorys
	this.Layout = this.GetTemplate() + "/base/layout.html"
	this.TplName = this.GetTemplate() + "/report/list.html"

}

func (this *ReportController) Query() {
	DB := orm.NewOrm()
	err := DB.Using("dbback")
	if err != nil {
		beego.Error("dbback err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	page, _ := this.GetInt("page", 1)
	rows, _ := this.GetInt("rows", 30)
	bdate := this.GetString("bdatename")
	edate := this.GetString("edatename")
	aws := this.GetString("aws")
	asin := strings.TrimSpace(this.GetString("asin"))
	bdates := strings.Replace(bdate, "-", "", -1)
	edates := strings.Replace(edate, "-", "", -1)
	start := (page - 1) * rows
	num := 0
	var maps []orm.Params

	where := []string{}
	if bdates != "" {
		_, e1 := util.SI(bdates)
		if e1 != nil {
			this.Data["json"] = &map[string]interface{}{"total": num, "rows": []interface{}{}}
			this.ServeJSON()
			return
		} else {
			where = append(where, "d>="+bdates)
		}
	}
	if edates != "" {
		_, e1 := util.SI(edates)
		if e1 != nil {
			this.Data["json"] = &map[string]interface{}{"total": num, "rows": []interface{}{}}
			this.ServeJSON()
			return
		} else {
			where = append(where, "d<="+edates)
		}
	}

	if aws != "" {
		where = append(where, `aws="`+aws+`"`)
	}
	if asin != "" {
		where = append(where, `asin="`+asin+`"`)
	}
	wheresql := ""
	where = append(where, "status=1")
	if len(where) == 0 {

	} else {
		wheresql = strings.Join(where, " and ")
		wheresql = "where " + wheresql
	}
	dudu := "SELECT * FROM `report` " + wheresql + " order by d,uv desc limit " + strconv.Itoa(start) + "," + strconv.Itoa(rows) + ";"

	DB.Raw(dudu).Values(&maps)

	dudu1 := "SELECT count(*) as num FROM `report` " + wheresql + ";"

	DB.Raw(dudu1).QueryRow(&num)

	if len(maps) == 0 {
		this.Data["json"] = &map[string]interface{}{"total": num, "rows": []interface{}{}}
	} else {
		this.Data["json"] = &map[string]interface{}{"total": num, "rows": &maps}
	}
	this.ServeJSON()
}

func (this *ReportController) Export() {
	DB := orm.NewOrm()
	err := DB.Using("dbback")
	if err != nil {
		beego.Error("dbback err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	bdate := this.GetString("bdatename")
	edate := this.GetString("edatename")
	aws := this.GetString("aws")
	bdates := strings.Replace(bdate, "-", "", -1)
	edates := strings.Replace(edate, "-", "", -1)
	asin := strings.TrimSpace(this.GetString("asin"))
	var maps []orm.Params

	where := []string{}
	if bdates != "" {
		_, e1 := util.SI(bdates)
		if e1 != nil {
			this.Rsp(false, "起始时间错误")
		} else {
			where = append(where, "d>="+bdates)
		}
	}
	if edates != "" {
		_, e1 := util.SI(edates)
		if e1 != nil {
			this.Rsp(false, "截止时间错误")
		} else {
			where = append(where, "d<="+edates)
		}
	}

	if aws != "" {
		where = append(where, `aws="`+aws+`"`)
	}
	if asin != "" {
		where = append(where, `asin="`+asin+`"`)
	}
	where = append(where, "status=1")
	wheresql := ""
	if len(where) == 0 {

	} else {
		wheresql = strings.Join(where, " and ")
		wheresql = "where " + wheresql
	}
	dudu := "SELECT * FROM `report` " + wheresql + " order by d,uv desc limit 10000;"

	num, err := DB.Raw(dudu).Values(&maps)

	if err != nil {
		this.Rsp(false, "内部错误："+err.Error())
	}
	if num == 0 {
		this.Rsp(false, "日期找不到或查询结果为空")
	}
	f, err := os.Create("file/back/" + bdates + "-" + edates + "-" + aws + ".csv")
	if err != nil {
		this.Rsp(false, "Excel创建出错")
	}
	defer f.Close()

	f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

	w := csv.NewWriter(f)
	w.Write([]string{`Asin(父)`, `Asin(子)`, `商品名称`, `买家访问次数`, `该日买家访问次数百分比`, `页面浏览次数`, `该日页面浏览次数百分比`, `该日购物车占比`, `该日已订购商品数量`, `该日订单商品数量转化率`, `该日已订购商品销售额`, `该日订单数`, `日期`, `店铺名`})
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
		w.Write([]string{temp["pasin"], temp["asin"], temp["title"], temp["uv"], temp["uvb"], temp["pv"], temp["pvb"], temp["bpvb"], temp["on"], temp["onr"], temp["v"], temp["c"], temp["d"], temp["aws"]})

		// iscatch:1
		// asin:B000BVXDZM
		// dbname:19
	}
	w.Flush()

	this.Rsp(true, "/file/back/"+bdates+"-"+edates+"-"+aws+".csv")
}

type Sizer interface {
	Size() int64
}

var Filebytes = 1 << 25 // (1<<25)/1000.0/1000.0 33.54 不能超出33M
func (this *ReportController) Import() {
	DB := orm.NewOrm()
	err := DB.Using("dbback")
	if err != nil {
		beego.Error("dbback err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	//初始化
	fileerror := 1 //上传不成功标志位
	fileSize := 0
	message := "什么都没发生"
	datastring := ""
	awsname := ""
	f, h, err := this.Ctx.Request.FormFile("imgFile")

	//关闭数据流
	defer f.Close()

	//出现错误
	if err != nil {
		message = err.Error()

	} else {
		message = h.Filename

		filename := strings.Split(message, ".")[0]
		filetemp := strings.Split(filename, "-")
		if len(filetemp) != 2 {
			message = "文件命名格式为  (20161012-店铺名)"
			goto END
		}

		datastring = filetemp[0]
		_, e := util.SI(datastring)
		if e != nil || len(filetemp[0]) != 8 {
			message = "文件命名日期不对 (20161012-店铺名)"
			goto END
		}

		awsname = filetemp[1]
		if len(awsname) > 20 {
			message = "文件店铺名过长 (20161012-店铺名)"
			goto END
		}
		//判断文件是否允许被添加

		filesuffix := lib.GetFileSuffix(message)
		//是否后缀正确
		if filesuffix == "csv" {
			//获取大小
			if fileSizer, ok := f.(Sizer); ok {
				fileSize = int(fileSizer.Size())
				// fmt.Printf("上传%v文件的大小为: %v", fileSize, h.Filename)
				if fileSize > Filebytes {
					message = "获取上传文件错误:文件大小超出" + util.IS(Filebytes) + "bytes"
					goto END
				}
			}

		} else {
			message = "文件后缀不被允许"
			goto END
		}

		bytess := make([]byte, int(fileSize))
		s, e := h.Open()
		if e != nil {
			message = e.Error()
			goto END
		} else {
			_, ee := s.Read(bytess)
			if ee != nil {
				message = ee.Error()
				goto END
			}
		}
		csv := string(bytess)
		good := strings.Split(csv, "\n")
		result := [][]string{}
		for _, v := range good {
			row := strings.Split(v, `","`)
			for i := 0; i <= len(row)-1; i++ {
				row[i] = strings.Replace(row[i], ",", "", -1)
				row[i] = strings.Replace(row[i], `"`, "", -1)
			}
			//println(len(row))
			if len(row) == 12 {
				result = append(result, row)
			}
		}
		//fmt.Printf("%#v",result)
		for _, v := range result {
			visitnum, e := util.SI(v[3])
			if e != nil {
				continue
			}
			pv, e := util.SI(v[5])
			if e != nil {
				continue
			}

			on, e := util.SI(v[8])
			if e != nil {
				continue
			}

			c, e := util.SI(v[11])
			if e != nil {
				continue
			}
			va := strings.Replace(v[10], "US$", "", -1)
			va = strings.Replace(va, "$", "", -1)
			va = strings.Replace(va, ",", "", -1)
			vv, e := strconv.ParseFloat(va, 32)
			if e != nil {
				continue
			}
			id := lib.Md5(datastring + "-" + awsname + "-" + v[1])

			//fmt.Println(visitnum, id)
			sql := "Replace INTO `report`(`id`,`pasin`,`asin`,`title`,`uv`,`uvb`,`pv`,`pvb`,`bpvb`,`on`,`onr`,`v`,`c`,`d`,`aws`,`status`)VALUES"
			sql = sql + "(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,1)"
			_, sqlr := DB.Raw(sql, id, v[0], v[1], v[2], visitnum, v[4], pv, v[6], v[7], on, v[9], vv, c, datastring, awsname).Exec()
			if sqlr != nil {
				beego.Error(sqlr.Error())
				continue
			}
		}
		f.Close()
		message = "成功执行"
		fileerror = 0
	}
END:

	this.Data["json"] = &map[string]interface{}{"error": fileerror, "message": message}

	this.ServeJSON()
}

func (this *ReportController) Delete() {
	id := this.GetString("id")
	if id == "" {
		this.Rsp(false, "主键为空")
	}
	sql := "update `report` set status=0 where id=?"
	DB := orm.NewOrm()
	err := DB.Using("dbback")
	if err != nil {
		beego.Error("dbback err:" + err.Error())
		this.Rsp(false, err.Error())
	}
	_, e := DB.Raw(sql, id).Exec()
	if e != nil {
		this.Rsp(false, e.Error())
	}
	this.Rsp(true, "成功！")
}
