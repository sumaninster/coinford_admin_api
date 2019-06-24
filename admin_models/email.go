package admin_models

import (
	"github.com/astaxie/beego/orm"
)

func (e *AdminEmail) TableName() string {
    return "admin_email"
}

func (e *AdminEmail) Insert() error {
	if _, err := orm.NewOrm().Insert(e); err != nil {
		return err
	}
	return nil
}

func (e *AdminEmail) Read(fields ...string) error {
	if err := orm.NewOrm().Read(e, fields...); err != nil {
		return err
	}
	return nil
}

func (e *AdminEmail) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(e, field, fields...)
}

func (e *AdminEmail) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(e, fields...); err != nil {
		return err
	}
	return nil
}

func (e *AdminEmail) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(e, fields ...); err != nil {
		return err
	}
	return nil
}

func AdminEmails(admin *Admin) (int64, []AdminEmail, error) {
	var table AdminEmail
	var emails []AdminEmail
	num, err := orm.NewOrm().QueryTable(table).Filter("admin_id", admin.Id).Filter("deleted_at__isnull", true).OrderBy("-updated_at").All(&emails)
	return num, emails, err
}

func init() {
    orm.RegisterModel(new(AdminEmail))
}