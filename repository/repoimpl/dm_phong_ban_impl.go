package repoimpl

import (
	"errors"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	. "hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
)

func NewDM_PhongBanRepo(db *gorm.DB) repository.DM_PhongBanRepo {
	return &DM_PhongBanRepoImpl{db: db}
}

type DM_PhongBanRepoImpl struct {
	db *gorm.DB
}

func (u *DM_PhongBanRepoImpl) GetAll(ctx echo.Context) ([]DM_PhongBan, error) {
	var data []DM_PhongBan

	err := u.db.Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *DM_PhongBanRepoImpl) GetById(ctx echo.Context, id int) (*DM_PhongBan, error) {
	var dsPhongKha DM_PhongBan
	err := u.db.Where("DM_PhongBanId = ?", id).Find(&dsPhongKha).Error

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

func (u *DM_PhongBanRepoImpl) Delete(ctx echo.Context, id int) error {
	var data DM_PhongBan

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

func (u *DM_PhongBanRepoImpl) Insert(ctx echo.Context, item DM_PhongBan) (*DM_PhongBan, error) {
	err := u.db.Create(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return &item, nil
}

func (u *DM_PhongBanRepoImpl) Update(ctx echo.Context, item DM_PhongBan) error {
	err := u.db.Save(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}
