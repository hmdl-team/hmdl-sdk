package repository

import (
	"github.com/labstack/echo/v4"
	. "hmdl-user-service/models/data_user"
)

type DM_PhongBanRepo interface {
	GetAll(ctx echo.Context) ([]DM_PhongBan, error)
	GetPhongBanComBobox(ctx echo.Context) ([]DM_PhongBan, error)
	GetById(ctx echo.Context, id int) (*DM_PhongBan, error)
	Delete(ctx echo.Context, id int) error
	Insert(ctx echo.Context, item DM_PhongBan) (*DM_PhongBan, error)
	Update(ctx echo.Context, item DM_PhongBan) error
}