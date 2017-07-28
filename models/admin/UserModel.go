package admin

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	. "github.com/hunterhug/AmazonBigSpiderWeb/lib"
)

type User struct {
	Id            int64
	Logincount    int
	Username      string    `orm:"unique;size(32)" form:"Username"  valid:"Required;MaxSize(20);MinSize(6)"`
	Password      string    `orm:"size(32)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)"`
	Repassword    string    `orm:"-" form:"Repassword" valid:"Required"`
	Nickname      string    `orm:"unique;size(32)" form:"Nickname" valid:"Required;MaxSize(20);MinSize(2)"`
	Email         string    `orm:"size(32)" form:"Email" valid:"Email"`
	Remark        string    `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
	Status        int       `orm:"default(2)" form:"Status" valid:"Range(1,2)"`
	Lastlogintime time.Time `orm:"null;type(datetime)" form:"-"`
	Createtime    time.Time `orm:"type(datetime)" `
	Lastip        string
	Role          []*Role `orm:"rel(m2m)"`
}

func (u *User) TableName() string {
	return beego.AppConfig.String("rbac_user_table")
}

func (u *User) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "两次输入的密码不一样")
	}
}

// 验证表单
func checkUser(u *User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(u)
	if !b {
		for _, err := range valid.Errors {
			return errors.New(err.Message)
		}
	}
	return nil
}

func init() {
	orm.RegisterModel(new(User))
}

// 列出用户
func Getuserlist(page int64, page_size int64, sort string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&users, "Createtime", "Email", "Id", "Lastip",
		"Lastlogintime", "Logincount", "Nickname", "Remark", "Status", "Username")
	count, _ = qs.Count()
	return users, count
}

func AddUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := new(User)
	user.Username = u.Username
	user.Password = Strtomd5(u.Password)
	user.Nickname = u.Nickname
	user.Email = u.Email
	user.Remark = u.Remark
	user.Status = u.Status
	user.Lastip = u.Lastip
	user.Createtime = GetTime()
	id, err := o.Insert(user)
	return id, err
}

func UpdateUser(u *User) (int64, error) {
	o := orm.NewOrm()
	user := make(orm.Params)
	valid := validation.Validation{}
	if len(u.Nickname) > 0 {
		valid.MinSize(u.Nickname, 6, "昵称最小长度")
		valid.MaxSize(u.Nickname, 20, "昵称最大长度")
		user["Nickname"] = u.Nickname
	}
	if len(u.Email) > 0 {
		valid.Email(u.Email, "邮箱")
		user["Email"] = u.Email
	}
	if len(u.Remark) > 0 {
		valid.MaxSize(u.Remark, 200, "最长备注")
		user["Remark"] = u.Remark
	}
	if u.Status != 0 {
		user["Status"] = u.Status
	}
	if len(user) == 0 {
		return 0, errors.New("字段不能为空")
	}
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			return 0, errors.New(err.Key + ":" + err.Message)
		}
	}
	var table User
	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
	return num, err
}

func UpdateUserPasswd(u *User) (int64, error) {
	valid := validation.Validation{}
	o := orm.NewOrm()
	user := make(orm.Params)
	if len(u.Password) > 0 {
		valid.MinSize(u.Password, 6, "最小长度")
		valid.MaxSize(u.Password, 20, "最大长度")
		user["Password"] = Strtomd5(u.Password)
	}
	if len(user) == 0 {
		return 0, errors.New("字段不能为空")
	}
	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			return 0, errors.New(err.Key + ":" + err.Message)
		}
	}
	var table User
	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
	return num, err
}

func UpdateLoginTime(u *User) User {
	u.Lastlogintime = GetTime()
	o := orm.NewOrm()
	o.Update(u)
	return *u
}

func DelUserById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&User{Id: Id})
	return status, err
}

func GetUserByUsername(username string) (user User) {
	user = User{Username: username}
	o := orm.NewOrm()
	o.Read(&user, "Username")
	return user
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
