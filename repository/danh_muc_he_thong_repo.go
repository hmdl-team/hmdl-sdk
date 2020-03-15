package repository

import (
	"hmdl-user-service/models/data_user"
)

type DanhMucHeThongRepository interface {
	DongBoChucdanh() error
	DongBoChucVu() error
	Insert(u *data_user.DanhMucHeThong) error
	GetDanhMucHeThongByDanhMucCode(DanhMucCode int, LoaiDanhMuc string) (*data_user.DanhMucHeThong, error)
	GetAllDanhMucHeThongByLoaiDanhMuc(LoaiDanhMuc string) ([]data_user.DanhMucHeThong, error)
	GetAllChucDanh() ([]data_user.DanhMucHeThong, error)
	GetAllChucVu() ([]data_user.DanhMucHeThong, error)
}