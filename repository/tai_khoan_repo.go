package repository

import (
	"github.com/labstack/echo/v4"
	"hmdl-user-service/models/data_user"
	"hmdl-user-service/models/request"
)

type TaiKhoanRepository interface {
	GetAll(ctx echo.Context) ([]data_user.DM_TaiKhoan, error)
	GetById(ctx echo.Context, id int) (*data_user.DM_TaiKhoan, error)
	Delete(ctx echo.Context, id int) error
	Insert(ctx echo.Context, item data_user.DM_TaiKhoan) (*data_user.DM_TaiKhoan, error)
	Update(ctx echo.Context, item data_user.DM_TaiKhoan) error

	//Login user acount
	Login(ctx echo.Context, loginReq request.ReqSignIn) (*data_user.DM_TaiKhoan, error)
}
