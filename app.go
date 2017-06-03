package main

import (
	"gopkg.in/kataras/iris.v6"
	"gopkg.in/kataras/iris.v6/adaptors/httprouter"
	"gopkg.in/kataras/iris.v6/adaptors/view"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/cms-go/routers"
)

func main() {
	db, err := gorm.Open("mysql", "cmsadmin:cmsadmin@tcp(127.0.0.1:3306)/cmsadmin?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		db = nil
	} else {
		defer db.Close()
	}

	app := iris.New(iris.Configuration{Gzip: false, Charset: "UTF-8"}) // defaults to these

	app.Adapt(iris.DevLogger())
	app.Adapt(httprouter.New())
	app.Adapt(view.HTML("./templates", ".html"))

	routers.Router(app, db)

	app.Listen(":8080")
}

