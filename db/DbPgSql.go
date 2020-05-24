package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type PgSql struct {
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
}

func (u *PgSql) PostgresConnect() (*gorm.DB, error) {
	conStr := fmt.Sprintf(`host=%v port=%v dbname=%v user=%v password=%v sslmode=disable`, u.Host, u.Port, u.DbName, u.UserName, u.Password)
	db, err := gorm.Open("postgres", conStr)

	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return db, nil
}