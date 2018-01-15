package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

// Hello World!
func (c *MainController) Get() {
	c.Ctx.WriteString("Hello World!\n")
}

//  UserController
type UserController struct {
	beego.Controller
}

// parse username
func (c *UserController) Get() {
	username := c.Ctx.Input.Param(":name")
	c.Ctx.WriteString("Hello " + username + "!\n")
}
