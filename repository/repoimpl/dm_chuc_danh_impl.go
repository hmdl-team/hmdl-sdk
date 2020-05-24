package repoimpl

import (
	"errors"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"hmdl-user-service/db/core"
	. "hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
)

func NewDmChucDanhRepo(db *core.DbData) repository.DmChucDanhRepo {
	return &dmChucDanhRepoImpl{db: db}
}

type dmChucDanhRepoImpl struct {
	db *core.DbData
}

func (u *dmChucDanhRepoImpl) GetAll(ctx echo.Context) ([]DmChucDanh, error) {
	var data []DmChucDanh

	err := u.db.DbSql01.Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *dmChucDanhRepoImpl) GetById(ctx echo.Context, id int) (*DmChucDanh, error) {
	var dsPhongKha DmChucDanh
	err := u.db.DbSql01.Find(&dsPhongKha, id).Error

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

func (u *dmChucDanhRepoImpl) Delete(ctx echo.Context, id int) error {
	var data DmChucDanh

	err := u.db.DbSql01.First(&data, id).Error
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

func (u *dmChucDanhRepoImpl) Insert(ctx echo.Context, item DmChucDanh) (*DmChucDanh, error) {
	err := u.db.DbSql01.Create(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return &item, nil
}

func (u *dmChucDanhRepoImpl) Update(ctx echo.Context, item DmChucDanh) error {
	err := u.db.DbSql01.Save(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}
