package blog

import (
	"github.com/astaxie/beego/orm"
)

type Config struct {
	Id         int64
	Title      string    `orm:"size(100)"`       //标题
	Content    string    `orm:"type(text)";null` //网站描述
	Address string    `orm:"type(text);null"` //地址
	Phone string //联系方式
	Webinfo string    `orm:"type(text);null"` //备案信息
	Photo string //logo
	Slogan string    `orm:"type(text);null"` //漂移通知
	Code1 string    `orm:"type(text);null"`  //评论短域名
	Code2 string    `orm:"type(text);null"`  //统计代码
	Code3 string    `orm:"type(text);null"` //友情链接
}

func init() {
	orm.RegisterModel(new(Config))
}

func (m *Config) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}


func (m *Config) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Config) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Config) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}

func (m *Config) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}
