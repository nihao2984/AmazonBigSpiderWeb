/*
	版权所有，侵权必究
	署名-非商业性使用-禁止演绎 4.0 国际
	警告： 以下的代码版权归属hunterhug，请不要传播或修改代码
	你可以在教育用途下使用该代码，但是禁止公司或个人用于商业用途(在未授权情况下不得用于盈利)
	商业授权请联系邮箱：gdccmcm14@live.com QQ:459527502

	All right reserved
	Attribution-NonCommercial-NoDerivatives 4.0 International
	Notice: The following code's copyright by hunterhug, Please do not spread and modify.
	You can use it for education only but can't make profits for any companies and individuals!
	For more information on commercial licensing please contact hunterhug.
	Ask for commercial licensing please contact Mail:gdccmcm14@live.com Or QQ:459527502

	2017.7 by hunterhug
*/
package routers

import (
	"github.com/astaxie/beego"
	"github.com/hunterhug/AmazonBigSpiderWeb/controllers/smart"
	"github.com/hunterhug/AmazonBigSpiderWeb/controllers/smartjp"
	"github.com/hunterhug/AmazonBigSpiderWeb/controllers/smartuk"
	"github.com/hunterhug/AmazonBigSpiderWeb/controllers/smartde"
)

func smartrouter() {

	//base data
	beego.Router("/auas/base/index", &smart.UsaController{}, "*:Index")
	beego.Router("/auas/base/query", &smart.UsaController{}, "*:Query")
	//beego.Router("/auas/base/export", &smart.UsaController{}, "*:Export")

	//big rank data
	beego.Router("/auas/big/index", &smart.BigController{}, "*:Index")
	beego.Router("/auas/big/query", &smart.BigController{}, "*:Query")
	beego.Router("/auas/big/export", &smart.BigController{}, "*:Export")
	beego.Router("/auas/big/asin", &smart.BigController{}, "*:Asin")

	//asin  data
	beego.Router("/auas/asin/index", &smart.AsinController{}, "*:Index")
	beego.Router("/auas/asin/query", &smart.AsinController{}, "*:Query")
	//beego.Router("/auas/asin/export", &smart.AsinController{}, "*:Export")

	//url  data
	beego.Router("/auas/url/index", &smart.UrlController{}, "*:Index")
	beego.Router("/auas/url/query", &smart.UrlController{}, "*:Query")
	beego.Router("/auas/url/update", &smart.UrlController{}, "*:Update")
	//beego.Router("/auas/url/export", &smart.UrlController{}, "*:Export")

	//itemfind
	beego.Router("/back/itemfind/index", &smart.ItemFindController{}, "*:Index")
	beego.Router("/back/itemfind/query", &smart.ItemFindController{}, "*:Query")
	//beego.Router("/back/itemfind/export", &smart.ItemFindController{}, "*:Export")

	beego.Router("/back/keep/index", &smart.KeepController{}, "*:Index")
	beego.Router("/back/keep/query", &smart.KeepController{}, "*:Query")
	beego.Router("/back/keep/export", &smart.KeepController{}, "*:Export")

	beego.Router("/csv/report/index", &smart.ReportController{}, "*:Index")
	beego.Router("/csv/report/query", &smart.ReportController{}, "*:Query")
	beego.Router("/csv/report/export", &smart.ReportController{}, "*:Export")
	beego.Router("/csv/report/import", &smart.ReportController{}, "*:Import")
	beego.Router("/csv/report/delete", &smart.ReportController{}, "*:Delete")

	//japan
	beego.Router("/ajp/base/index", &smartjp.UsaController{}, "*:Index")
	beego.Router("/ajp/base/query", &smartjp.UsaController{}, "*:Query")

	//big rank data
	beego.Router("/ajp/big/index", &smartjp.BigController{}, "*:Index")
	beego.Router("/ajp/big/query", &smartjp.BigController{}, "*:Query")
	beego.Router("/ajp/big/export", &smartjp.BigController{}, "*:Export")
	beego.Router("/ajp/big/asin", &smartjp.BigController{}, "*:Asin")

	//asin  data
	beego.Router("/ajp/asin/index", &smartjp.AsinController{}, "*:Index")
	beego.Router("/ajp/asin/query", &smartjp.AsinController{}, "*:Query")

	//url  data
	beego.Router("/ajp/url/index", &smartjp.UrlController{}, "*:Index")
	beego.Router("/ajp/url/query", &smartjp.UrlController{}, "*:Query")
	beego.Router("/ajp/url/update", &smartjp.UrlController{}, "*:Update")

	//uk
	beego.Router("/uk/base/index", &smartuk.UsaController{}, "*:Index")
	beego.Router("/uk/base/query", &smartuk.UsaController{}, "*:Query")

	//big rank data
	beego.Router("/uk/big/index", &smartuk.BigController{}, "*:Index")
	beego.Router("/uk/big/query", &smartuk.BigController{}, "*:Query")
	beego.Router("/uk/big/export", &smartuk.BigController{}, "*:Export")
	beego.Router("/uk/big/asin", &smartuk.BigController{}, "*:Asin")

	//asin  data
	beego.Router("/uk/asin/index", &smartuk.AsinController{}, "*:Index")
	beego.Router("/uk/asin/query", &smartuk.AsinController{}, "*:Query")

	//url  data
	beego.Router("/uk/url/index", &smartuk.UrlController{}, "*:Index")
	beego.Router("/uk/url/query", &smartuk.UrlController{}, "*:Query")
	beego.Router("/uk/url/update", &smartuk.UrlController{}, "*:Update")

	//de
	beego.Router("/de/base/index", &smartde.UsaController{}, "*:Index")
	beego.Router("/de/base/query", &smartde.UsaController{}, "*:Query")

	//big rank data
	beego.Router("/de/big/index", &smartde.BigController{}, "*:Index")
	beego.Router("/de/big/query", &smartde.BigController{}, "*:Query")
	beego.Router("/de/big/export", &smartde.BigController{}, "*:Export")
	beego.Router("/de/big/asin", &smartde.BigController{}, "*:Asin")

	//asin  data
	beego.Router("/de/asin/index", &smartde.AsinController{}, "*:Index")
	beego.Router("/de/asin/query", &smartde.AsinController{}, "*:Query")

	//url  data
	beego.Router("/de/url/index", &smartde.UrlController{}, "*:Index")
	beego.Router("/de/url/query", &smartde.UrlController{}, "*:Query")
	beego.Router("/de/url/update", &smartde.UrlController{}, "*:Update")
}
