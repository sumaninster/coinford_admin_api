package admin

import (
	"coinford_admin_api/admin_models"
	"github.com/astaxie/beego"
	"encoding/json"
)

type CurrencyController struct {
	beego.Controller
}

// @Title GetAll
// @Description get the list of currencies
// @Param	body		body 	CurrencyAdd		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /add [post]
func (c *CurrencyController) Add() {
	var rqd CurrencyAdd
	json.Unmarshal(c.Ctx.Input.RequestBody, &rqd)
	jres, admin, adminGroup, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		if adminGroup.Key == "SUPER_ADMIN" {
			currency := admin_models.Currency{Name: rqd.Name, Code: rqd.Code, Description: rqd.Description, Type: rqd.Type, CountryId: rqd.CountryId}
			err := currency.Insert()
			if err == nil {
				c.Data["json"] = c.prepareCurrencyList(admin)
			} else {
				debugMessage("Currency Add: ", err)
				c.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to add currency."}
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
// @Description get the list of currencies
// @Param	body		body 	CurrencyUpdate		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /update [post]
func (c *CurrencyController) Update() {//adminGroup
	var rqd CurrencyUpdate
	json.Unmarshal(c.Ctx.Input.RequestBody, &rqd)
	jres, admin, adminGroup, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		if adminGroup.Key == "SUPER_ADMIN" {
			currency := admin_models.Currency{Id: rqd.Id}
			err := currency.Read("id")
			if err == nil {
				currency.Name = rqd.Name
				currency.Code = rqd.Code
				currency.Description = rqd.Description
				currency.Type = rqd.Type
				currency.CountryId = rqd.CountryId
				err = currency.Update()
				if err == nil {
					c.Data["json"] = c.prepareCurrencyList(admin)
				} else {
					debugMessage("Currency Update: ", err)
					c.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to update currency."}
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
// @Description get the list of currencies
// @Param	body		body 	AdminToken		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /list [post]
func (c *CurrencyController) GetAll() {//adminGroup
	var rqd AdminToken
	json.Unmarshal(c.Ctx.Input.RequestBody, &rqd)
	jres, admin, _, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		c.Data["json"] = c.prepareCurrencyList(admin)
	} else {
		c.Data["json"] = jres
	}
	c.ServeJSON()
}
func (c *CurrencyController) prepareCurrencyList(admin *admin_models.Admin) CurrencyList {
	_, currencies, _ := admin_models.Currencies(admin)
	return CurrencyList{Currencies: currencies, ResponseCode: 200, ResponseDescription: "List of Currencies"}
}