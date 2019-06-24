package admin_models

import (
	"github.com/astaxie/beego/orm"
)

func (ap *AdminPassword) TableName() string {
    return "admin_password"
}

func (ap *AdminPassword) Insert() error {
	if _, err := orm.NewOrm().Insert(ap); err != nil {
		return err
	}
	return nil
}

func (ap *AdminPassword) Read(fields ...string) error {
	if err := orm.NewOrm().Read(ap, fields...); err != nil {
		return err
	}
	return nil
}

func (ap *AdminPassword) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(ap, field, fields...)
}

func (ap *AdminPassword) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(ap, fields...); err != nil {
		return err
	}
	return nil
}

func (ap *AdminPassword) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(ap, fields ...); err != nil {
		return err
	}
	return nil
}

func (ap *AdminPassword) IsValid(adminId int64) (bool, AdminPassword){
	var table AdminPassword
	var passwords []AdminPassword
	num, err := orm.NewOrm().QueryTable(table).Filter("admin_id", adminId).Filter("password", ap.Password).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&passwords)
	if err == nil {
		if num == 1 {
			return true, passwords[0]
		}
		return false, AdminPassword{}
	}
	return false, AdminPassword{}
}

func init() {
	orm.RegisterModel(new(AdminPassword))
}
