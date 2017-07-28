package blog

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Paper struct {
	Id         int64
	Title      string    `orm:"size(100)"`       //标题
	Content    string    `orm:"type(text)";null` //内容
	Descontent string    `orm:"type(text);null"` //内容简介
	Createtime time.Time `orm:"type(datetime);null"`
	Updatetime time.Time `orm:"type(datetime);null"`
	Sort       int64     //排序
	Status     int64     `orm:"default(0)"` //0 未审核 1审核 2回收站
	Author   string    //昵称
	Photo      string    //图片加密地址
	View       int64     //浏览量
	Cid        int64     //分类
	Istop     int64     `orm:"default(0)"` //是否置顶 1置顶
	Ishot      int64     `orm:"default(0)"` //是否热门 1热门
	Isroll int64  `orm:"default(0)"` //是否轮转
	Rollpath string //自定义轮转地址
	Type int64 //0表示文章，1表示图片
}

func init() {
	orm.RegisterModel(new(Paper))
}

func (m *Paper) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}


func (m *Paper) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Paper) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Paper) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Paper) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
