package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"App/models"
)

type TareaController struct {
	beego.Controller
}


func (this *TareaController) List() {
	res := struct{ Tasks []*models.Task }{models.DefaultTaskList.All()}
	this.Data["json"] = res
	this.ServeJSON()
}


func (this *TareaController) New() {
	req := struct{ Title string }{}
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &req); err != nil {
		this.Ctx.Output.SetStatus(400)
		this.Ctx.Output.Body([]byte("empty title"))
		return
	}
	t, _ := models.New(req.Title)

	models.DefaultTaskList.Save(t)
}



