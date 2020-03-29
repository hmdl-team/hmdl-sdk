package repository

import (
	"github.com/labstack/echo/v4"
	"hmdl-user-service/models/data_user"
	"hmdl-user-service/models/request"
)

type PhanQuyenMenuRepository interface {
	GetAllPhanQuyenMenu() ([]data_user.DM_PhanQuyenMenu, error)
	GetMenuByPhanQuyenId(phanQuyenId int, duAnId int) ([]data_user.DM_MenuWeb, error)

	GetById(id int) (*data_user.DM_PhanQuyenMenu, error)
	Insert(u data_user.DM_PhanQuyenMenu) error
	Update(u data_user.DM_PhanQuyenMenu) error
	Delete(id int) error
	UpdatePhanQuyen(ctx echo.Context, req request.PhanQuyenMenuReq) error
}
