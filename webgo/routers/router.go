package routers

import (
	"github.com/astaxie/beego"
	"github.com/haswelliris/cloudgo-static/webgo/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/user/:name", &controllers.UserController{})
}
