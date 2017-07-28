package blog

import (
	_ "fmt"
	_ "github.com/astaxie/beego"
	. "github.com/hunterhug/AmazonBigSpiderWeb/lib"
	"io/ioutil"
	"strings"
)

var FileAllow = map[string][]string{
	"image": {
		"jpg", "jpeg", "png", "bmp", "gif"},
	"flash": {
		"swf", "flv"},
	"media": {
		"swf", "flv", "mp3", "wav", "wma", "wmv", "mid", "avi", "mpg", "asf", "rm", "rmvb"},
	"file": {
		"doc", "docx", "xls", "xlsx", "ppt", "htm", "html", "txt", "zip", "rar", "gz", "bz2"},
	"other": {
		"jpg", "jpeg", "png", "bmp", "gif","swf", "flv","mp3",
		"wav", "wma", "wmv", "mid", "avi", "mpg", "asf", "rm", "rmvb",
		"doc", "docx", "xls", "xlsx", "ppt", "htm", "html", "txt", "zip", "rar", "gz", "bz2"}}

var Filebytes = 1 << 25 // (1<<25)/1000.0/1000.0 33.54 不能超出33M

type UploadController struct {
	baseController
}

type Sizer interface {
	Size() int64
}

	/*

	imgFile: 文件form名称
	dir: 上传类型，分别为image、flash、media、file、other
	返回格式(JSON)

	//成功时
	{
		"error" : 0,
		"url" : "http://www.example.com/path/to/file.ext"
	}
	//失败时
	{
		"error" : 1,
		"message" : "错误信息"
		"token":"加密文件地址"
	}
	*/
func (this *UploadController) UploadFile() {
	//初始化
	fileerror := 1       //上传不成功标志位
	dirpath := ""        //保存路径
	filename := ""       //文件名
	filetype := this.GetString("dir", "other")

	message := "什么都没发生"

	//得到表单数据
	// f, h, err := this.GetFile("imgFile")
	f, h, err := this.Ctx.Request.FormFile("imgFile")

	//关闭数据流
	defer f.Close()

	//出现错误
	if err != nil {
		message = err.Error()
	} else {

		//判断文件是否允许被添加
		//dir类型正确
		fileallowarray, ok := FileAllow[filetype]
		if ok {
			//得到文件后缀
			filesuffix := GetFileSuffix(h.Filename)
			//是否后缀正确
			if InArray(fileallowarray, filesuffix) {
				//获取大小
				if fileSizer, ok := f.(Sizer); ok {
					fileSize := fileSizer.Size()
					// fmt.Printf("上传%v文件的大小为: %v", fileSize, h.Filename)
					if fileSize > int64(Filebytes) {
						message = "获取上传文件错误:文件大小超出33M"
						goto END
					}
				} else {
					message = "获取上传文件错误:无法读取文件大小"
				}
				//创建文件夹
				dirpath, err = MakeFileDir(filetype + "/" + GetTodayString())
				if err != nil {
					message = err.Error()
				} else {
					//新建文件名
					filename = h.Filename
					if HasFile(dirpath + "/" + filename) {
						//文件名存在，放大招，改名
						filename = GetTimeString() + h.Filename
						// message = "文件重名"
						// goto END
					}
					//复制文件
					err = CopyFS(f, dirpath + "/" + filename)
					if err != nil {
						message = err.Error()
					} else {
						fileerror = 0
					}
				}
			} else {
				message = "文件后缀不被允许"
			}
		} else {
			message = "dir参数不允许"
		}
	}
	END:
	if fileerror == 1 {
		this.Data["json"] = &map[string]interface{}{"error": fileerror, "message": message}
	} else {
		name := dirpath + "/" + filename
		//http://lulijuan505.blog.163.com/blog/static/308369112015322102455860/
		//Base64产生的/ + =出现在url会有问题
		/*
			base64
			1、包含A-Z a-z 0-9 和加号“+”，斜杠“/” 用来作为开始的64个数字. 等号“=”用来作为后缀用途。
			2、2进制的.
			3、要比源数据多33%。
			4、常用于邮件。

			urlencode
			除了 -_. 之外的所有非字母数字字符都将被替换成百分号（%）后跟两位十六进制数，空格则编码为加号（+）
			  在神马情况下用

		*/
		token:=Base64E(UrlE(name))
		//urlstring := "/public/file/getfile?token=" + token
		//fmt.Println(name)
		this.Data["json"] = &map[string]interface{}{"error": fileerror, "url":"/"+name,"token":token}
	}
	this.ServeJSON()
}

func (this *UploadController) GetWebFile() {
	id := this.GetString("token", "")
	id = UrlD(Base64D(id))
	//fmt.Println(id)
	if id == "" {
		this.StopRun()
	}
	if strings.HasPrefix(id, GetFileBaseDir()) {
		data, e := ioutil.ReadFile(id)
		if e != nil {
			this.StopRun()
		}
		this.Ctx.ResponseWriter.WriteHeader(200)
		this.Ctx.ResponseWriter.Write(data)
	} else {
		this.StopRun()
	}

}

func (this *UploadController) FileManage() {

}
