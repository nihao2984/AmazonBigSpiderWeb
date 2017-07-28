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
// RBAC权限包
package controllers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	. "github.com/hunterhug/AmazonBigSpiderWeb/lib"
	m "github.com/hunterhug/AmazonBigSpiderWeb/models/admin"
)

func init() {
	AccessRegister()
}

//check access and register user's nodes
func AccessRegister() {
	var Check = func(ctx *context.Context) {
		user_auth_type, _ := strconv.Atoi(beego.AppConfig.String("user_auth_type"))
		rbac_auth_gateway := beego.AppConfig.String("rbac_auth_gateway")
		var accesslist map[string]bool
		if user_auth_type > 0 {
			params := strings.Split(strings.ToLower(strings.Split(ctx.Request.RequestURI, "?")[0]), "/")
			if CheckAccess(params) {
				uinfo := ctx.Input.Session("userinfo")
				if uinfo == nil && beego.AppConfig.String("cookie7") == "1" {
					arr := strings.Split(ctx.GetCookie("auth"), "|")
					if len(arr) == 2 {
						idstr, password := arr[0], arr[1]
						userid, _ := strconv.ParseInt(idstr, 10, 0)
						if userid > 0 {
							var user m.User
							user.Id = userid
							if user.Read() == nil && password == Md5(GetClientIp(ctx)+"|"+user.Password) {
								uinfo = user

							}
						}
					}
				}
				if uinfo == nil {
					ctx.Redirect(302, rbac_auth_gateway)
					return
				} else {
					//增加sessioN
					ctx.Output.Session("userinfo", uinfo)
				}
				//admin用户不用认证权限
				adminuser := beego.AppConfig.String("rbac_admin_user")
				if uinfo.(m.User).Username == adminuser {
					return
				}

				if user_auth_type == 1 {
					listbysession := ctx.Input.Session("accesslist")
					if listbysession != nil {
						accesslist = listbysession.(map[string]bool)
					} else {
						accesslist, _ = GetAccessList(uinfo.(m.User).Id)
					}
				} else if user_auth_type == 2 {

					accesslist, _ = GetAccessList(uinfo.(m.User).Id)
				}

				ret := AccessDecision(params, accesslist)
				if !ret {
					ctx.Output.JSON(&map[string]interface{}{"status": false, "info": "权限不足"}, true, false)
				}
			}

		}
	}
	beego.InsertFilter("/*", beego.BeforeRouter, Check)
}

//Determine whether need to verify
func CheckAccess(params []string) bool {
	if len(params) <= 3 {
		return false
	}
	for _, nap := range strings.Split(beego.AppConfig.String("not_auth_package"), ",") {
		if params[1] == nap {
			return false
		}
	}
	return true
}

//To test whether permissions
func AccessDecision(params []string, accesslist map[string]bool) bool {
	if CheckAccess(params) {
		s := fmt.Sprintf("%s/%s/%s", params[1], params[2], params[3])
		if len(accesslist) < 1 {
			return false
		}
		_, ok := accesslist[s]
		if ok != false {
			return true
		}
	} else {
		return true
	}
	return false
}

type AccessNode struct {
	Id        int64
	Name      string
	Childrens []*AccessNode
}

//Access permissions list
func GetAccessList(uid int64) (map[string]bool, error) {
	list, err := m.AccessList(uid)
	if err != nil {
		return nil, err
	}
	alist := make([]*AccessNode, 0)
	for _, l := range list {
		if l["Pid"].(int64) == 0 && l["Level"].(int64) == 1 && l["Status"].(int64) == 1 { //最严最好！！！
			anode := new(AccessNode)
			anode.Id = l["Id"].(int64)
			anode.Name = l["Name"].(string)
			alist = append(alist, anode)
		}
	}
	for _, l := range list {
		if l["Level"].(int64) == 2 && l["Status"].(int64) == 1 {
			for _, an := range alist {
				if an.Id == l["Pid"].(int64) {
					anode := new(AccessNode)
					anode.Id = l["Id"].(int64)
					anode.Name = l["Name"].(string)
					an.Childrens = append(an.Childrens, anode)
				}
			}
		}
	}
	for _, l := range list {
		if l["Level"].(int64) == 3 && l["Status"].(int64) == 1 { //补充，如果第三层节点被禁用，则无法访问
			for _, an := range alist {
				for _, an1 := range an.Childrens {
					if an1.Id == l["Pid"].(int64) {
						anode := new(AccessNode)
						anode.Id = l["Id"].(int64)
						anode.Name = l["Name"].(string)
						an1.Childrens = append(an1.Childrens, anode)
					}
				}

			}
		}
	}
	accesslist := make(map[string]bool)
	for _, v := range alist {
		for _, v1 := range v.Childrens {
			for _, v2 := range v1.Childrens {
				vname := strings.Split(v.Name, "/")
				v1name := strings.Split(v1.Name, "/")
				v2name := strings.Split(v2.Name, "/")
				str := fmt.Sprintf("%s/%s/%s", strings.ToLower(vname[0]), strings.ToLower(v1name[0]), strings.ToLower(v2name[0]))
				accesslist[str] = true
			}
		}
	}
	return accesslist, nil
}
