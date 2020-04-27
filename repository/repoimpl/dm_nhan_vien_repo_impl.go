package repoimpl

import (
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"hmdl-user-service/db/core"
	"hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
	"strings"
)

func NewNhanVienRepo(db *core.DbData) repository.NhanVienRepository {
	return &NhanVienRepoImpl{
		db: db,
	}
}

type NhanVienRepoImpl struct {
	db *core.DbData
}

func (u *NhanVienRepoImpl) GetNhanVienByPhongBanId(id int) (data []data_user.NhanVien, err error) {

	err = u.db.DbSql01.Raw(`
	EXEC sp_DM_NhanVien_PhongBan @PhongBanID = ?
		`, id).Scan(&data).Error

	if err != nil && gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}
	return data, err
}

func (u *NhanVienRepoImpl) GetNhanVienCombobox() ([]data_user.NhanVien, error) {
	var data []data_user.NhanVien
	err := u.db.DbSql01.Where(&data_user.NhanVien{TinhTrang: true}).Order("DM_NhanVienId desc").Find(&data).Error
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}
	if data != nil {
		return data, nil
	}
	return nil, nil
}

func (u *NhanVienRepoImpl) GetDanhSachNhanVienByChucDanhId(chucDanhId int) []data_user.NhanVien {
	data := make([]data_user.NhanVien, 0)

	err := u.db.DbSql01.Order("DM_NhanVienId desc").Where("chuc_danh_id = ?", chucDanhId).Find(&data).Error
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return nil
	}

	return data

}
func (u *NhanVienRepoImpl) GetDanhSachBacSi() []data_user.NhanVien {
	data := make([]data_user.NhanVien, 0)

	err := u.db.DbSql01.Raw("exec sp_GetDanhSachBacSi").Scan(&data).Error
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return nil
	}

	return data
}
func (u *NhanVienRepoImpl) GetNhanVienByNhanVienId(nhanVienId int) (*data_user.NhanVien, error) {
	data := &data_user.NhanVien{}
	err := u.db.DbSql01.Where("DM_NhanVienId = ?", nhanVienId).Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}
	return data, nil

}
func (u *NhanVienRepoImpl) GetNhanVienById(nhanVienId int) (*data_user.NhanVien, error) {
	data := &data_user.NhanVien{}
	err := u.db.DbSql01.Where("DM_NhanVienId = ?", nhanVienId).
		Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, nil

}
func (u *NhanVienRepoImpl) GetNhanVienByUserName(userName string) *data_user.NhanVien {
	taikhoan := &data_user.DM_TaiKhoan{}
	u.db.DbSql01.Where(" LOWER(TenTaiKhoan) like ?", strings.ToLower(userName)).Find(&taikhoan)
	if taikhoan != nil {
		nhanvien := &data_user.NhanVien{}
		err := u.db.DbSql01.Where("DM_NhanVienId = ?", taikhoan.DM_NhanVienId).Find(&nhanvien).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			return nil
		}
		return nhanvien

	}

	return nil
}

func (u *NhanVienRepoImpl) Insert(NhanVien *data_user.NhanVien) error {
	err := u.db.DbSql01.Create(&NhanVien).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}
	return nil
}
func (u *NhanVienRepoImpl) Update(NhanVien *data_user.NhanVien) error {
	err := u.db.DbSql01.Save(NhanVien).Error
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}
	return nil
}
func (u *NhanVienRepoImpl) Delete(id int) error {
	data := &data_user.NhanVien{}
	err := u.db.DbSql01.Where("DM_NhanVienId = ?", id).Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	} else if data != nil {
		err := u.db.DbSql01.Delete(&data).Error
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			return err
		}

	}

	return nil
}
func (u *NhanVienRepoImpl) GetAll() ([]data_user.NhanVien, error) {
	var data []data_user.NhanVien
	err := u.db.DbSql01.Order("DM_NhanVienId desc").Find(&data).Error
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}
	if data != nil {
		return data, nil
	}
	return nil, nil
}
