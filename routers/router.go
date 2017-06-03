package routers

import (
	"gopkg.in/kataras/iris.v6"
	"github.com/jinzhu/gorm"
	"github.com/cms-go/controllers"
)

func Router(app *iris.Framework, db ...*gorm.DB)  {
	app.Get("/login", controllers.NewAdminUserController(db...).Login())
}