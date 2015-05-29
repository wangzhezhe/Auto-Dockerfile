package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

// Operations about object
type TestController struct {
	beego.Controller
}

// @Title render main page
// @Description : start the websocket connection
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [get]
func (o *TestController) Get() {

	fmt.Println("test")
}
