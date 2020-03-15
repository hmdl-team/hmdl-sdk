package db

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
)

type MsSql struct {
	Db       *gorm.DB
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
}

func (u *MsSql) SqlServeConnect() {
	conStr := fmt.Sprintf(`sqlserver://%v:%v@%s:%v?database=%v`, u.UserName, u.Password, u.Host, u.Port, u.DbName)
	db, err := gorm.Open("mssql", conStr)

	if err != nil {
		fmt.Println(err)
		raven.CaptureErrorAndWait(err, nil)

	}

	u.Db = db
	u.Db.LogMode(true)
}

func (u *MsSql) Close() {
	defer u.Db.Close()
}
