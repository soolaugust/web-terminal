package main

import (
	"web-terminal/controller"

	"github.com/astaxie/beego"
)

func main() {
	beego.BConfig.WebConfig.AutoRender = false
	beego.Router("/terminal", &controller.TerminalController{})
	beego.Run(":8081")
}
