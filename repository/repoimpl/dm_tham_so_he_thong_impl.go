package repoimpl

import (
	"hmdl-user-service/db/core"
	. "hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
	"errors"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

func NewDmThamSoHeThongRepo(db *core.DbData) repository.DmThamSoHeThongRepo {
	return &DmThamSoHeThongRepoImpl{db: db}
}

type DmThamSoHeThongRepoImpl struct {
	db *core.DbData
}

func (u *DmThamSoHeThongRepoImpl) GetAll(ctx echo.Context) ([]DmThamSoHeThong, error) {
	var data []DmThamSoHeThong

	err := u.db.DbSql01.Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *DmThamSoHeThongRepoImpl) GetById(ctx echo.Context, id int) (*DmThamSoHeThong, error) {
	var dsPhongKha DmThamSoHeThong
	err := u.db.DbSql01.Where("Id = ?", id).Find(&dsPhongKha).Error

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

func (u *DmThamSoHeThongRepoImpl) Delete(ctx echo.Context, id int) error {
	var data DmThamSoHeThong

	err := u.db.DbSql01.Find(&data, id).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	if gorm.IsRecordNotFoundError(err) {
		return errors.New("Không tìm thấy dữ liệu")
	}

	err = u.db.DbSql01.Delete(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}

func (u *DmThamSoHeThongRepoImpl) Insert(ctx echo.Context, item DmThamSoHeThong) (*DmThamSoHeThong, error) {
	err := u.db.DbSql01.Create(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return &item, nil
}

func (u *DmThamSoHeThongRepoImpl) Update(ctx echo.Context, item DmThamSoHeThong) error {
	err := u.db.DbSql01.Save(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}
