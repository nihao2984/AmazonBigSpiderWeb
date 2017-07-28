package lib

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"io"
	"io/ioutil"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

//字符串base64加密
func Base64E(urlstring string) string {
	str := []byte(urlstring)
	data := base64.StdEncoding.EncodeToString(str)
	return data
}

//字符串base64解密
func Base64D(urlxxstring string) string {
	data, err := base64.StdEncoding.DecodeString(urlxxstring)
	if err != nil {
		return ""
	}
	s := fmt.Sprintf("%q", data)
	s = strings.Replace(s, "\"", "", -1)
	return s
}

//url转义
func UrlE(s string) string {
	return url.QueryEscape(s)
}

//url解义
func UrlD(s string) string {
	s, e := url.QueryUnescape(s)
	if e != nil {
		return e.Error()
	} else {
		return s
	}
}

//得到系统时间
func GetTime() time.Time {
	timezone := float64(0)
	v := beego.AppConfig.String("timezone")
	timezone, _ = strconv.ParseFloat(v, 64)
	add := timezone * float64(time.Hour)
	return time.Now().UTC().Add(time.Duration(add))
}

/*"2006-01-02 15:04:05"*/
//得到今天日期字符串
func GetTodayString() string {
	timezone := float64(0)
	v := beego.AppConfig.String("timezone")
	timezone, _ = strconv.ParseFloat(v, 64)
	add := timezone * float64(time.Hour)
	return time.Now().UTC().Add(time.Duration(add)).Format("20060102")
}

//得到时间字符串
func GetTimeString() string {
	timezone := float64(0)
	v := beego.AppConfig.String("timezone")
	timezone, _ = strconv.ParseFloat(v, 64)
	add := timezone * float64(time.Hour)
	return time.Now().UTC().Add(time.Duration(add)).Format("20060102150405")
}

//上传文件根目录
func GetFileBaseDir() string {
	return beego.AppConfig.String("filebasepath")
}

//创建上传文件夹子文件夹
func MakeFileDir(s string) (filedir string, err error) {
	filedir = GetFileBaseDir() + "/" + s
	err = os.MkdirAll(filedir, 0777)
	return filedir, err
}

//判断文件或文件夹是否存在
func HasFile(s string) bool {
	f, err := os.Open(s)
	if err != nil && os.IsNotExist(err) {
		return false
	}
	f.Close()
	return true
}

//File-File复制文件
func CopyFF(src io.Reader, dst io.Writer) error {
	_, err := io.Copy(dst, src)
	return err
}

//File-String复制文件
func CopyFS(src io.Reader, dst string) error {
	f, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = io.Copy(f, src)
	return err
}

//判断是否是文件
func IsFile(filepath string) bool {
	fielinfo, err := os.Stat(filepath)
	if err != nil {
		return false
	} else {
		if fielinfo.IsDir() {
			return false
		} else {
			return true
		}
	}
}

//判断是否是文件夹
func IsDir(filepath string) bool {
	fielinfo, err := os.Stat(filepath)
	if err != nil {
		return false
	} else {
		if fielinfo.IsDir() {
			return true
		} else {
			return false
		}
	}
}

//文件状态
func FileStatus(filepath string) {
	fielinfo, err := os.Stat(filepath)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("%v", fielinfo)
	}
}

//文件夹下数量
func SizeofDir(dirPth string) int {
	if IsDir(dirPth) {
		files, _ := ioutil.ReadDir(dirPth)
		return len(files)
	}

	return 0
}

//字符串是否在字符串数组中
func InArray(sa []string, a string) bool {
	for _, v := range sa {
		if a == v {
			return true
		}
	}
	return false
}

//获取文件后缀
func GetFileSuffix(f string) string {
	fa := strings.Split(f, ".")
	if len(fa) == 0 {
		return ""
	} else {
		return fa[len(fa)-1]
	}
}

//create md5 string
func Strtomd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

//password hash function
func Pwdhash(str string) string {
	return Strtomd5(str)
}

func Md5(str string) string {
	return Strtomd5(str)
}

func StringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
}

func Rawurlencode(str string) string {
	return strings.Replace(url.QueryEscape(str), "+", "%20", -1)
}

//获取用户IP地址
func GetClientIp(this *context.Context) string {
	s := strings.Split(this.Request.RemoteAddr, ":")
	//fmt.Println(s[0])
	if s[0] == "127.0.0.1" {
		//fmt.Printf("%#v\n",this.Request.Header)
		if v, ok := this.Request.Header["X-Real-Ip"]; ok {
			if len(v) > 0 {
				return v[0]
			}
		}
	}
	return s[0]
}
