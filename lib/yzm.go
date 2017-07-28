package lib

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/utils/captcha"
)

var cpt *captcha.Captcha

func init() {
	store := cache.NewMemoryCache()
	cpt = captcha.NewWithFilter("/captcha/", store) //一定要写在构造函数里面，要不然第一次打开页面有可能是X
	//原数据
	cpt.FieldIDName = "yzm_id"

	//用户数据
	cpt.FieldCaptchaName = "yzm"
	cpt.ChallengeNums = 4
	// cpt.StdWidth = 100
	// cpt.StdHeight = 40
}

func Verify(this *context.Context) bool {
	if cpt.VerifyReq(this.Request) {
		return true
	}
	return false
}
