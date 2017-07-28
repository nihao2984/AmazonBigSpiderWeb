package rbac

import (
	"github.com/astaxie/beego"
	m "github.com/hunterhug/AmazonBigSpiderWeb/models/admin"
)

type CommonController struct {
	baseController
	Templatetype string //ui template type
}

func (this *CommonController) Prepare() {
	this.Data["version"] = beego.AppConfig.String("version")
}

// 请求状态响应
func (this *CommonController) Rsp(status bool, str string) {
	this.Data["json"] = &map[string]interface{}{"status": status, "info": str}
	this.ServeJSON()
	this.StopRun()
}

// 获取模板位置
func (this *CommonController) GetTemplate() string {
	templatetype := beego.AppConfig.String("admin_template")
	if templatetype == "" {
		templatetype = "default"
	}
	return templatetype
}

// 获取权限各节点URL   权限控制器 用户节点  /rbac /node/index
func (this *CommonController) GetTree() []Tree {
	nodes, _ := m.GetNodeTree(0, 1) //第一层
	tree := make([]Tree, len(nodes))
	for k, v := range nodes {
		tree[k].Id = v["Id"].(int64)
		tree[k].Text = v["Title"].(string)
		tree[k].GroupId = v["Group"].(int64)
		children, _ := m.GetNodeTree(v["Id"].(int64), 2) //第二层
		tree[k].Children = make([]Tree, len(children))
		for k1, v1 := range children {
			tree[k].Children[k1].Id = v1["Id"].(int64)
			tree[k].Children[k1].Text = v1["Title"].(string)
			tree[k].Children[k1].Attributes.Url = "/" + v["Name"].(string) + "/" + v1["Name"].(string)
		}
	}
	// beego.Trace("%v", tree)
	return tree
}
