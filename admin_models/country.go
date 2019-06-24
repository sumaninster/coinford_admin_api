package admin_models

import (
	"github.com/astaxie/beego/orm"
	//"fmt"
)

func (c *Country) TableName() string {
    return "country"
}

func (c *Country) Insert() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

func (c *Country) Read(fields ...string) error {
	if err := orm.NewOrm().Read(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Country) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(c, field, fields...)
}

func (c *Country) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Country) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(c, fields ...); err != nil {
		return err
	}
	return nil
}

func Countries(admin *Admin) (int64, []Country, error) {
	var table Country
	var countries []Country
	var num int64
	var err error
	adminGroup := GetAdminGroup(admin)
	if adminGroup.Key == "SUPER_ADMIN" || adminGroup.Key == "ADMIN" || adminGroup.Key == "CUSTOMER_SUPPORT" {
		num, err = orm.NewOrm().QueryTable(table).Filter("deleted_at__isnull", true).OrderBy("name").All(&countries)
	}
	return num, countries, err
}

func CountriesUser(admin *Admin, user *User, eligible string) (int64, []Country, error) {
	var table Country
	var countries []Country
	var num int64
	var err error
	o := orm.NewOrm()
	qs := o.QueryTable(table)
	cond := orm.NewCondition()
	var cond1 *orm.Condition
	icountry_ids := UserCountryIds(admin, user, eligible)

	cond1 = cond.And("id__in", icountry_ids...)
	num, err = qs.SetCond(cond1).Filter("deleted_at__isnull", true).OrderBy("name").All(&countries)

	return num, countries, err
}

func init() {
	orm.RegisterModel(new(Country))
}