package admin_models

import (
	"github.com/astaxie/beego/orm"
	"coinford_admin_api/admin_configs"
)

func (f *Field) TableName() string {
    return "field"
}

func (f *Field) Insert() error {
	if _, err := orm.NewOrm().Insert(f); err != nil {
		return err
	}
	return nil
}

func (f *Field) Read(fields ...string) error {
	if err := orm.NewOrm().Read(f, fields...); err != nil {
		return err
	}
	return nil
}

func (f *Field) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(f, field, fields...)
}

func (f *Field) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(f, fields...); err != nil {
		return err
	}
	return nil
}

func (f *Field) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(f, fields ...); err != nil {
		return err
	}
	return nil
}

func FieldIds(countryId int64, fieldType string) []interface{} {
	num, fields, err := Fields(countryId, fieldType)
	field_ids := []int64{}
	if num > 0 && err == nil {
		for _, field := range fields {
		    field_ids = append(field_ids, field.Id)
		}
	}
	ifield_ids := admin_configs.Int64ToInterface(field_ids)
	return ifield_ids
}

func Fields(countryId int64, fieldType string) (int64, []Field, error) {
	var table Field
	var fields []Field
	o := orm.NewOrm()
	qs := o.QueryTable(table)
	cond := orm.NewCondition()
	var qcond *orm.Condition
	if fieldType != "ALL" && fieldType != "" {
		qcond = cond.And("field_type", fieldType)
	}
	num, err := qs.SetCond(qcond).Filter("country_id", countryId).Filter("deleted_at__isnull", true).OrderBy("order").All(&fields)
	return num, fields, err
}

func init() {
	orm.RegisterModel(new(Field))
}
