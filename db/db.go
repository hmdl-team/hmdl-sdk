package db

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// https://jmoiron.github.io/sqlx/
type Sql struct {
	Db       *gorm.DB
	Host     string
	Port     int
	UserName string
	Password string
	DbName   string
}
type DbMySql struct {
	Db       *gorm.DB
	Host     string
	Port     int
	UserName string
	Password string
	DbName   string
}

func (u *DbMySql) MySqlConnect() {
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

func (u *Sql) PostgresConnect() {
	conStr := fmt.Sprintf(`host=%v port=%v dbname=%v user=%v password=%v sslmode=disable`,
		u.Host, u.Port, u.DbName, u.UserName, u.Password)

	db, err := gorm.Open("postgres", conStr)

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Println(err)
	}

	// Show log sql
	//db.LogMode(true)
	u.Db = db

}

func (u *Sql) SqlServeConnect() {
	conStr := fmt.Sprintf(`sqlserver://%v:%v@%v:%v?database=%v`, u.UserName, u.Password, u.Host, u.Port, u.DbName)

	db, err := gorm.Open("mssql", conStr)

	if err != nil {
		fmt.Println(err)
		raven.CaptureErrorAndWait(err, nil)
	}

	u.Db = db

}

func (u *Sql) Close() {
	_ = u.Db.Close()
}
func (u *DbMySql) Close() {
	_ = u.Db.Close()
}
