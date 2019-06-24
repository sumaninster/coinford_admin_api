package admin_models

import (
    "github.com/astaxie/beego"
    "github.com/astaxie/beego/orm"
    "time"
    "fmt"
    _ "github.com/lib/pq"
)

var (
    DB_PREFIX string
    DB_USER string
    DB_PASSWORD string
    DB_HOST string
    DB_NAME string
    DB_PORT string
    Runmode string
)

type Admin struct {
    Id              int64
    Name            string      `valid:"Required;MaxSize(250)"`
    Adminname       string      `valid:"Required;MaxSize(250)"`
    AdminGroupId    int64
    EditNameTimes   int
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type AdminAccess struct {

}

type AdminEmail struct {
    Id              int64
    AdminId         int64
    Email           string      `valid:"Required;Email;MaxSize(250)"`
    Primary         string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type AdminGroup struct {
    Id              int64
    Key             string      `orm:"unique"`
    Name            string
    Description     string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type AdminPassword struct {
    Id              int64
    AdminId         int64 
    Password        string      `valid:"Required;MinSize(8)"`
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type AdminToken struct {
    Id              int64
    AdminId         int64 
    Token           string      `orm:"unique"`
    ExpirationTime  time.Time
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type AdminCountry struct {
    Id              int64
    AdminId         int64
    CountryId       int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type AdminSetting struct {
    Id              int64
    AdminId         int64 
    Key             string
    Value           string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Country struct {
    Id              int64
    Name            string 
    IsoCode         string
    DialCode        string
    Code            string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Currency struct {
    Id              int64
    Code            string 
    Name            string
    Description     string
    Type            string
    CountryId       int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Field struct {
    Id              int64
    CountryId       int64 
    Name            string
    Description     string
    FieldType       string
    DataType        string
    Order           int64
    Key             string
    HasCategory     bool
    HasInputText    bool
    HasFile         bool
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type FieldCategory struct {
    Id              int64
    FieldId         int64 
    Name            string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type Data struct {
    Id              int64
    UserId          int64 
    CountryId       int64
    FieldType       string
    Nickname        string
    Primary         bool
    Active          bool
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type DataText struct {
    Id              int64
    DataId          int64 
    FieldId         int64
    Text            string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type DataCategory struct {
    Id              int64
    DataId          int64 
    FieldId         int64
    FieldCategoryId int64
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type DataFile struct {
    Id              int64
    DataId          int64 
    FieldId         int64
    Name            string
    Extension       string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type User struct {
    Id              int64
    Name            string      `valid:"Required;MaxSize(250)"`
    Username        string      `valid:"Required;MaxSize(250)"`
    EditNameTimes   int
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

type UserCountry struct {
    Id              int64
    UserId          int64
    CountryId       int64
    Eligible        string
    CreatedAt       time.Time   `orm:"auto_now_add;type(datetime)"`
    UpdatedAt       time.Time   `orm:"auto_now;type(datetime)"`
    DeletedAt       time.Time
}

func init() {
    DB_PREFIX = beego.AppConfig.String("DB_PREFIX")
    DB_USER = beego.AppConfig.String("DB_USER")
    DB_PASSWORD = beego.AppConfig.String("DB_PASSWORD")
    DB_HOST = beego.AppConfig.String("DB_HOST")
    DB_NAME = beego.AppConfig.String("DB_NAME")
    DB_PORT = beego.AppConfig.String("DB_PORT")
    Runmode = beego.AppConfig.String("runmode")

    switch Runmode {
    case "dev":
        orm.RegisterDriver("postgres", orm.DRPostgres)
        orm.RegisterDataBase("default", "postgres", fmt.Sprintf("postgres://%s:%s@%s/%s?port=%i&sslmode=disable",
        DB_USER, DB_PASSWORD, DB_HOST, DB_NAME, DB_PORT), 30)
        //orm.Debug = true
    }
}