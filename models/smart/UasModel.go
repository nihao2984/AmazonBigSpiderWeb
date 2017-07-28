package smart

import (
	"github.com/astaxie/beego/orm"
	"time"
)
// 此模型报废，由于表名是不确定的!!!!
var (
	o = orm.NewOrm()
)

/*CREATE TABLE `smartdb`.`20161028` (
  `id` VARCHAR(255),
  `purl` varchar(255) DEFAULT NULL COMMENT '父类类目链接',
  `dbname` varchar(255) DEFAULT NULL COMMENT 'dbname',
  `col1` varchar(255) DEFAULT NULL COMMENT '预留字段',
  `col2` varchar(255) DEFAULT NULL,
  `col3` varchar(255) DEFAULT NULL,
  `iscatch` tinyint(4) DEFAULT '0' COMMENT '已抓取是1',
  `smallrank` INT NULL COMMENT '小类排名',
  `name` VARCHAR(255) NULL COMMENT '小类名',
  `bigname` VARCHAR(255) NULL COMMENT '大类名',
  `title` TINYTEXT NULL COMMENT '商品标题',
  `asin` VARCHAR(255) NULL,
  `url` VARCHAR(255) NULL,
  `rank` INT NULL COMMENT '大类排名',
  `soldby` VARCHAR(255) NULL COMMENT '卖家',
  `shipby` VARCHAR(255) NULL COMMENT '物流',
  `price` FLOAT NULL COMMENT '价格',
  `score` FLOAT NULL COMMENT '打分',
  `commentnum` INT NULL COMMENT '评论数',
  `commenttime` VARCHAR(255) NULL COMMENT '第一条评论时间',
  `createtime` DATETIME NULL,
  PRIMARY KEY (`id`)ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='类目表';*/

type UsaBase struct {
	Id          string `orm:"pk"`
	Purl        string
	Dbname      string
	Col1        string
	Col2        string
	Col3        string
	Iscatch     int64 `orm:"default(0)"`
	Smallrank   int64
	Name        string
	Bigname     string
	Title       string
	Asin        string
	Url         string
	Rank        int64
	Soldby      string
	Shipby      string
	Price       float64
	Score       float64
	Commentnum  int64
	Commenttime string
	Createtime  time.Time `orm:"type(datetime);null"`
}

func init() {
	orm.RegisterModel(new(UsaBase))
	o.Using("smartdb")
}

func (m *UsaBase) Read(fields ...string) error {
	if err := o.Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *UsaBase) Update(fields ...string) error {
	if _, err := o.Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *UsaBase) Delete() error {
	if _, err := o.Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *UsaBase) Query() orm.QuerySeter {
	return o.QueryTable(m)
}

func (m *UsaBase) Insert() error {
	if _, err := o.Insert(m); err != nil {
		return err
	}
	return nil
}
