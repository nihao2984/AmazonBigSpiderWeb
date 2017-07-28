package blog

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Roll struct {
	Id         int64
	Title      string    `orm:"size(100)"`       //标题
	Content    string    `orm:"type(text)";null` //内容
	Createtime time.Time `orm:"type(datetime);null"`
	Updatetime time.Time `orm:"type(datetime);null"`
	Sort       int64     //排序
	Status     int64     `orm:"default(0)"` //0 关闭 1开启
	Photo      string    //图片加密地址
	View       int64     //浏览量
	Url        string
}

func init() {
	orm.RegisterModel(new(Roll))
}

func (m *Roll) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Roll) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Roll) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Roll) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Roll) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
