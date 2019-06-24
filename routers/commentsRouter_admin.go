package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/register`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"],
		beego.ControllerComments{
			Method: "AdminGroups",
			Router: `/admingroups`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/delete`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"],
		beego.ControllerComments{
			Method: "ChangeAdminname",
			Router: `/changeadminname`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"],
		beego.ControllerComments{
			Method: "ChangeName",
			Router: `/changename`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"],
		beego.ControllerComments{
			Method: "IssueToken",
			Router: `/token`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"],
		beego.ControllerComments{
			Method: "Login",
			Router: `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"],
		beego.ControllerComments{
			Method: "Authenticate",
			Router: `/auth`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"],
		beego.ControllerComments{
			Method: "Logout",
			Router: `/logout`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:AdminController"],
		beego.ControllerComments{
			Method: "IsUniqueAdminname",
			Router: `/isuniqueadminname`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:CountryController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:CountryController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:CountryController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:CountryController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:CountryController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:CountryController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:CurrencyController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:CurrencyController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:CurrencyController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:CurrencyController"],
		beego.ControllerComments{
			Method: "Update",
			Router: `/update`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:CurrencyController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:CurrencyController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["coinford_admin_api/admin:UserController"] = append(beego.GlobalControllerRouter["coinford_admin_api/admin:UserController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/list`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
