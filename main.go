package main

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"hmdl-user-service/db"
	_ "hmdl-user-service/docs"
	"hmdl-user-service/migration"
	"hmdl-user-service/router"
	"log"
	"os"
)

func init() {
	_ = raven.SetDSN("https://bb5a77f6d3e04677a8f01cb2a62e1811:f808d36381bb42709cc650b620cca8d9@sentry.io/1480050")
}

// @title hmdl-user-service api
// @version 1.0
// @description Service api HMDL
// @contact.name IT HMDL
// @contact.url https://hoanmydalat.com
// @contact.email nguyen.nguyen@hoanmy.com
// @securityDefinitions.apikey jwt
// @in header
// @name Authorization
// @host localhost:7001
func main() {
	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	msSql := &db.MsSql{
		Host:     os.Getenv("SQL_DATA_USER_HOST"),
		Password: os.Getenv("SQL_DATA_USER_PASSWORD"),
		UserName: os.Getenv("SQL_DATA_USER_USER"),
		DbName:   os.Getenv("SQL_DATA_USER_DBNAME"),
		Port:     os.Getenv("SQL_DATA_USER_PORT"),
	}

	msSql.SqlServeConnect()
	defer msSql.Close()

	err = migration.NewDatabaseRepo(msSql.Db).Migrate()

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Print(err)
	}

	e := echo.New()

	e.GET("/docs/*", echoSwagger.WrapHandler)
	//swag init

	api := router.API{
		Echo: e,
		Db:   msSql.Db,
	}

	api.NewRouter()
	e.Logger.Fatal(e.Start(":7001"))
}
