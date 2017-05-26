package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/cms-go/controllers"
)

func main() {
	db, err := sql.Open("mysql", "cmsadmin:cmsadmin@/cmsadmin?parseTime=true")

	if err != nil {
		db = nil
	} else {
		defer db.Close()
	}

	app := iris.New(iris.Configuration{Gzip: false, Charset: "UTF-8"}) // defaults to these

	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())
	app.Adapt(view.HTML("./templates", ".html"))

	app.Get("/login", controllers.NewAdminUserController(db).Login())

	app.Listen(":8080")
}

