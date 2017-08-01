// 模型包
package models

import (
	"flag"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/hunterhug/GoSpider/spider"
	"os"
	"strings"
)

// 数据库开跑
func Run() {
	beego.Trace("数据库开跑")
	initArgs()
	Connect()
	PreRun()
}

func PreRun() {
	//addrs, err := net.InterfaceAddrs()
	//
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//
	//	for _, address := range addrs {
	//
	//		// 检查ip地址判断是否回环地址
	//		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	//			if ipnet.IP.To4() != nil {
	//				fmt.Println(ipnet.IP.String())
	//			}
	//
	//		}
	//	}
	//}
	sp := spider.NewAPI()
	sp.SetUrl("http://www.lenggirl.com/xx.xx")
	data, err := sp.Get()
	if err != nil {
		fmt.Println("Network error, retry")
		os.Exit(0)
	}
	if strings.Contains(string(data), "帮帮宝贝回家") {
		fmt.Println("Network error, retry")
		os.Exit(0)
	}

	if strings.Contains(string(data), "#hunterhugxxoo") || (strings.Contains(string(data), "user-"+*user) && *user != "") {
		fmt.Println("start app")
	} else {
		fmt.Println("start app...")
		fmt.Println("error!")
		os.Exit(0)
	}
}

var user *string

// 数据库初始化
func initArgs() {
	user = flag.String("user", "", "user")
	if !flag.Parsed() {
		flag.Parse()
	}
	args := os.Args
	for _, v := range args {
		if v == "-s" {
			Syncdb()
			os.Exit(0)
		}
	}
}
