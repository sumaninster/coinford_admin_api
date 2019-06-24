package admin

import (
	"coinford_admin_api/admin_models"
)

type CommonResponse struct {
	ResponseCode 			int
	ResponseDescription 	string
}

type LoginResponse struct {
	Token 					string
	Name 					string
	AdminGroup 				admin_models.AdminGroup
	ResponseCode 			int
	ResponseDescription 	string
}

type AdminList struct {
	Admins 					[]admin_models.Admin
	ResponseCode 			int
	ResponseDescription 	string
}

type AdminGroupList struct {
	AdminGroups 			[]admin_models.AdminGroup
	ResponseCode 			int
	ResponseDescription 	string
}

type CountryList struct {
	Countries 				[]admin_models.Country
	ResponseCode 			int
	ResponseDescription 	string
}

type CurrencyList struct {
	Currencies 				[]admin_models.Currency
	ResponseCode 			int
	ResponseDescription 	string
}

type UserCountryDetail struct {
	Country 				admin_models.Country
	UserCountry 			admin_models.UserCountry
}

type UserDetail struct {
	User 					admin_models.User
	UserCountryDetail 		[]UserCountryDetail
}

type UserList struct {
	UserDetail 				[]UserDetail
	ResponseCode 			int
	ResponseDescription 	string
}