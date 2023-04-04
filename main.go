package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"github.com/hotel/app/service"
	"time"

	_ "github.com/hotel/routers"
)

const VERSION = "1.0"

func main() {
	service.Init()

	_ = beego.AppConfig.Set("version", VERSION)
	if beego.AppConfig.String("runmode") != "prod" {
		beego.SetLevel(beego.LevelDebug)
	} else {
		beego.SetLevel(beego.LevelInformational)
		beego.SetLogger("file", `{"filename":"`+beego.AppConfig.String("log_file")+`"}`)
		beego.BeeLogger.DelLogger("console")
	}

	// 记录启动时间
	beego.AppConfig.Set("up_time", fmt.Sprintf("%d", time.Now().Unix()))

	// todo: what's i18n
	beego.AddFuncMap("i18n", i18n.Tr)

	// todo: what's assets
	beego.SetStaticPath("/assets", "assets")
	beego.Run()
}
