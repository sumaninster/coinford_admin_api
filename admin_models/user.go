package admin_models

import (
	"github.com/astaxie/beego/orm"
)

func (u *User) TableName() string {
    return "user"
}

func (u *User) Insert() error {
	if _, err := orm.NewOrm().Insert(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(u, field, fields...)
}

func (u *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(u, fields...); err != nil {
		return err
	}
	return nil
}

func (u *User) Delete(fields ...string) error {
	if _, err := orm.NewOrm().Delete(u, fields ...); err != nil {
		return err
	}
	return nil
}

func Users(admin *Admin, countryId int64, eligible string) (int64, []User, error) {
	var table User
	var users []User
	var num int64
	var err error
	adminGroup := GetAdminGroup(admin)
	if adminGroup.Key == "SUPER_ADMIN" || adminGroup.Key == "ADMIN" || adminGroup.Key == "CUSTOMER_SUPPORT" {
		o := orm.NewOrm()
		qs := o.QueryTable(table)
		cond := orm.NewCondition()
		var cond1 *orm.Condition

		num, userIds := AdminCountryIdsUser(admin, countryId, eligible)
		if num > 0 {
			cond1 = cond.And("id__in", userIds)
			num, err = qs.SetCond(cond1).Filter("deleted_at__isnull", true).OrderBy("-created_at").All(&users)
		}		
	}
	return num, users, err
}

func init() {
	orm.RegisterModel(new(User))
}
