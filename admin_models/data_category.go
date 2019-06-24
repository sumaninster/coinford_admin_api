package admin_models

import (
	"github.com/astaxie/beego/orm"
)

func (dc *DataCategory) TableName() string {
    return "data_category"
}

func (dc *DataCategory) Insert() error {
	if _, err := orm.NewOrm().Insert(dc); err != nil {
		return err
	}
	return nil
}

func (dc *DataCategory) Read(fields ...string) error {
	if err := orm.NewOrm().Read(dc, fields...); err != nil {
		return err
	}
	return nil
}

func (dc *DataCategory) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(dc, field, fields...)
}

func (dc *DataCategory) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(dc, fields...); err != nil {
		return err
	}
	return nil
}

func (dc *DataCategory) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(dc, fields ...); err != nil {
		return err
	}
	return nil
}

func DataCategoryList(dataId int64, fieldId int64) (int64, []DataCategory, error) {
	var table DataCategory
	var datas []DataCategory
	num, err := orm.NewOrm().QueryTable(table).Filter("data_id", dataId).Filter("field_id", fieldId).Filter("deleted_at__isnull", true).OrderBy("-created_at").All(&datas)
	return num, datas, err
}

func init() {
	orm.RegisterModel(new(DataCategory))
}
