package repository

import (
	"github.com/labstack/echo/v4"
	. "hmdl-user-service/models/data_user"
)

type DanhSachNhanVienQuanLyRepo interface {
	GetAll(ctx echo.Context) ([]DM_NhanVienQuanLy, error)
	GetById(ctx echo.Context, id int) (*DM_NhanVienQuanLy, error)
	Delete(ctx echo.Context, id int) error
	Insert(ctx echo.Context, item DM_NhanVienQuanLy) (*DM_NhanVienQuanLy, error)
	Update(ctx echo.Context, item DM_NhanVienQuanLy) error
}
