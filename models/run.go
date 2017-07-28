// 模型包
package models

import (
	"github.com/astaxie/beego"
	"os"
)

// 数据库开跑
func Run() {
	beego.Trace("数据库开跑")
	initArgs()
	Connect()
}

// 数据库初始化
func initArgs() {
	args := os.Args
	for _, v := range args {
		if v == "-s" {
			Syncdb()
			os.Exit(0)
		}
	}
}
