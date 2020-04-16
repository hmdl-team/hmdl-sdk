package repository

import (
	"github.com/labstack/echo/v4"
	"hmdl-user-service/models/data_user"
)

type MenuRepository interface {
	GetAll(ctx echo.Context) ([]data_user.DM_MenuWeb, error)
	GetById(ctx echo.Context, id int) (*data_user.DM_MenuWeb, error)
	Delete(ctx echo.Context, id int) error
	Insert(ctx echo.Context, item data_user.DM_MenuWeb) (*data_user.DM_MenuWeb, error)
	Update(ctx echo.Context, item data_user.DM_MenuWeb) error
	GetMenuByPhanQuyenId(phanQuyenId int) ([]data_user.DM_MenuWeb, error)
	GetMenuByPhanQuyenIdAndDuAnId(phanQuyenId, duAnId int) ([]*data_user.DM_MenuWeb, error)
}
