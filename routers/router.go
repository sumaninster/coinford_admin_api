// @APIVersion 1.0.0
// @Title Coinford Admin API
// @Description Coinford Admin API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://coinford.com/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"coinford_admin_api/admin"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	    AllowAllOrigins: true,
	    AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	    AllowHeaders:     []string{"Origin", "Authorization", "X-Requested-With"},//, "Access-Control-Allow-Origin"
	    ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},//, "Access-Control-Allow-Headers"
	}))
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/admin",
			beego.NSInclude(
				&admin.AdminController{},
			),
		),
		beego.NSNamespace("/country",
			beego.NSInclude(
				&admin.CountryController{},
			),
		),
		beego.NSNamespace("/currency",
			beego.NSInclude(
				&admin.CurrencyController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&admin.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
