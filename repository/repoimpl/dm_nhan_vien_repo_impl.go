package repoimpl

import (
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
	"strings"
)

//Nhân viên khởi tạo
type NhanVienRepoImpl struct {
	DbSql *gorm.DB
	DbPos *gorm.DB
}

//NewTaiKhoanRepo : khởi tạo
func NewNhanVienRepo(DbPos *gorm.DB) repository.NhanVienRepository {
	return &NhanVienRepoImpl{
		DbPos: DbPos,
	}
}

func (u *NhanVienRepoImpl) GetDanhSachNhanVienByChucDanhId(chucDanhId int) []data_user.NhanVien {
	data := make([]data_user.NhanVien, 0)

	err := u.DbPos.Order("DM_NhanVienId desc").Where("chuc_danh_id = ?", chucDanhId).Find(&data).Error
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return nil
	}

	return data

}
func (u *NhanVienRepoImpl) GetDanhSachBacSi() []data_user.NhanVien {
	data := make([]data_user.NhanVien, 0)

	err := u.DbPos.Raw("exec sp_GetDanhSachBacSi").Scan(&data).Error
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return nil
	}

	return data
}
func (u *NhanVienRepoImpl) GetNhanVienByNhanVienId(nhanVienId int) (*data_user.NhanVien, error) {
	data := &data_user.NhanVien{}
	err := u.DbPos.Where("DM_NhanVienId = ?", nhanVienId).Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}
	return data, nil

}
func (u *NhanVienRepoImpl) GetNhanVienById(nhanVienId int) (*data_user.NhanVien, error) {
	data := &data_user.NhanVien{}
	err := u.DbPos.Where("DM_NhanVienId = ?", nhanVienId).Preload("PhongKham").Preload("ChucDanhNhanVien").Preload("ChucVuNhanVien").First(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, nil

}
func (u *NhanVienRepoImpl) GetNhanVienByUserName(userName string) *data_user.NhanVien {
	taikhoan := &data_user.DM_TaiKhoan{}
	u.DbPos.Where(" LOWER(TenTaiKhoan) like ?", strings.ToLower(userName)).First(&taikhoan)
	if taikhoan != nil {
		nhanvien := &data_user.NhanVien{}
		err := u.DbPos.Where("DM_NhanVienId = ?", taikhoan.DM_NhanVienId).Find(&nhanvien).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return nil
		}
		return nhanvien

	}

	return nil
}

func (u *NhanVienRepoImpl) Insert(NhanVien *data_user.NhanVien) error {
	err := u.DbPos.Create(&NhanVien).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}
	return nil
}
func (u *NhanVienRepoImpl) Update(NhanVien *data_user.NhanVien) error {
	err := u.DbPos.Save(NhanVien).Error
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}
	return nil
}
func (u *NhanVienRepoImpl) Delete(id int) error {
	data := &data_user.NhanVien{}
	err := u.DbPos.Where("DM_NhanVienId = ?", id).Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	} else if data != nil {
		err := u.DbPos.Delete(&data).Error
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			return err
		}

	}

	return nil
}
func (u *NhanVienRepoImpl) GetAll() ([]data_user.NhanVien, error) {
	var data []data_user.NhanVien
	err := u.DbPos.Order("DM_NhanVienId desc").Find(&data).Error
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}
	if data != nil {
		return data, nil
	}
	return nil, nil
}
