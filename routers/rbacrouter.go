package routers

import (
	"github.com/astaxie/beego"
	"github.com/hunterhug/AmazonBigSpiderWeb/controllers/admin/rbac"
)

// 后台RBAC路由
func rbacrouter() {
	beego.Router("/", &rbac.MainController{}, "*:Index")
	beego.Router("/public", &rbac.MainController{}, "*:Index")
	beego.Router("/public/Index", &rbac.MainController{}, "*:Index")
	beego.Router("/public/Login", &rbac.MainController{}, "*:Login")
	beego.Router("/public/Logout", &rbac.MainController{}, "*:Logout")
	beego.Router("/public/Changepwd", &rbac.MainController{}, "*:Changepwd")

	beego.Router("/rbac/user/AddUser", &rbac.UserController{}, "*:AddUser")
	beego.Router("/rbac/user/UpdateUser", &rbac.UserController{}, "*:UpdateUser")
	beego.Router("/rbac/user/UpdateUserPasswd", &rbac.UserController{}, "*:UpdateUserPasswd")
	beego.Router("/rbac/user/DelUser", &rbac.UserController{}, "*:DelUser")
	beego.Router("/rbac/user/Index", &rbac.UserController{}, "*:Index")

	beego.Router("/rbac/node/AddAndEdit", &rbac.NodeController{}, "*:AddAndEdit")
	beego.Router("/rbac/node/DelNode", &rbac.NodeController{}, "*:DelNode")
	beego.Router("/rbac/node/Index", &rbac.NodeController{}, "*:Index")

	beego.Router("/rbac/group/AddGroup", &rbac.GroupController{}, "*:AddGroup")
	beego.Router("/rbac/group/UpdateGroup", &rbac.GroupController{}, "*:UpdateGroup")
	beego.Router("/rbac/group/DelGroup", &rbac.GroupController{}, "*:DelGroup")
	beego.Router("/rbac/group/Index", &rbac.GroupController{}, "*:Index")

	beego.Router("/rbac/role/AddAndEdit", &rbac.RoleController{}, "*:AddAndEdit")
	beego.Router("/rbac/role/DelRole", &rbac.RoleController{}, "*:DelRole")
	beego.Router("/rbac/role/AccessToNode", &rbac.RoleController{}, "*:AccessToNode")
	beego.Router("/rbac/role/AddAccess", &rbac.RoleController{}, "*:AddAccess")
	beego.Router("/rbac/role/RoleToUserList", &rbac.RoleController{}, "*:RoleToUserList")
	beego.Router("/rbac/role/AddRoleToUser", &rbac.RoleController{}, "*:AddRoleToUser")
	beego.Router("/rbac/role/GetList", &rbac.RoleController{}, "*:Getlist")
	beego.Router("/rbac/role/Index", &rbac.RoleController{}, "*:Index")

}
