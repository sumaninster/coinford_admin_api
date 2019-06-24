package admin_models

import (
	"github.com/astaxie/beego/orm"
	"coinford_admin_api/admin_configs"
	"fmt"
)

func (ac *AdminCountry) TableName() string {
    return "admin_country"
}

func (ac *AdminCountry) Insert() error {
	if _, err := orm.NewOrm().Insert(ac); err != nil {
		return err
	}
	return nil
}

func (ac *AdminCountry) Read(fields ...string) error {
	if err := orm.NewOrm().Read(ac, fields...); err != nil {
		return err
	}
	return nil
}

func (ac *AdminCountry) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(ac, field, fields...)
}

func (ac *AdminCountry) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(ac, fields...); err != nil {
		return err
	}
	return nil
}

func (ac *AdminCountry) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(ac, fields ...); err != nil {
		return err
	}
	return nil
}

func AdminCountryIds(adminId int64) []interface{} {
	num, admin_countries, err := AdminCountries(adminId)
	country_ids := []int64{admin_configs.GLOBAL_CODE}
	if num > 0 && err == nil {
		for _, v := range admin_countries {
		    country_ids = append(country_ids, v.CountryId)
		}
	}
	icountry_ids := admin_configs.Int64ToInterface(country_ids)
	return icountry_ids
}

func IsAdminEligible(adminId int64, countryId int64) bool {
	var table AdminCountry
	var admincountries []AdminCountry
	num, err := orm.NewOrm().QueryTable(table).Filter("admin_id", adminId).Filter("country_id", countryId).Filter("deleted_at__isnull", true).All(&admincountries)
	if num == 1 && err == nil{
		return true
	}
	num, err = orm.NewOrm().QueryTable(table).Filter("admin_id", adminId).Filter("country_id", admin_configs.GLOBAL_CODE).Filter("deleted_at__isnull", true).All(&admincountries)
	fmt.Println(num, err)
	if num == 1 && err == nil{
		return true
	}
	return false
}

func AdminCountries(adminId int64) (int64, []AdminCountry, error) {
	var table AdminCountry
	var admincountries []AdminCountry
	var num int64
	var err error
	num, err = orm.NewOrm().QueryTable(table).Filter("admin_id", adminId).Filter("deleted_at__isnull", true).OrderBy("country_id").All(&admincountries)
	return num, admincountries, err
}

func init() {
	orm.RegisterModel(new(AdminCountry))
	
}
