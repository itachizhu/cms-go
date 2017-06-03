package controllers

import (
	"gopkg.in/kataras/iris.v6"
	"fmt"
	"github.com/cms-go/services"
	"github.com/jinzhu/gorm"
)

type AdminUserController struct {
	db *gorm.DB
}

func NewAdminUserController(db ...*gorm.DB) (*AdminUserController)  {
	if len(db) > 0 {
		return &AdminUserController{
			db: db[0],
		}
	} else {
		return &AdminUserController{}
	}
}

func (this *AdminUserController) Login() iris.HandlerFunc {
	return func(ctx *iris.Context) {
		service := services.NewAdminUserService(this.db)
		user, err := service.Login("admin", "111111")
		if err != nil {
			fmt.Printf("%v", err)
			ctx.JSON(503, "login failed")
		}

		ctx.JSON(200, user)
	}
}