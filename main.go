package main

import (
	"fmt"
	"github.com/congnguyendl/hmdl-sdk/db"
	"github.com/congnguyendl/hmdl-sdk/sdk"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
	"hmdl-user-service/migration"
	"log"

	_ "hmdl-user-service/docs"
	"hmdl-user-service/router"
	"os"
)

func init() {
	_ = raven.SetDSN("https://bb5a77f6d3e04677a8f01cb2a62e1811:f808d36381bb42709cc650b620cca8d9@sentry.io/1480050")

	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initDB() *gorm.DB {
	msSql := &db.MsSql{
		Host:     os.Getenv("SQL_DATA_USER_HOST"),
		Password: os.Getenv("SQL_DATA_USER_PASSWORD"),
		UserName: os.Getenv("SQL_DATA_USER_USER"),
		DbName:   os.Getenv("SQL_DATA_USER_DBNAME"),
		Port:     os.Getenv("SQL_DATA_USER_PORT")}

	db, err := msSql.SqlServeConnect()

	if err != nil {
		panic(err)
	}

	if err := migration.NewDatabaseRepo(db).Migrate(); err != nil {
		raven.CaptureErrorAndWait(err, nil)
		fmt.Print(err)
	}

	return db
}

func initKong() int {
	//Đăng ký service với kong gate
	port := 7001

	kong := sdk.KongServer{
		ServerKong:  os.Getenv("KONG_ADDRESS"),
		NameService: "HMDL-USER-SERVICE",
		PathService: "/user-service",
		UrlService:  fmt.Sprintf("http://%s:%s", sdk.GetLocalIP(), "7001"),
		IpService:   sdk.GetLocalIP().String() + ":7001",
	}
	if err := kong.RegisterKong(); err != nil {
		fmt.Println(err)
	}
	return port
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

	msSql := initDB()
	port := initKong()

	err := sdk.ConnectNat()
	if err != nil {
		panic(err)
	}

	consulAddress := os.Getenv("CONSUL_ADDRESS")

	sdk.RegisterServiceWithConsul("hmdl-user-service", port, consulAddress)

	e := echo.New()
	e.GET("/docs/*", echoSwagger.WrapHandler)
	e.HideBanner = true

	api := router.API{
		Echo: e,
		Db:   msSql,
	}

	api.NewRouter()

	log.Fatal(e.Start(":7001"))
}
