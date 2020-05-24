package repository

import (
	"github.com/labstack/echo/v4"
	"hmdl-user-service/models/data_user"
)

type DmChucVuRepo interface {
	GetAll(ctx echo.Context) ([]data_user.DmChucVu, error)
	GetById(ctx echo.Context, id int) (*data_user.DmChucVu, error)
	Delete(ctx echo.Context, id int) error
	Insert(ctx echo.Context, item data_user.DmChucVu) (*data_user.DmChucVu, error)
	Update(ctx echo.Context, item data_user.DmChucVu) error
}
