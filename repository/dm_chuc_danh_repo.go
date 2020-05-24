package repository

import (

	"github.com/labstack/echo/v4"
	"hmdl-user-service/models/data_user"
)

type DmChucDanhRepo interface {
	GetAll(ctx echo.Context) ([]data_user.DmChucDanh, error)
	GetById(ctx echo.Context, id int) (*data_user.DmChucDanh, error)
	Delete(ctx echo.Context, id int) error
	Insert(ctx echo.Context, item data_user.DmChucDanh) (*data_user.DmChucDanh, error)
	Update(ctx echo.Context, item data_user.DmChucDanh) error
}
