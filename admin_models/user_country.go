package admin_models

import (
	"github.com/astaxie/beego/orm"
	"coinford_admin_api/admin_configs"
	//"fmt"
)

func (uc *UserCountry) TableName() string {
    return "user_country"
}

func (uc *UserCountry) Insert() error {
	if _, err := orm.NewOrm().Insert(uc); err != nil {
		return err
	}
	return nil
}

func (uc *UserCountry) Read(fields ...string) error {
	if err := orm.NewOrm().Read(uc, fields...); err != nil {
		return err
	}
	return nil
}

func (uc *UserCountry) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(uc, field, fields...)
}

func (uc *UserCountry) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(uc, fields...); err != nil {
		return err
	}
	return nil
}

func (uc *UserCountry) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(uc, fields ...); err != nil {
		return err
	}
	return nil
}

func UserCountryIds(admin *Admin, user *User, eligible string) []interface{} {
	num, user_countries, err := UserCountries(user, eligible)
	country_ids := []int64{}//admin_configs.GLOBAL_CODE
	if num > 0 && err == nil {
		for _, v := range user_countries {
		    country_ids = append(country_ids, v.CountryId)
		}
	}
	icountry_ids := admin_configs.Int64ToInterface(country_ids)
	return icountry_ids
}

func IsUserEligible(user *User, countryId int64) bool {
	var table UserCountry
	var usercountries []UserCountry
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", user.Id).Filter("country_id", countryId).Filter("eligible", "YES").Filter("deleted_at__isnull", true).All(&usercountries)
	if num == 1 && err == nil{
		return true
	}
	return false
}

func UserCountries(user *User, eligible string) (int64, []UserCountry, error) {
	var table UserCountry
	var usercountries []UserCountry
	var num int64
	var err error
	if eligible == "YES" || eligible == "NO" {
		num, err = orm.NewOrm().QueryTable(table).Filter("user_id", user.Id).Filter("eligible", eligible).Filter("deleted_at__isnull", true).OrderBy("country_id").All(&usercountries)
	} else {
		num, err = orm.NewOrm().QueryTable(table).Filter("user_id", user.Id).Filter("deleted_at__isnull", true).OrderBy("country_id").All(&usercountries)
	}
	return num, usercountries, err
}

func AdminCountryIdsUser(admin *Admin, countryId int64, eligible string) (int64, []interface{}) {
	num, user_countries, err := AdminCountriesUser(admin, countryId, eligible)
	user_ids := []int64{}
	if num > 0 && err == nil {
		for _, v := range user_countries {
		    user_ids = append(user_ids, v.UserId)
		}
	}
	iuser_ids := admin_configs.Int64ToInterface(user_ids)
	return num, iuser_ids
}

func AdminCountriesUser(admin *Admin, countryId int64, eligible string) (int64, []UserCountry, error) {
	var table UserCountry
	var usercountries []UserCountry
	var num int64
	var err error
	adminGroup := GetAdminGroup(admin)
	if adminGroup.Key == "SUPER_ADMIN" || adminGroup.Key == "ADMIN" || adminGroup.Key == "CUSTOMER_SUPPORT" {
		isAdminEligible := IsAdminEligible(admin.Id, countryId)
		if isAdminEligible {
			o := orm.NewOrm()
			qs := o.QueryTable(table)
			cond := orm.NewCondition()
			var cond1 *orm.Condition
			if countryId != admin_configs.GLOBAL_CODE {
				cond1 = cond.And("country_id__in", countryId)
			}

			if eligible == "YES" || eligible == "NO" {
				cond1 = cond.And("eligible", eligible).AndCond(cond1)
			}

			num, err = qs.SetCond(cond1).Filter("deleted_at__isnull", true).OrderBy("-created_at").OrderBy("country_id").All(&usercountries)
		}
	}
	return num, usercountries, err
}

func init() {
	orm.RegisterModel(new(UserCountry))
	
}
