package core

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type DbData struct {
	Echo    *echo.Echo
	DbSql01 *gorm.DB
}
