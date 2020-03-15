package repository

import (
	"hmdl-user-service/models/data_user"
)

type PhanQuyenMenuRepository interface {
	GetAllPhanQuyenMenu() ([]data_user.DM_PhanQuyenMenu, error)
	GetMenuByPhanQuyenId(phanQuyenId int,duAnId int) ([]data_user.DM_MenuWeb, error)

	GetById(id int) (*data_user.DM_PhanQuyenMenu, error)
	Insert(u *data_user.DM_PhanQuyenMenu) error
	Update(u *data_user.DM_PhanQuyenMenu) error
	Delete(id int) error
}
