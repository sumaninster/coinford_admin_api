package admin_models

import (
	"github.com/astaxie/beego/orm"
)

func (f *FieldCategory) TableName() string {
    return "field_category"
}

func (f *FieldCategory) Insert() error {
	if _, err := orm.NewOrm().Insert(f); err != nil {
		return err
	}
	return nil
}

func (f *FieldCategory) Read(fields ...string) error {
	if err := orm.NewOrm().Read(f, fields...); err != nil {
		return err
	}
	return nil
}

func (f *FieldCategory) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(f, field, fields...)
}

func (f *FieldCategory) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(f, fields...); err != nil {
		return err
	}
	return nil
}

func (f *FieldCategory) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(f, fields ...); err != nil {
		return err
	}
	return nil
}

func FieldCategories(fieldId int64) (int64, []FieldCategory, error) {
	var table FieldCategory
	var fieldcategories []FieldCategory
	num, err := orm.NewOrm().QueryTable(table).Filter("field_id", fieldId).Filter("deleted_at__isnull", true).OrderBy("name").All(&fieldcategories)
	return num, fieldcategories, err
}

func init() {
	orm.RegisterModel(new(FieldCategory))
}
