package main

import (
	_ "coinford_admin_api/routers"
	"coinford_admin_api/admin_configs"
	"github.com/astaxie/beego"
)

func init() {
    admin_configs.Init()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
