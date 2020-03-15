package repository

import (
	"github.com/labstack/echo/v4"
	"hmdl-user-service/models/data_user"
)

type DmPhanQuyenRepo interface {
	GetAll(ctx echo.Context) ([]data_user.DM_PhanQuyen, error)
	GetById(ctx echo.Context, id int) (*data_user.DM_PhanQuyen, error)
	Delete(ctx echo.Context, id int) error
	Insert(ctx echo.Context, item data_user.DM_PhanQuyen) (*data_user.DM_PhanQuyen, error)
	Update(ctx echo.Context, item data_user.DM_PhanQuyen) error
}
