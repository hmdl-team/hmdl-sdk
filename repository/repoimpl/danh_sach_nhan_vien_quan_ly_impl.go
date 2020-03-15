package repoimpl

import (
	"errors"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
)

func NewDanhSachNhanVienQuanLyRepo(db *gorm.DB) repository.DanhSachNhanVienQuanLyRepo {
	return &DanhSachNhanVienQuanLyRepoImpl{db: db}
}

type DanhSachNhanVienQuanLyRepoImpl struct {
	db *gorm.DB
}

func (u *DanhSachNhanVienQuanLyRepoImpl) GetAll(ctx echo.Context) ([]data_user.DM_NhanVienQuanLy, error) {
	var data []data_user.DM_NhanVienQuanLy

	err := u.db.Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *DanhSachNhanVienQuanLyRepoImpl) GetById(ctx echo.Context, id int) (*data_user.DM_NhanVienQuanLy, error) {
	var dsPhongKha data_user.DM_NhanVienQuanLy
	err := u.db.Where("DanhSachNhanVienQuanLyId = ?", id).Find(&dsPhongKha).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return &dsPhongKha, err
	}

	if err != nil && gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}
	return &dsPhongKha, nil
}

func (u *DanhSachNhanVienQuanLyRepoImpl) Delete(ctx echo.Context, id int) error {
	var data data_user.DM_NhanVienQuanLy

	err := u.db.Find(&data, id).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Không tìm thấy dữ liệu")
	}

	err = u.db.Delete(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}

func (u *DanhSachNhanVienQuanLyRepoImpl) Insert(ctx echo.Context, item data_user.DM_NhanVienQuanLy) (*data_user.DM_NhanVienQuanLy, error) {
	err := u.db.Create(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return &item, nil
}

func (u *DanhSachNhanVienQuanLyRepoImpl) Update(ctx echo.Context, item data_user.DM_NhanVienQuanLy) error {
	err := u.db.Save(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}
