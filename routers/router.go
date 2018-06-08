package routers

import (
	"minego/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/users/?:id", &controllers.UsersControllers{})
}
