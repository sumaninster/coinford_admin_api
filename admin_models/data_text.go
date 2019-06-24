package admin_models

import (
	"github.com/astaxie/beego/orm"
)

func (df *DataText) TableName() string {
    return "data_text"
}

func (df *DataText) Insert() error {
	if _, err := orm.NewOrm().Insert(df); err != nil {
		return err
	}
	return nil
}

func (df *DataText) Read(fields ...string) error {
	if err := orm.NewOrm().Read(df, fields...); err != nil {
		return err
	}
	return nil
}

func (df *DataText) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(df, field, fields...)
}

func (df *DataText) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(df, fields...); err != nil {
		return err
	}
	return nil
}

func (df *DataText) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(df, fields ...); err != nil {
		return err
	}
	return nil
}

func DataTextList(dataId int64, fieldId int64) (int64, DataText, error) {
	var table DataText
	var datas []DataText
	num, err := orm.NewOrm().QueryTable(table).Filter("data_id", dataId).Filter("field_id", fieldId).Filter("deleted_at__isnull", true).OrderBy("-created_at").All(&datas)
	return num, datas[0], err
}

func init() {
	orm.RegisterModel(new(DataText))
}
