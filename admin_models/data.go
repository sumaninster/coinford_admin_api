package admin_models

import (
	"github.com/astaxie/beego/orm"
)

func (d *Data) TableName() string {
    return "data"
}

func (d *Data) Insert() error {
	if _, err := orm.NewOrm().Insert(d); err != nil {
		return err
	}
	return nil
}

func (d *Data) Read(fields ...string) error {
	if err := orm.NewOrm().Read(d, fields...); err != nil {
		return err
	}
	return nil
}

func (d *Data) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(d, field, fields...)
}

func (d *Data) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(d, fields...); err != nil {
		return err
	}
	return nil
}

func (d *Data) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(d, fields ...); err != nil {
		return err
	}
	return nil
}

func DataList(userId int64, countryId int64, fieldType string) (int64, []Data, error) {
	var table Data
	var datas []Data
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", userId).Filter("country_id", countryId).Filter("field_type", fieldType).Filter("deleted_at__isnull", true).OrderBy("-created_at").All(&datas, "id", "country_id", "field_type", "nickname", "primary", "active", "created_at", "updated_at")
	return num, datas, err
}

func DataPrimaryFalse(userId int64, countryId int64, fieldType string) (int64, []Data, error) {
	var table Data
	var datas []Data
	num, err := orm.NewOrm().QueryTable(table).Filter("user_id", userId).Filter("country_id", countryId).Filter("field_type", fieldType).Filter("deleted_at__isnull", true).OrderBy("-created_at").All(&datas)
	if err == nil && num > 0 {
		for _, v := range datas {
			v.Primary = false
			v.Update()
		}
	}
	return num, datas, err
}

func init() {
	orm.RegisterModel(new(Data))
}
