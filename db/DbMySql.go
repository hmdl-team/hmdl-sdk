package db

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type MySql struct {
	Db       *gorm.DB
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
}

func (u *MySql) MySqlConnect() {
	conStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local", u.UserName, u.Password, u.Host, u.Port, u.DbName)
	db, err := gorm.Open("mysql", conStr)


	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Println(err)
	}

	// Show log sql
	//db.LogMode(true)
	u.Db = db
}
func (u *MySql) Close() {
	_ = u.Db.Close()
}
