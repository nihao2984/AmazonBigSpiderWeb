// 后台模型数据填充包
package admin

import (
	"fmt"
	. "github.com/hunterhug/AmazonBigSpiderWeb/lib"
	"github.com/hunterhug/AmazonBigSpiderWeb/models/blog"
)

func InitData() {
	InsertUser()
	InsertGroup()
	InsertRole()
	InsertNodes()
	InsertConfig()
}

//插入网站配置
func InsertConfig() {
	fmt.Println("insert config start")
	c := new(blog.Config)
	c.Title = "大数据智能平台"
	err := c.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("insert config end")
}

// 用户数据
func InsertUser() {
	fmt.Println("insert user ...")
	u := new(User)
	u.Username = "admin"
	u.Nickname = "admin"
	u.Password = Pwdhash("admin")
	u.Email = "569929309@qq.com"
	u.Remark = "最高权限的王"
	u.Status = 2
	u.Createtime = GetTime()
	err := u.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}

	u1 := new(User)
	u1.Username = "test"
	u1.Nickname = "测试用户"
	u1.Password = Pwdhash("test")
	u1.Email = "569929309@qq.com"
	u1.Remark = "测试用户"
	u1.Status = 2
	u1.Createtime = GetTime()
	err1 := u1.Insert()
	if err1 != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("insert user end")
}

// 模组数据
func InsertGroup() {
	fmt.Println("insert group ...")
	g := new(Group)
	g.Name = "后台管理"
	g.Title = "后台管理"
	g.Sort = 1
	g.Id = 1
	g.Status = 1
	e := g.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}
	gb := new(Group)
	gb.Name = "美国亚马逊"
	gb.Title = "美国亚马逊"
	gb.Sort = 2
	gb.Id = 2
	gb.Status = 1
	e = gb.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}

	jgb := new(Group)
	jgb.Name = "日本亚马逊"
	jgb.Title = "日本亚马逊"
	jgb.Sort = 3
	jgb.Id = 5
	jgb.Status = 1
	e = jgb.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}

	de := new(Group)
	de.Name = "德国亚马逊"
	de.Title = "德国亚马逊"
	de.Sort = 4
	de.Id = 6
	de.Status = 1
	e = de.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}

	uk := new(Group)
	uk.Name = "英国亚马逊"
	uk.Title = "英国亚马逊"
	uk.Sort = 5
	uk.Id = 7
	uk.Status = 1
	e = uk.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}

	g1 := new(Group)
	g1.Name = "文章管理"
	g1.Title = "文章管理"
	g1.Sort = 6
	g1.Id = 3
	g1.Status = 0
	e = g1.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}
	g2 := new(Group)
	g2.Name = "图片管理"
	g2.Title = "图片管理"
	g2.Sort = 7
	g2.Id = 4
	g2.Status = 0
	e = g2.Insert()
	if e != nil {
		fmt.Println(e.Error())
	}
	fmt.Println("insert group end")
}

// 角色数据
func InsertRole() {
	fmt.Println("insert role ...")
	r := new(Role)
	r.Name = "管理员"
	r.Remark = "权限最高的一群人"
	r.Status = 1
	r.Title = "管理员角色"
	err := r.Insert()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("insert role end")
}

// 节点数据
func InsertNodes() {
	fmt.Println("insert node ...")
	g := new(Group)
	g.Id = 1
	gb := new(Group)
	gb.Id = 2
	g1 := new(Group)
	g1.Id = 3
	g2 := new(Group)
	g2.Id = 4

	jp := new(Group)
	jp.Id = 5

	de := new(Group)
	de.Id = 6

	uk := new(Group)
	uk.Id = 7

	//nodes := make([20]Node)
	nodes := []Node{
		/*

			RBAC管理中心

		*/
		{Id: 1, Name: "rbac", Title: "权限中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g},
		{Id: 2, Name: "node/index", Title: "节点管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 3, Name: "Index", Title: "节点首页", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},
		{Id: 4, Name: "AddAndEdit", Title: "增编节点", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},
		{Id: 5, Name: "DelNode", Title: "删除节点", Remark: "", Level: 3, Pid: 2, Status: 1, Group: g},
		{Id: 6, Name: "user/index", Title: "用户管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 7, Name: "Index", Title: "用户首页", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 8, Name: "AddUser", Title: "增加用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 9, Name: "UpdateUser", Title: "更新用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 10, Name: "DelUser", Title: "删除用户", Remark: "", Level: 3, Pid: 6, Status: 1, Group: g},
		{Id: 11, Name: "group/index", Title: "分组管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 12, Name: "Index", Title: "分组首页", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 13, Name: "AddGroup", Title: "增加分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 14, Name: "UpdateGroup", Title: "更新分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 15, Name: "DelGroup", Title: "删除分组", Remark: "", Level: 3, Pid: 11, Status: 1, Group: g},
		{Id: 16, Name: "role/index", Title: "角色管理", Remark: "", Level: 2, Pid: 1, Status: 1, Group: g},
		{Id: 17, Name: "index", Title: "角色首页", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 18, Name: "AddAndEdit", Title: "增编角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 19, Name: "DelRole", Title: "删除角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 20, Name: "GetList", Title: "列出角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 21, Name: "AccessToNode", Title: "显示权限", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 22, Name: "AddAccess", Title: "增加权限", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 23, Name: "RoleToUserList", Title: "列出角色下用户", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},
		{Id: 24, Name: "AddRoleToUser", Title: "授予用户角色", Remark: "", Level: 3, Pid: 16, Status: 1, Group: g},

		/*

			配置中心

		*/
		{Id: 25, Name: "config", Title: "配置中心", Remark: "", Level: 1, Pid: 0, Status: 0, Group: g},
		//-------
		//网站配置
		{Id: 26, Name: "option/index", Title: "网站配置", Remark: "", Level: 2, Pid: 25, Status: 1, Group: g},
		{Id: 27, Name: "Index", Title: "网站配置首页", Remark: "", Level: 3, Pid: 26, Status: 1, Group: g},
		{Id: 28, Name: "UpdateOption", Title: "更新网站配置", Remark: "", Level: 3, Pid: 26, Status: 1, Group: g},
		//网站配置
		//个人信息
		{Id: 29, Name: "user/index", Title: "个人信息", Remark: "", Level: 2, Pid: 25, Status: 1, Group: g},
		{Id: 30, Name: "Index", Title: "个人信息首页", Remark: "", Level: 3, Pid: 29, Status: 1, Group: g},
		{Id: 31, Name: "UpdateUser", Title: "更新个人信息", Remark: "", Level: 3, Pid: 29, Status: 1, Group: g},
		//个人信息

		/*

			文章中心

		*/
		{Id: 32, Name: "blog", Title: "文章中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g1},
		//------
		//文章目录
		{Id: 33, Name: "category/index", Title: "目录列表", Remark: "", Level: 2, Pid: 32, Status: 1, Group: g1},
		{Id: 34, Name: "Index", Title: "目录列表首页", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		{Id: 35, Name: "AddCategory", Title: "增加目录", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		{Id: 36, Name: "UpdateCategory", Title: "修改目录", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},
		//文章目录
		//文章
		{Id: 37, Name: "paper/index", Title: "文章列表", Remark: "", Level: 2, Pid: 32, Status: 1, Group: g1},
		{Id: 38, Name: "Index", Title: "文章列表首页", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 39, Name: "AddPaper", Title: "增加文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 40, Name: "UpdatePaper", Title: "修改文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 41, Name: "DeletePaper", Title: "回收文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		{Id: 42, Name: "RealDelPaper", Title: "删除文章", Remark: "", Level: 3, Pid: 37, Status: 1, Group: g1},
		//文章

		/*

			图片管理

		*/
		{Id: 43, Name: "picture", Title: "图片中心", Remark: "", Level: 1, Pid: 0, Status: 1, Group: g2},
		//---------
		//相册
		{Id: 44, Name: "album/index", Title: "相册管理", Remark: "", Level: 2, Pid: 43, Status: 1, Group: g2},
		{Id: 45, Name: "Index", Title: "相册首页", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 46, Name: "AddAlbum", Title: "增加相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 47, Name: "DeleteAlbum", Title: "删除相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		{Id: 48, Name: "UpdateAlbum", Title: "修改相册", Remark: "", Level: 3, Pid: 44, Status: 1, Group: g2},
		//相册
		//图片
		{Id: 49, Name: "photo/index", Title: "图片管理", Remark: "", Level: 2, Pid: 43, Status: 1, Group: g2},
		{Id: 50, Name: "Index", Title: "图片首页", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 51, Name: "AddPhoto", Title: "增加图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 52, Name: "DeletePhoto", Title: "回收图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 53, Name: "UpdatePhoto", Title: "修改图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		{Id: 54, Name: "RealDelPhoto", Title: "删除图片", Remark: "", Level: 3, Pid: 49, Status: 1, Group: g2},
		//图片

		//补充的。ID无效
		{Id: 55, Name: "DeleteCategory", Title: "删除目录", Remark: "", Level: 3, Pid: 33, Status: 1, Group: g1},

		///
		{Id: 56, Name: "paper/rubbish", Title: "文章回收站", Remark: "", Level: 2, Pid: 32, Status: 1, Group: g1},
		{Id: 57, Name: "Rubbish", Title: "文章回收站", Remark: "", Level: 3, Pid: 56, Status: 1, Group: g1},

		///
		{Id: 58, Name: "photo/rubbish", Title: "图片回收站", Remark: "", Level: 2, Pid: 43, Status: 1, Group: g2},
		{Id: 59, Name: "Rubbish", Title: "图片回收站", Remark: "", Level: 3, Pid: 58, Status: 1, Group: g2},

		//首页图片轮转
		{Id: 60, Name: "roll/index", Title: "首页轮转", Remark: "", Level: 2, Pid: 25, Status: 1, Group: g},
		{Id: 61, Name: "Index", Title: "轮转列表", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
		{Id: 62, Name: "AddRoll", Title: "增加轮转", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
		{Id: 63, Name: "UpdateRoll", Title: "更新轮转", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},
		{Id: 64, Name: "DeleteRoll", Title: "删除轮转", Remark: "", Level: 3, Pid: 60, Status: 1, Group: g},

		//智干亚马逊美国数据
		{Id: 65, Name: "auas", Title: "基础数据", Remark: "", Level: 1, Pid: 0, Status: 1, Group: gb},
		{Id: 66, Name: "base/index", Title: "小类数据", Remark: "", Level: 2, Pid: 65, Status: 1, Group: gb},
		{Id: 67, Name: "Index", Title: "美国站小类数据列表", Remark: "", Level: 3, Pid: 66, Status: 1, Group: gb},
		{Id: 671, Name: "Query", Title: "美国站小类数据查询", Remark: "", Level: 3, Pid: 66, Status: 1, Group: gb},
		{Id: 672, Name: "Export", Title: "美国站小类数据导出", Remark: "", Level: 3, Pid: 66, Status: 1, Group: gb},

		{Id: 68, Name: "big/index", Title: "大类数据", Remark: "", Level: 2, Pid: 65, Status: 1, Group: gb},
		{Id: 69, Name: "Index", Title: "美国站大类数据列表", Remark: "", Level: 3, Pid: 68, Status: 1, Group: gb},
		{Id: 70, Name: "Query", Title: "美国站大类数据查询", Remark: "", Level: 3, Pid: 68, Status: 1, Group: gb},
		{Id: 71, Name: "Export", Title: "美国站大类数据导出", Remark: "", Level: 3, Pid: 68, Status: 1, Group: gb},
		{Id: 711, Name: "Asin", Title: "美国站Asin历史记录", Remark: "", Level: 3, Pid: 68, Status: 1, Group: gb},

		{Id: 72, Name: "asin/index", Title: "Asin数据", Remark: "", Level: 2, Pid: 65, Status: 1, Group: gb},
		{Id: 73, Name: "Index", Title: "美国站Asin数据列表", Remark: "", Level: 3, Pid: 72, Status: 1, Group: gb},
		{Id: 74, Name: "Query", Title: "美国站Asin数据查询", Remark: "", Level: 3, Pid: 72, Status: 1, Group: gb},
		{Id: 75, Name: "Export", Title: "美国站Asin数据导出", Remark: "", Level: 3, Pid: 72, Status: 1, Group: gb},

		{Id: 76, Name: "url/index", Title: "类目数据", Remark: "", Level: 2, Pid: 65, Status: 1, Group: gb},
		{Id: 77, Name: "Index", Title: "美国站类目数据列表", Remark: "", Level: 3, Pid: 76, Status: 1, Group: gb},
		{Id: 78, Name: "Query", Title: "美国站类目数据查询", Remark: "", Level: 3, Pid: 76, Status: 1, Group: gb},
		{Id: 79, Name: "Update", Title: "美国站类目数据更新", Remark: "", Level: 3, Pid: 76, Status: 1, Group: gb},

		{Id: 80, Name: "back", Title: "竞品数据", Remark: "", Level: 1, Pid: 0, Status: 1, Group: gb},
		{Id: 81, Name: "itemfind/index", Title: "选款数据", Remark: "", Level: 2, Pid: 80, Status: 1, Group: gb},
		{Id: 82, Name: "Index", Title: "美国站选款数据列表", Remark: "", Level: 3, Pid: 81, Status: 1, Group: gb},
		{Id: 83, Name: "Query", Title: "美国站选款数据查询", Remark: "", Level: 3, Pid: 81, Status: 1, Group: gb},
		{Id: 84, Name: "Export", Title: "美国站选款数据导出", Remark: "", Level: 3, Pid: 81, Status: 1, Group: gb},

		{Id: 85, Name: "keep/index", Title: "库存数据", Remark: "", Level: 2, Pid: 80, Status: 1, Group: gb},
		{Id: 86, Name: "Index", Title: "美国站库存数据列表", Remark: "", Level: 3, Pid: 85, Status: 1, Group: gb},
		{Id: 87, Name: "Query", Title: "美国站库存数据查询", Remark: "", Level: 3, Pid: 85, Status: 1, Group: gb},
		{Id: 88, Name: "Export", Title: "美国站库存数据导出", Remark: "", Level: 3, Pid: 85, Status: 1, Group: gb},

		{Id: 89, Name: "csv", Title: "业务数据", Remark: "", Level: 1, Pid: 0, Status: 1, Group: gb},
		{Id: 90, Name: "report/index", Title: "报告数据", Remark: "", Level: 2, Pid: 89, Status: 1, Group: gb},
		{Id: 91, Name: "Index", Title: "美国站报告数据列表", Remark: "", Level: 3, Pid: 90, Status: 1, Group: gb},
		{Id: 92, Name: "Query", Title: "美国站报告数据查询", Remark: "", Level: 3, Pid: 90, Status: 1, Group: gb},
		{Id: 93, Name: "Export", Title: "美国站报告数据导出", Remark: "", Level: 3, Pid: 90, Status: 1, Group: gb},
		{Id: 94, Name: "Import", Title: "美国站报告数据导入", Remark: "", Level: 3, Pid: 90, Status: 1, Group: gb},
		{Id: 95, Name: "Delete", Title: "美国站报告数据删除", Remark: "", Level: 3, Pid: 90, Status: 1, Group: gb},

		{Id: 96, Name: "code/index", Title: "折扣码数据", Remark: "", Level: 2, Pid: 89, Status: 1, Group: gb},
		{Id: 97, Name: "Index", Title: "美国站折扣码列表", Remark: "", Level: 3, Pid: 96, Status: 1, Group: gb},
		{Id: 98, Name: "Query", Title: "美国站折扣码查询", Remark: "", Level: 3, Pid: 96, Status: 1, Group: gb},
		{Id: 99, Name: "Export", Title: "美国站折扣码导出", Remark: "", Level: 3, Pid: 96, Status: 1, Group: gb},
		{Id: 100, Name: "Import", Title: "美国站折扣码导入", Remark: "", Level: 3, Pid: 96, Status: 1, Group: gb},
		{Id: 101, Name: "Delete", Title: "美国站折扣码数据删除", Remark: "", Level: 3, Pid: 96, Status: 1, Group: gb},

		//---------------------------------------------------
		//智干亚马逊日本数据
		{Id: 102, Name: "ajp", Title: "基础数据", Remark: "", Level: 1, Pid: 0, Status: 1, Group: jp},
		{Id: 103, Name: "base/index", Title: "小类数据", Remark: "", Level: 2, Pid: 102, Status: 1, Group: jp},
		{Id: 104, Name: "Index", Title: "日本站小类数据列表", Remark: "", Level: 3, Pid: 103, Status: 1, Group: jp},
		{Id: 105, Name: "Query", Title: "日本站小类数据查询", Remark: "", Level: 3, Pid: 103, Status: 1, Group: jp},
		{Id: 106, Name: "Export", Title: "日本站小类数据导出", Remark: "", Level: 3, Pid: 103, Status: 1, Group: jp},

		{Id: 107, Name: "big/index", Title: "大类数据", Remark: "", Level: 2, Pid: 102, Status: 1, Group: jp},
		{Id: 108, Name: "Index", Title: "日本站大类数据列表", Remark: "", Level: 3, Pid: 107, Status: 1, Group: jp},
		{Id: 109, Name: "Query", Title: "日本站大类数据查询", Remark: "", Level: 3, Pid: 107, Status: 1, Group: jp},
		{Id: 110, Name: "Export", Title: "日本站大类数据导出", Remark: "", Level: 3, Pid: 107, Status: 1, Group: jp},
		{Id: 712, Name: "Asin", Title: "日本站Asin历史记录", Remark: "", Level: 3, Pid: 107, Status: 1, Group: jp},

		{Id: 111, Name: "asin/index", Title: "Asin数据", Remark: "", Level: 2, Pid: 102, Status: 1, Group: jp},
		{Id: 112, Name: "Index", Title: "日本站Asin数据列表", Remark: "", Level: 3, Pid: 111, Status: 1, Group: jp},
		{Id: 113, Name: "Query", Title: "日本站Asin数据查询", Remark: "", Level: 3, Pid: 111, Status: 1, Group: jp},
		{Id: 114, Name: "Export", Title: "日本站Asin数据导出", Remark: "", Level: 3, Pid: 111, Status: 1, Group: jp},

		{Id: 115, Name: "url/index", Title: "类目数据", Remark: "", Level: 2, Pid: 102, Status: 1, Group: jp},
		{Id: 116, Name: "Index", Title: "日本站类目数据列表", Remark: "", Level: 3, Pid: 115, Status: 1, Group: jp},
		{Id: 117, Name: "Query", Title: "日本站类目数据查询", Remark: "", Level: 3, Pid: 115, Status: 1, Group: jp},
		{Id: 118, Name: "Update", Title: "日本站类目数据更新", Remark: "", Level: 3, Pid: 115, Status: 1, Group: jp},

		//---------------------------------------------------
		//智干亚马逊德国数据
		{Id: 119, Name: "de", Title: "基础数据", Remark: "", Level: 1, Pid: 0, Status: 1, Group: de},
		{Id: 120, Name: "base/index", Title: "小类数据", Remark: "", Level: 2, Pid: 119, Status: 1, Group: de},
		{Id: 121, Name: "Index", Title: "德国站小类数据列表", Remark: "", Level: 3, Pid: 120, Status: 1, Group: de},
		{Id: 122, Name: "Query", Title: "德国站小类数据查询", Remark: "", Level: 3, Pid: 120, Status: 1, Group: de},
		{Id: 123, Name: "Export", Title: "德国站小类数据导出", Remark: "", Level: 3, Pid: 120, Status: 1, Group: de},

		{Id: 124, Name: "big/index", Title: "大类数据", Remark: "", Level: 2, Pid: 119, Status: 1, Group: de},
		{Id: 125, Name: "Index", Title: "德国站大类数据列表", Remark: "", Level: 3, Pid: 124, Status: 1, Group: de},
		{Id: 126, Name: "Query", Title: "德国站大类数据查询", Remark: "", Level: 3, Pid: 124, Status: 1, Group: de},
		{Id: 127, Name: "Export", Title: "德国站大类数据导出", Remark: "", Level: 3, Pid: 124, Status: 1, Group: de},
		{Id: 128, Name: "Asin", Title: "德国站Asin历史记录", Remark: "", Level: 3, Pid: 124, Status: 1, Group: de},

		{Id: 129, Name: "asin/index", Title: "Asin数据", Remark: "", Level: 2, Pid: 119, Status: 1, Group: de},
		{Id: 130, Name: "Index", Title: "德国站Asin数据列表", Remark: "", Level: 3, Pid: 129, Status: 1, Group: de},
		{Id: 131, Name: "Query", Title: "德国站Asin数据查询", Remark: "", Level: 3, Pid: 129, Status: 1, Group: de},
		{Id: 132, Name: "Export", Title: "德国站Asin数据导出", Remark: "", Level: 3, Pid: 129, Status: 1, Group: de},

		{Id: 133, Name: "url/index", Title: "类目数据", Remark: "", Level: 2, Pid: 119, Status: 1, Group: de},
		{Id: 1331, Name: "Index", Title: "德国站类目数据列表", Remark: "", Level: 3, Pid: 133, Status: 1, Group: de},
		{Id: 1332, Name: "Query", Title: "德国站类目数据查询", Remark: "", Level: 3, Pid: 133, Status: 1, Group: de},
		{Id: 134, Name: "Update", Title: "德国站类目数据更新", Remark: "", Level: 3, Pid: 133, Status: 1, Group: de},

		//---------------------------------------------------
		//智干亚马逊英国数据
		{Id: 135, Name: "uk", Title: "基础数据", Remark: "", Level: 1, Pid: 0, Status: 1, Group: uk},
		{Id: 136, Name: "base/index", Title: "小类数据", Remark: "", Level: 2, Pid: 135, Status: 1, Group: uk},
		{Id: 137, Name: "Index", Title: "英国站小类数据列表", Remark: "", Level: 3, Pid: 136, Status: 1, Group: uk},
		{Id: 138, Name: "Query", Title: "英国站小类数据查询", Remark: "", Level: 3, Pid: 136, Status: 1, Group: uk},
		{Id: 139, Name: "Export", Title: "英国站小类数据导出", Remark: "", Level: 3, Pid: 136, Status: 1, Group: uk},

		{Id: 140, Name: "big/index", Title: "大类数据", Remark: "", Level: 2, Pid: 135, Status: 1, Group: uk},
		{Id: 141, Name: "Index", Title: "英国站大类数据列表", Remark: "", Level: 3, Pid: 140, Status: 1, Group: uk},
		{Id: 142, Name: "Query", Title: "英国站大类数据查询", Remark: "", Level: 3, Pid: 140, Status: 1, Group: uk},
		{Id: 143, Name: "Export", Title: "英国站大类数据导出", Remark: "", Level: 3, Pid: 140, Status: 1, Group: uk},
		{Id: 144, Name: "Asin", Title: "英国站Asin历史记录", Remark: "", Level: 3, Pid: 140, Status: 1, Group: uk},

		{Id: 145, Name: "asin/index", Title: "Asin数据", Remark: "", Level: 2, Pid: 135, Status: 1, Group: uk},
		{Id: 146, Name: "Index", Title: "英国站Asin数据列表", Remark: "", Level: 3, Pid: 145, Status: 1, Group: uk},
		{Id: 147, Name: "Query", Title: "英国站Asin数据查询", Remark: "", Level: 3, Pid: 145, Status: 1, Group: uk},
		{Id: 148, Name: "Export", Title: "英国站Asin数据导出", Remark: "", Level: 3, Pid: 145, Status: 1, Group: uk},

		{Id: 149, Name: "url/index", Title: "类目数据", Remark: "", Level: 2, Pid: 135, Status: 1, Group: uk},
		{Id: 150, Name: "Index", Title: "英国站类目数据列表", Remark: "", Level: 3, Pid: 149, Status: 1, Group: uk},
		{Id: 151, Name: "Query", Title: "英国站类目数据查询", Remark: "", Level: 3, Pid: 149, Status: 1, Group: uk},
		{Id: 152, Name: "Update", Title: "英国站类目数据更新", Remark: "", Level: 3, Pid: 149, Status: 1, Group: uk},
	}
	for _, v := range nodes {
		n := new(Node)
		n.Id = v.Id //这句是无效的
		n.Name = v.Name
		n.Title = v.Title
		n.Remark = v.Remark
		n.Level = v.Level
		n.Pid = v.Pid
		n.Status = v.Status
		n.Group = v.Group
		err := n.Insert()
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("insert node end")
}
