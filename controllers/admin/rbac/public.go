package rbac

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	. "github.com/hunterhug/AmazonBigSpiderWeb/controllers"
	. "github.com/hunterhug/AmazonBigSpiderWeb/lib"
	"github.com/hunterhug/AmazonBigSpiderWeb/models/admin"
	// "github.com/hunterhug/AmazonBigSpiderWeb/models/home"
	// "os"
	// "runtime"
	"strconv"
	"strings"
)

type MainController struct {
	CommonController
}

var Cookie7 = beego.AppConfig.String("cookie7")

// 后台首页
func (this *MainController) Index() {
	userinfo := this.GetSession("userinfo")
	//如果没有seesion
	if userinfo == nil && Cookie7 == "1" {
		success, userinfo := CheckCookie(this.Ctx)
		//查看是否有cookie
		if success {
			//有
			//更新登陆时间
			userinfo = admin.UpdateLoginTime(&userinfo)
			userinfo.Logincount += 1
			userinfo.Lastip = GetClientIp(this.Ctx)
			userinfo.Update()
			this.SetSession("userinfo", userinfo)
			//设置权限列表session
			accesslist, _ := GetAccessList(userinfo.Id)
			this.SetSession("accesslist", accesslist)

		} else {
			//没有
			this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
			return
		}
	}

	if userinfo == nil {
		this.Ctx.Redirect(302, beego.AppConfig.String("rbac_auth_gateway"))
		return
	}

	// 权限最重要的部分，入口在此
	// 获取模块rbac-节点 public/index    /rbac/public/index
	// 分组关闭，则该分组所有菜单没有
	// 第一/二层节点关闭,则菜单没有
	tree := this.GetTree()

	//userinfo = this.GetSession("userinfo")
	groups := admin.GroupList()
	this.Data["tree"] = tree
	this.Data["user"] = userinfo.(admin.User)
	this.Data["groups"] = groups
	// this.Data["hostname"], _ = os.Hostname()
	// this.Data["gover"] = runtime.Version()
	// this.Data["os"] = runtime.GOOS
	// this.Data["cpunum"] = runtime.NumCPU()
	// this.Data["arch"] = runtime.GOARCH
	// this.Data["postnum"], _ = new(home.Post).Query().Count()
	// this.Data["tagnum"], _ = new(home.Tag).Query().Count()
	// this.Data["usernum"], _ = new(admin.User).Query().Count()
	this.TplName = this.GetTemplate() + "/public/admin.html"

}

//登录
func (this *MainController) Login() {
	// 查看是否已经登陆过
	userinfo := this.GetSession("userinfo")
	if userinfo == nil && Cookie7 == "1" {
		success, userinfo := CheckCookie(this.Ctx)
		//查看是否有cookie
		if success {
			//更新登陆时间
			userinfo = admin.UpdateLoginTime(&userinfo)
			userinfo.Logincount += 1
			userinfo.Lastip = GetClientIp(this.Ctx)
			userinfo.Update()
			this.SetSession("userinfo", userinfo)
			//设置权限列表session
			accesslist, _ := GetAccessList(userinfo.Id)
			this.SetSession("accesslist", accesslist)
			this.Ctx.Redirect(302, "/public/index")
			return
		}
	} else if userinfo != nil {
		this.Ctx.Redirect(302, "/public/index")
		return
	} else {

	}

	//登陆中
	isajax := this.GetString("isajax")
	if isajax == "1" {
		if Verify(this.Ctx) {
			account := strings.TrimSpace(this.GetString("account"))
			password := strings.TrimSpace(this.GetString("password"))
			remember := this.GetString("remember")
			user, err := CheckLogin(account, password)
			if err == nil {

				//更新登陆时间
				user = admin.UpdateLoginTime(&user)
				user.Logincount += 1
				user.Lastip = GetClientIp(this.Ctx)
				user.Update()
				authkey := Md5(GetClientIp(this.Ctx) + "|" + user.Password)
				if Cookie7 == "1" {
					if remember == "yes" {
						this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authkey, 7*86400)
					} else {
						this.Ctx.SetCookie("auth", strconv.FormatInt(user.Id, 10)+"|"+authkey)
					}
				}
				//设置登陆session
				this.SetSession("userinfo", user)
				//设置权限列表session
				accesslist, _ := GetAccessList(user.Id)
				this.SetSession("accesslist", accesslist)

				this.Ctx.Redirect(302, "/public/index")
				return

			} else {
				this.Data["errmsg"] = err.Error()
			}
		} else {
			this.Data["errmsg"] = "验证码错误"
		}
	}

	this.TplName = this.GetTemplate() + "/public/login.html"
}

//退出登陆
func (this *MainController) Logout() {
	this.DelSession("userinfo")
	this.DelSession("accesslist")
	this.Ctx.SetCookie("auth", "")
	this.Ctx.Redirect(302, "/public/login")
	return
}

//修改密码
func (this *MainController) Changepwd() {
	isajax := this.GetString("isajax")
	userinfo := this.GetSession("userinfo")
	if userinfo == nil {
		this.Rsp(false, "没有登陆")
	}
	if isajax == "1" {
		nowpassword := this.GetString("nowpassword")
		user := new(admin.User)
		user.Id = userinfo.(admin.User).Id
		err := user.Read()
		if err != nil {
			this.Rsp(false, err.Error())
		}
		if Md5(nowpassword) != user.Password {
			this.Rsp(false, "原始密码错误")
		}
		password := this.GetString("password")
		repassword := this.GetString("repassword")
		if password == "" || repassword == "" {
			this.Rsp(false, "不能为空")
		} else if password != repassword {
			this.Rsp(false, "两次密码不一致")
		} else {
			if len(password) < 6 || len(password) > 20 {
				this.Rsp(false, "长度应该6-20")
			}

			user.Password = Pwdhash(password)
			err := user.Update("password")
			if err != nil {
				this.Rsp(false, err.Error())
			} else {
				this.Rsp(true, "更改成功")
			}
		}
	} else {
		this.TplName = this.GetTemplate() + "/public/changepwd.html"
	}
}

func CheckLogin(username string, password string) (user admin.User, err error) {
	//根据名字查找用户
	user = admin.GetUserByUsername(username)
	if user.Id == 0 {
		return user, errors.New("用户不存在或者密码错误")
	}
	if user.Password != Pwdhash(password) {
		return user, errors.New("用户不存在或者密码错误")
	}

	adminuser := beego.AppConfig.String("rbac_admin_user")
	if user.Username != adminuser && user.Status == 2 {
		return user, errors.New("用户未激活")
	}

	return user, nil
}

func CheckCookie(ctx *context.Context) (bool, admin.User) {
	var user admin.User
	//查看是否有cookie
	arr := strings.Split(ctx.GetCookie("auth"), "|")
	if len(arr) == 2 {
		idstr, password := arr[0], arr[1]
		userid, _ := strconv.ParseInt(idstr, 10, 0)
		if userid > 0 {
			user.Id = userid
			// cookie没问题,且已经激活
			adminuser := beego.AppConfig.String("rbac_admin_user")
			if user.Read() == nil && password == Md5(GetClientIp(ctx)+"|"+user.Password) && (user.Username == adminuser || user.Status == 1) {
				return true, user
			} else {
				return false, user
			}
		}
	}
	return false, user
}
