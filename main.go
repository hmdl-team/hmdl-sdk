package main

import (
	"fmt"
	"github.com/getsentry/raven-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"hmdl-user-service/migration"
	"log"

	"hmdl-user-service/db"
	_ "hmdl-user-service/docs"
	"hmdl-user-service/router"
	"os"
	"runtime"
)

func init() {
	_ = raven.SetDSN("https://bb5a77f6d3e04677a8f01cb2a62e1811:f808d36381bb42709cc650b620cca8d9@sentry.io/1480050")

	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
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

	ng := runtime.NumCPU()
	fmt.Println("NumCPU: ", ng)

	msSql := &db.MsSql{
		Host:     os.Getenv("SQL_DATA_USER_HOST"),
		Password: os.Getenv("SQL_DATA_USER_PASSWORD"),
		UserName: os.Getenv("SQL_DATA_USER_USER"),
		DbName:   os.Getenv("SQL_DATA_USER_DBNAME"),
		Port:     os.Getenv("SQL_DATA_USER_PORT"),
	}

	msSql.SqlServeConnect()
	defer msSql.Close()

	if err := migration.NewDatabaseRepo(msSql.Db).Migrate(); err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Print(err)
	}

	e := echo.New()

	//e.Debug = true

	e.GET("/docs/*", echoSwagger.WrapHandler)
	//swag init

	api := router.API{
		Echo: e,
		Db:   msSql.Db,
	}

	api.NewRouter()

	//log.Fatal(e.Start(":7001"))
}
