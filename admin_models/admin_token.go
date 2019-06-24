package admin_models

import (
	"github.com/astaxie/beego/orm"
)

func (at *AdminToken) TableName() string {
    return "admin_token"
}

func (at *AdminToken) Insert() error {
	if _, err := orm.NewOrm().Insert(at); err != nil {
		return err
	}
	return nil
}

func (at *AdminToken) Read(fields ...string) error {
	if err := orm.NewOrm().Read(at, fields...); err != nil {
		return err
	}
	return nil
}

func (at *AdminToken) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(at, field, fields...)
}

func (at *AdminToken) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(at, fields...); err != nil {
		return err
	}
	return nil
}

func (at *AdminToken) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(at, fields ...); err != nil {
		return err
	}
	return nil
}

func init() {
    orm.RegisterModel(new(AdminToken))
}