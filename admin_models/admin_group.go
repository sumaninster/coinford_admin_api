package admin_models

import (
	"github.com/astaxie/beego/orm"
)

func (ag *AdminGroup) TableName() string {
    return "admin_group"
}

func (ag *AdminGroup) Insert() error {
	if _, err := orm.NewOrm().Insert(ag); err != nil {
		return err
	}
	return nil
}

func (ag *AdminGroup) Read(fields ...string) error {
	if err := orm.NewOrm().Read(ag, fields...); err != nil {
		return err
	}
	return nil
}

func (ag *AdminGroup) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(ag, field, fields...)
}

func (ag *AdminGroup) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(ag, fields...); err != nil {
		return err
	}
	return nil
}

func (ag *AdminGroup) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(ag, fields ...); err != nil {
		return err
	}
	return nil
}

func AdminGroups() (int64, []AdminGroup, error) {
	var table AdminGroup
	var adminGroups []AdminGroup
	var num int64
	var err error
	num, err = orm.NewOrm().QueryTable(table).Filter("deleted_at__isnull", true).OrderBy("name").All(&adminGroups, "id", "key", "name", "description")
	return num, adminGroups, err
}

func GetAdminGroup(admin *Admin) AdminGroup {
	adminGroup := AdminGroup{Id: admin.AdminGroupId}
	adminGroup.Read("id")
    return adminGroup
}

func init() {
    orm.RegisterModel(new(AdminGroup))
}