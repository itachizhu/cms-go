package controllers

import (
	"database/sql"
	"gopkg.in/kataras/iris.v6"
	"fmt"
	"github.com/cms-go/services"
)

type AdminUserController struct {
	db *sql.DB
}

func NewAdminUserController(db *sql.DB) (*AdminUserController)  {
	return &AdminUserController {
		db: db,
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