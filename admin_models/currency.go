package admin_models

import (
	"github.com/astaxie/beego/orm"
)

func (c *Currency) TableName() string {
    return "currency"
}

func (c *Currency) Insert() error {
	if _, err := orm.NewOrm().Insert(c); err != nil {
		return err
	}
	return nil
}

func (c *Currency) Read(fields ...string) error {
	if err := orm.NewOrm().Read(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Currency) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(c, field, fields...)
}

func (c *Currency) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(c, fields...); err != nil {
		return err
	}
	return nil
}

func (c *Currency) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(c, fields ...); err != nil {
		return err
	}
	return nil
}

func Currencies(admin *Admin) (int64, []Currency, error) {
	var table Currency
	var currencies []Currency
	var num int64
	var err error
	adminGroup := GetAdminGroup(admin)
	if adminGroup.Key == "SUPER_ADMIN" || adminGroup.Key == "ADMIN" || adminGroup.Key == "CUSTOMER_SUPPORT" {
		num, err = orm.NewOrm().QueryTable(table).Filter("deleted_at__isnull", true).OrderBy("code").All(&currencies)	
	}
	return num, currencies, err
}

func init() {
    orm.RegisterModel(new(Currency))
}