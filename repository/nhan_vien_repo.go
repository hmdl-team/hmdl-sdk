package repository

import (
	"hmdl-user-service/models/data_user"
)

type NhanVienRepository interface {
	GetDanhSachNhanVienByChucDanhId(chucDanhId int) []data_user.NhanVien
	GetDanhSachBacSi() []data_user.NhanVien
	GetNhanVienByNhanVienId(nhanVienId int) (*data_user.NhanVien, error)
	GetNhanVienById(id int) (*data_user.NhanVien, error)
	GetNhanVienByUserName(userName string) *data_user.NhanVien

	Insert(u *data_user.NhanVien) error
	Update(u *data_user.NhanVien) error
	Delete(id int) error
	GetAll() ([]data_user.NhanVien, error)
}
