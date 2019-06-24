package admin

import (
	"coinford_admin_api/admin_models"
	"github.com/astaxie/beego"
	"encoding/json"
)

type UserController struct {
	beego.Controller
}
/*
// @Title GetAll
// @Description get the list of users
// @Param	body		body 	UserAdd		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /add [post]
func (u *UserController) Add() {
	var rqd UserAdd
	json.Unmarshal(u.Ctx.Input.RequestBody, &rqd)
	jres, admin, adminGroup, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		if adminGroup.Key == "SUPER_ADMIN" {
			user := admin_models.User{Name: rqd.Name, IsoCode: rqd.IsoCode, DialCode: rqd.DialCode, Code: rqd.Code}
			err := user.Insert()
			if err == nil {
				u.Data["json"] = u.prepareUserList(admin)
			} else {
				debugMessage("User Add: ", err)
				u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to add user."}
			}
		} else {
			u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Admin not authorized."}
		}
	} else {
		u.Data["json"] = jres
	}
	u.ServeJSON()
}

// @Title GetAll
// @Description get the list of users
// @Param	body		body 	UserUpdate		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /update [post]
func (u *UserController) Update() {//adminGroup
	var rqd UserUpdate
	json.Unmarshal(u.Ctx.Input.RequestBody, &rqd)
	jres, admin, adminGroup, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		if adminGroup.Key == "SUPER_ADMIN" {
			user := admin_models.User{Id: rqd.Id}
			err := user.Read("id")
			if err == nil {
				user.Name = rqd.Name
				user.IsoCode = rqd.IsoCode
				user.DialCode = rqd.DialCode
				user.Code = rqd.Code
				err = user.Update()
				if err == nil {
					u.Data["json"] = u.prepareUserList(admin)
				} else {
					debugMessage("User Update: ", err)
					u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to update user."}
				}
			} else {
				u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Unable to read and write data."}
			}
		} else {
			u.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Admin not authorized."}
		}
	} else {
		u.Data["json"] = jres
	}
	u.ServeJSON()
}
*/
// @Title GetAll
// @Description get the list of users
// @Param	body		body 	UserGet		true		"Token for Authentication"
// @Success 200 {int} response
// @Failure 403 Authentication Failed
// @router /list [post]
func (u *UserController) GetAll() {//adminGroup
	var rqd UserGet
	json.Unmarshal(u.Ctx.Input.RequestBody, &rqd)
	jres, admin, _, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		u.Data["json"] = u.prepareUserList(admin, &rqd.UserSearch)
	} else {
		u.Data["json"] = jres
	}
	u.ServeJSON()
}

func (u *UserController) prepareUserList(admin *admin_models.Admin, userSearch *UserSearch) UserList {
	_, users, _ := admin_models.Users(admin, userSearch.CountryId, userSearch.Eligible)
	var userDetails []UserDetail
	for _, user := range users {
		userDetail := UserDetail{User: user}
		_, countries, _ := admin_models.CountriesUser(admin, &user, userSearch.Eligible)
		var userCountryDetails []UserCountryDetail
		for _, country := range countries {
			userCountry := admin_models.UserCountry{UserId: user.Id, CountryId: country.Id}
			userCountry.Read("user_id", "country_id")
			userCountryDetail := UserCountryDetail{Country: country, UserCountry: userCountry}
			userCountryDetails = append(userCountryDetails, userCountryDetail)
		}
		userDetail.UserCountryDetail = userCountryDetails
		userDetails = append(userDetails, userDetail)
	}
	return UserList{UserDetail: userDetails, ResponseCode: 200, ResponseDescription: "List of Users"}
}