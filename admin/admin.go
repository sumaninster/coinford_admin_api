package admin

import (
	"coinford_admin_api/admin_models"
	"coinford_admin_api/admin_configs"
	"github.com/astaxie/beego/validation"
	"time"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type AdminController struct {
	beego.Controller
}

// @Title RegisterAdmin
// @Description Register New Admin
// @Param	body		body 	AdminAdd		true		"New Admin Registration Data"
// @Success 200 {int} response
// @Failure 403 body is empty
// @router /register [post]
func (a *AdminController) Add() {
	var rqd AdminAdd
	json.Unmarshal(a.Ctx.Input.RequestBody, &rqd)
	fmt.Println(rqd)
	jres, admin, adminGroup, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		if adminGroup.Key == "SUPER_ADMIN" || adminGroup.Key == "ADMIN" {
			adminNew := admin_models.Admin{Name: rqd.Name, Adminname: rqd.Adminname, AdminGroupId: rqd.AdminGroupId, EditNameTimes: 0, CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *admin_configs.NullTime}
			err := a.validateAdmin(&adminNew)
			if err == nil {
				if a.isUniqueAdminname(adminNew.Adminname) {
					err = adminNew.Insert()
					if err == nil {
						password := admin_models.AdminPassword{AdminId: adminNew.Id, Password: admin_configs.GetSha512(rqd.Password), CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *admin_configs.NullTime}
						err = password.Insert()
						if err == nil {
							email := admin_models.AdminEmail{AdminId: adminNew.Id, Email: rqd.Email, Primary: "YES", CreatedAt: time.Now(), UpdatedAt: time.Now(), DeletedAt: *admin_configs.NullTime}
							err := email.Insert()
							if err == nil {
								a.Data["json"] = a.listAdmins(admin, "Admin registered successfully.")
							} else {
								debugMessage("RegisterAdmin Error (Email): ", err)
								a.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Unable to add email."}
							}
						} else {
							debugMessage("RegisterAdmin Error (Password): ", err)
							a.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Unable to add password."}
						}
					} else {
						debugMessage("RegisterAdmin Error 2: ", err)
						a.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Unable to read and write data. Please try later."}
					}
				} else {
					debugMessage("RegisterAdmin Error 3: ", err)
					a.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Adminname exists"}
				}
			} else {
				debugMessage("RegisterAdmin Error 4: ", err)
				a.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Invalid data"}	
			}
		} else {
			a.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Admin not authorized."}
		}
	} else {
		a.Data["json"] = jres
	}
	a.ServeJSON()
}

// @Title Delete
// @Description List the admin groups
// @Param	body		body 	AdminToken		true		"Token for Authentication"
// @Success 200 {string} list success!
// @Failure 403 uid is empty
// @router /admingroups [post]
func (a *AdminController) AdminGroups() {
	var rqd AdminToken
	json.Unmarshal(a.Ctx.Input.RequestBody, &rqd)
	jres, _, _, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		//if adminGroup.Key == "SUPER_ADMIN" || adminGroup.Key == "ADMIN"  || adminGroup.Key == "CUSTOMER_CARE"{
			_, adminGroups, err := admin_models.AdminGroups()
			if err == nil {
				a.Data["json"] = AdminGroupList{AdminGroups: adminGroups, ResponseCode: 200, ResponseDescription: "List of admin groups"}
			} else {
				a.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Read and write failed."}
			}
		//}
	} else {
		a.Data["json"] = jres
	}
	a.ServeJSON()
}

// @Title Delete
// @Description delete the admin
// @Param	body		body 	AdminId		true		"Token for Authentication"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /delete [delete]
func (a *AdminController) Delete() {
	var rqd AdminId
	json.Unmarshal(a.Ctx.Input.RequestBody, &rqd)
	jres, admin, adminGroup, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		if adminGroup.Key == "SUPER_ADMIN" || adminGroup.Key == "ADMIN" {
			adminDelete := admin_models.Admin{Id: rqd.Id}
			adminDelete.UpdatedAt = time.Now()
			adminDelete.DeletedAt = time.Now()
			err := adminDelete.Update()
			if err == nil {
				a.Data["json"] = a.listAdmins(admin, "Admin deleted successfully")
			} else {
				a.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Delete failed"}
			}
		}
	} else {
		a.Data["json"] = jres
	}
	a.ServeJSON()
}

// @Title ChangeAdminname
// @Description change Adminname for the admin
// @Param	body		body 	AdminChangeAdminname		true		"Change Adminname"
// @Success 200 {string} password change success!
// @Failure 403 password change failed
// @router /changeadminname [post]
func (a *AdminController) ChangeAdminname() {
	var rqd AdminChangeAdminname
	json.Unmarshal(a.Ctx.Input.RequestBody, &rqd)
	jres, admin, _, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		var isvalid bool
		password := admin_models.AdminPassword{AdminId: admin.Id, Password: admin_configs.GetSha512(rqd.CurrentPassword)}
		isvalid, _ = password.IsValid(admin.Id)
		if isvalid {
			admin.Adminname = rqd.NewAdminname
			if a.isUniqueAdminname(admin.Adminname) {
				a.updateAdmin(admin, "Adminname changed successfully")
			} else {
				a.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Adminname exists"}				
			}
		} else {
			a.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid admin"}
		}
	} else {
		a.Data["json"] = jres
	}
	a.ServeJSON()
}

// @Title ChangeName
// @Description change name for the admin (allowed only once). Please make sure this matches your bank account. You will not be able to change the name a second time.
// @Param	body		body 	AdminChangeName		true		"Change Name"
// @Success 200 {string} password change success!
// @Failure 403 password change failed
// @router /changename [post]
func (a *AdminController) ChangeName() {
	var rqd AdminChangeName
	json.Unmarshal(a.Ctx.Input.RequestBody, &rqd)
	jres, admin, _, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		var isvalid bool
		password := admin_models.AdminPassword{AdminId: admin.Id, Password: admin_configs.GetSha512(rqd.CurrentPassword)}
		isvalid, _ = password.IsValid(admin.Id)
		if isvalid {
			if admin.EditNameTimes < admin_configs.EditNameMaximumTimes {
				admin.Name = rqd.NewName
				admin.EditNameTimes += 1
				a.updateAdmin(admin, "Name Changed Successfully")
			} else {
				a.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Edit name not allowed. You have exceeded the number of attemts"}
			}
		} else {
			a.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid admin"}
		}
	} else {
		a.Data["json"] = jres
	}
	a.ServeJSON()
}

// @Title IssueToken
// @Description Issue a new token
// @Success 200 {admin} admin_models.Admin
// @Failure 403 :uid is empty
// @router /token [get]
func (a *AdminController) IssueToken() {
	expirationTime := time.Now().Add(time.Hour * time.Duration(admin_configs.PreLoginTokenTime)).Unix()
	tokenString, _ := tokenAdmin(expirationTime)
    a.Data["json"] = map[string]string{"Token": tokenString}
	a.ServeJSON()
}

// @Title Login
// @Description Logs admin into the system
// @Param	body		body 	AdminLogin		true		"Login Details"
// @Success 200 {string} login success
// @Failure 403 admin not exist
// @router /login [post]
func (a *AdminController) Login() {
	var rqd AdminLogin
	json.Unmarshal(a.Ctx.Input.RequestBody, &rqd)
	err := parseAdminToken(rqd.Token)
	if err == nil {
		admin := admin_models.Admin{Adminname: rqd.Adminname}
		err := admin.Read("adminname")
		if err == nil && admin.DeletedAt == *admin_configs.NullTime {
			var isvalid bool
			password := admin_models.AdminPassword{AdminId: admin.Id, Password: admin_configs.GetSha512(rqd.Password)}
			isvalid, password = password.IsValid(admin.Id)
			if isvalid {
				var tokenString string
				var jres CommonResponse
				jres, tokenString, err = saveAdminToken(&admin)
				if err == nil {
					adminGroup := admin_models.GetAdminGroup(&admin)
					a.Data["json"] = LoginResponse{Token: tokenString, Name: admin.Name, AdminGroup: adminGroup, ResponseCode: 200, ResponseDescription: "Login successful"}
				} else {
					debugMessage("Login Error 1: ", err)
					a.Data["json"] = jres
				}
			} else {
				debugMessage("Login Error 3: ", err)
				a.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid password"}
			}	
		} else {
			debugMessage("Login Error 4: ", err)
			a.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Admin does not exists"}
		}
	} else {
		debugMessage("Login Error 5: ", err)
		a.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid token"}
	}
	a.ServeJSON()
}

// @Title Authenticate
// @Description Authenticates the admin into the system
// @Param	body		body 	AdminToken		true		"Token for Authentication"
// @Success 200 {string} login success
// @Failure 403 admin not exist
// @router /auth [post]
func (a *AdminController) Authenticate() {
	var rqd AdminToken
	json.Unmarshal(a.Ctx.Input.RequestBody, &rqd)
	jres, admin, _, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		jres, tokenString, err := saveAdminToken(admin)
		if err == nil {
			a.Data["json"] = LoginResponse{Token: tokenString, Name: admin.Name, ResponseCode: 200, ResponseDescription: "Login successful"}
		} else {
			debugMessage("Login Error 1: ", err)
			a.Data["json"] = jres
		}
	} else {
		a.Data["json"] = jres
	}
	a.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in admin session
// @Param	body		body 	AdminToken		true		"Token for Authentication"
// @Success 200 {string} logout success
// @router /logout [post]
func (a *AdminController) Logout() {
	var rqd AdminToken
	json.Unmarshal(a.Ctx.Input.RequestBody, &rqd)
	jres, _, _, isLogin, _ := apiAuthenticateAdmin(rqd.Token)
	if isLogin {
		authToken := admin_models.AdminToken{Token: rqd.Token}
		authToken.Read()
		authToken.ExpirationTime = time.Now()
		authToken.UpdatedAt = time.Now()
		authToken.DeletedAt = time.Now()
		err := authToken.Update()
		if err == nil {
			a.Data["json"] = CommonResponse{ResponseCode: 200, ResponseDescription: "Logout successful"}
		} else {
			a.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Unable to read or write data"}
		}
	} else {
		a.Data["json"] = jres
	}
	a.ServeJSON()
}

// @Title IsUniqueAdminname
// @Description Logs out current logged in admin session
// @Param	body		body 	AdminAdminname		true		"Adminname for uniqueness"
// @Success 200 {string} unique adminname
// @router /isuniqueadminname [post]
func (a *AdminController) IsUniqueAdminname() {
	var rqd AdminAdminname
	json.Unmarshal(a.Ctx.Input.RequestBody, &rqd)
	err := parseAdminToken(rqd.Token)
	if err == nil {
		if a.isUniqueAdminname(rqd.Adminname) {
			a.Data["json"] = CommonResponse{ResponseCode: 200, ResponseDescription: "Adminname is unique"}
		} else {
			a.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Adminname exists"}
		}
	} else {
		a.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Invalid token"}
	}
	a.ServeJSON()
}

func (a *AdminController) validateAdmin(admin *admin_models.Admin) error {
	valid := validation.Validation{}
	isvalid, err := valid.Valid(admin)
	if isvalid {
		return nil
	}
	return err
}

func (a *AdminController) isUniqueAdminname(adminname string) bool {
	admin := admin_models.Admin{Adminname: adminname}
	err := admin.Read("adminname")
	fmt.Println(err)
	if err == orm.ErrNoRows {
		return true
	}
	return false
}

func (a *AdminController) updateAdmin(admin *admin_models.Admin, message string) {
	err := a.validateAdmin(admin)
	if err == nil {
		err = admin.Update()
		if err == nil {
			a.Data["json"] = CommonResponse{ResponseCode: 200, ResponseDescription: message}
		} else {
			a.Data["json"] = CommonResponse{ResponseCode: 404, ResponseDescription: "Read and write error"}
		}	
	} else {
		a.Data["json"] = CommonResponse{ResponseCode: 403, ResponseDescription: "Invalid data"}
	}
}

func (a *AdminController) listAdmins(admin *admin_models.Admin, message string) AdminList {
	_, admins, _ := admin_models.Admins(admin)
	return AdminList{Admins: admins, ResponseCode: 200, ResponseDescription: message}
}