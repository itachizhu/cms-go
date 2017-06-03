package services

import (
	"github.com/cms-go/models"
	"github.com/jinzhu/gorm"
)

type AdminUserService struct {
	db *gorm.DB
}

func NewAdminUserService(db *gorm.DB) *AdminUserService {
	return &AdminUserService {
		db: db,
	}
}

func (this *AdminUserService) Login(account string, password string) (*models.AdminUser, error) {
	adminUser := &models.AdminUser{}
	this.db.Where("accout=? and `password`=password(?) and isdel=0", account, password).Find(adminUser)
	adminUser.Password = ""
	/*
	rows, err := a.db.Query("select * from t_admuser where accout=? and `password`=password(?)", account, password)
	if err != nil {
		return adminUser, err
	}
	rows.Next()
	err = rows.Scan(&adminUser.Id, &adminUser.Account, &adminUser.Mail, &adminUser.Name, &adminUser.Phone, &adminUser.Department, &adminUser.Password, &adminUser.CreateTime, &adminUser.UpdateTime, &adminUser.Deleted)
	*/
	return adminUser, this.db.Error
}