package db

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"os"
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

	//conStr := fmt.Sprintf(`sqlserver://%v:%s@%s:%v?database=%v`, u.UserName, u.Password, u.Host, u.Port, u.DbName)
	connectionString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%v;database=%s", u.Host, u.UserName, u.Password, u.Port, u.DbName)
	db, err := gorm.Open("mssql", connectionString)

	if err != nil {
		fmt.Println(err)
		raven.CaptureErrorAndWait(err, nil)
	}

	evr := os.Getenv("ENVIRONMENT")

	if evr == "DEV" {
		// Show log sql
		db.LogMode(true)
	}

	u.Db = db

}

func (u *MsSql) Close() {
	defer u.Db.Close()
}
