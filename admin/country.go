package admin

import (
	"coinford_admin_api/admin_models"
	"github.com/astaxie/beego"
	"encoding/json"
)

type CountryController struct {
	beego.Controller
}

// @Title GetAll
// @Description get the list of countries
// @Param	body		body 	CountryAdd		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /add [post]
func (c *CountryController) Add() {
	var rqd CountryAdd
	json.Unmarshal(c.Ctx.Input.RequestBody, &rqd)
	jres, admin, adminGroup, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		if adminGroup.Key == "SUPER_ADMIN" {
			country := admin_models.Country{Name: rqd.Name, IsoCode: rqd.IsoCode, DialCode: rqd.DialCode, Code: rqd.Code}
			err := country.Insert()
			if err == nil {
				c.Data["json"] = c.prepareCountryList(admin)
			} else {
				debugMessage("Country Add: ", err)
				c.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to add country."}
			}
		} else {
			c.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Admin not authorized."}
		}
	} else {
		c.Data["json"] = jres
	}
	c.ServeJSON()
}

// @Title GetAll
// @Description get the list of countries
// @Param	body		body 	CountryUpdate		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /update [post]
func (c *CountryController) Update() {//adminGroup
	var rqd CountryUpdate
	json.Unmarshal(c.Ctx.Input.RequestBody, &rqd)
	jres, admin, adminGroup, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		if adminGroup.Key == "SUPER_ADMIN" {
			country := admin_models.Country{Id: rqd.Id}
			err := country.Read("id")
			if err == nil {
				country.Name = rqd.Name
				country.IsoCode = rqd.IsoCode
				country.DialCode = rqd.DialCode
				country.Code = rqd.Code
				err = country.Update()
				if err == nil {
					c.Data["json"] = c.prepareCountryList(admin)
				} else {
					debugMessage("Country Update: ", err)
					c.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to update country."}
				}
			} else {
				c.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to read and write data."}
			}
		} else {
			c.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Admin not authorized."}
		}
	} else {
		c.Data["json"] = jres
	}
	c.ServeJSON()
}

// @Title GetAll
// @Description get the list of countries
// @Param	body		body 	AdminToken		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /list [post]
func (c *CountryController) GetAll() {//adminGroup
	var rqd AdminToken
	json.Unmarshal(c.Ctx.Input.RequestBody, &rqd)
	jres, admin, _, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		c.Data["json"] = c.prepareCountryList(admin)
	} else {
		c.Data["json"] = jres
	}
	c.ServeJSON()
}
func (c *CountryController) prepareCountryList(admin *admin_models.Admin) CountryList {
	_, countries, _ := admin_models.Countries(admin)
	return CountryList{Countries: countries, ResponseCode: 200, ResponseDescription: "List of Countries"}
}