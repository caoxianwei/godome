package main

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Ctx.WriteString("hello world")
	this.Ctx.WriteString("appname" + beego.AppConfig.String("appname")+
		"\nhttpport" + beego.AppConfig.String("httpport")+
		"\nrunmode" + beego.AppConfig.String("runmode"))
}

func main() {
	beego.Router("/", &MainController{})
	beego.Run()
}