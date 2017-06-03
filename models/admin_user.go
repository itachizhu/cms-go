package models

import (
	"time"
)

type AdminUser struct {
	Id          int64      `json:"id",gorm:"AUTO_INCREMENT;primary_key"`
	Account     string     `json:"account",gorm:"column:accout"`
	Mail        string     `json:"mail"`
	Name        string     `json:"name"`
	Phone       string     `json:"phone"`
	Department  string     `json:"department"`
	Password    string     `json:"password,omitempty"`
	CreateTime  time.Time  `json:"create-time"`
	UpdateTime  time.Time  `json:"update-time"`
	Deleted     int8       `json:"deleted",gorm:"column:isdel"`
}

/*
func NewAdminUser(db *sql.DB) *AdminUser {
	return &AdminUser {
		db: db,
	}
}

func (a *AdminUser) Get(account string, password string) (*AdminUser, error) {
	rows, err := a.db.Query("select * from adm where account=$1 and `password`=password($2)", account, password)
	if err != nil {
		return a, err
	}
	rows.Next()
	err = rows.Scan(&a.Account)
	return a, err
}
*/

func (AdminUser) TableName() string  {
	return "t_admuser"
}