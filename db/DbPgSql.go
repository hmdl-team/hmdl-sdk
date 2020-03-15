package db

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PgSql struct {
	Db       *gorm.DB
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
}

func (u *PgSql) PostgresConnect() {
	conStr := fmt.Sprintf(`host=%v port=%v dbname=%v user=%v password=%v sslmode=disable`, u.Host, u.Port, u.DbName, u.UserName, u.Password)
	db, err := gorm.Open("postgres", conStr)

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Println(err)
	}

	// Show log sql
	//db.LogMode(true)
	u.Db = db
}

func (u *PgSql) Close() {
	_ = u.Db.Close()
}
