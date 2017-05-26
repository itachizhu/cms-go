package services

import (
	"database/sql"
	"github.com/cms-go/models"
)

type AdminUserService struct {
	db *sql.DB
}

func NewAdminUserService(db *sql.DB) *AdminUserService {
	return &AdminUserService {
		db: db,
	}
}

func (a *AdminUserService) Login(account string, password string) (*models.AdminUser, error) {
	adminUser := &models.AdminUser{}
	rows, err := a.db.Query("select * from t_admuser where accout=? and `password`=password(?)", account, password)
	if err != nil {
		return adminUser, err
	}
	rows.Next()
	err = rows.Scan(&adminUser.Id, &adminUser.Account, &adminUser.Mail, &adminUser.Name, &adminUser.Phone, &adminUser.Department, &adminUser.Password, &adminUser.CreateTime, &adminUser.UpdateTime, &adminUser.Deleted)
	return adminUser, err
}