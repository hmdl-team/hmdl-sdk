package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MySql struct {
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
}

func (u *MySql) MySqlConnect() (*gorm.DB, error) {
	conStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", u.UserName, u.Password, u.Host, u.Port, u.DbName)
	db, err := gorm.Open("mysql", conStr)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return db, nil
}
