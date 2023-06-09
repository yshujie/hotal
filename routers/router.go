package routers

import (
	"github.com/astaxie/beego"
	"github.com/hotel/app/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "*:Index")
	beego.Router("/register", &controllers.MainController{}, "*:Register")
	beego.Router("/login", &controllers.MainController{}, "*:Login")
	beego.Router("/logout", &controllers.MainController{}, "*:Logout")

	beego.AutoRouter(&controllers.HotelController{})
	beego.AutoRouter(&controllers.RoomController{})
	beego.AutoRouter(&controllers.UserController{})
	beego.AutoRouter(&controllers.BookOrderController{})
}
