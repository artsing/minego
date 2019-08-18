package routers

import (
	"github.com/astaxie/beego"
	"minego/controllers"
)

func init() {
	beego.Router("/users/?:id", &controllers.UsersControllers{})
}
