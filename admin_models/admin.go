package admin_models

import (
	"github.com/astaxie/beego/orm"
)

func (a *Admin) TableName() string {
    return "admin"
}

func (a *Admin) Insert() error {
	if _, err := orm.NewOrm().Insert(a); err != nil {
		return err
	}
	return nil
}

func (a *Admin) Read(fields ...string) error {
	if err := orm.NewOrm().Read(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *Admin) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(a, field, fields...)
}

func (a *Admin) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(a, fields...); err != nil {
		return err
	}
	return nil
}

func (a *Admin) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(a, fields ...); err != nil {
		return err
	}
	return nil
}

func Admins(admin *Admin) (int64, []Admin, error) {
	var table Admin
	var admins []Admin
	var num int64
	var err error
	adminGroup := GetAdminGroup(admin)
	if adminGroup.Key == "SUPER_ADMIN" && adminGroup.Key == "ADMIN" {
		num, err = orm.NewOrm().QueryTable(table).Filter("deleted_at__isnull", true).OrderBy("name").All(&admins)
	}
	return num, admins, err
}

func init() {
	orm.RegisterModel(new(Admin))
}
