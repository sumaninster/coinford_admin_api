package admin_models

import (
	"github.com/astaxie/beego/orm"
)

func (df *DataFile) TableName() string {
    return "data_file"
}

func (df *DataFile) Insert() error {
	if _, err := orm.NewOrm().Insert(df); err != nil {
		return err
	}
	return nil
}

func (df *DataFile) Read(fields ...string) error {
	if err := orm.NewOrm().Read(df, fields...); err != nil {
		return err
	}
	return nil
}

func (df *DataFile) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(df, field, fields...)
}

func (df *DataFile) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(df, fields...); err != nil {
		return err
	}
	return nil
}

func (df *DataFile) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(df, fields ...); err != nil {
		return err
	}
	return nil
}

func DataFileList(dataId int64, fieldId int64) (int64, []DataFile, error) {
	var table DataFile
	var datas []DataFile
	num, err := orm.NewOrm().QueryTable(table).Filter("data_id", dataId).Filter("field_id", fieldId).Filter("deleted_at__isnull", true).OrderBy("-created_at").All(&datas)
	return num, datas, err
}

func init() {
	orm.RegisterModel(new(DataFile))
}
