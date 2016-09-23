package main

import (
	"github.com/astaxie/beego"
	"App/controllers"
)

func main() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/tarea/", &controllers.TareaController{}, "get:List;post:New")
	beego.Run()
}
