package models

import (
	"time"
)

type AdminUser struct {
	Id          int64      `json:"id"`
	Account     string     `json:"account"`
	Mail        string     `json:"mail"`
	Name        string     `json:"name"`
	Phone       string     `json:"phone"`
	Department  string     `json:"department"`
	Password    string     `json:"password"`
	CreateTime  time.Time  `json:"create-time"`
	UpdateTime  time.Time  `json:"update-time"`
	Deleted     int8       `json:"deleted"`
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