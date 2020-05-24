package sdkcm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	Db *gorm.DB
)

func ConnectDb() error {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Asia%%2fSaigon",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	if os.Getenv("ENVIRONMENT") == "development" {
		logrus.Info(connectionString)
	}

	db, err := gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}

	db = db.LogMode(true)
	Db = db
	return nil
}
