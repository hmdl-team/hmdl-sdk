package repoimpl

import (
	"errors"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"hmdl-user-service/core"
	. "hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
)

func NewDmChucVuRepo(db *core.DbData) repository.DmChucVuRepo {
	return &dmChucVuRepoImpl{db: db}
}

type dmChucVuRepoImpl struct {
	db *core.DbData
}

func (u *dmChucVuRepoImpl) GetAll(ctx echo.Context) ([]DmChucVu, error) {
	var data []DmChucVu
	err := u.db.DbSql01.Find(&data).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}
	return data, err
}

func (u *dmChucVuRepoImpl) GetById(ctx echo.Context, id int) (*DmChucVu, error) {
	var dsPhongKha DmChucVu
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

func (u *dmChucVuRepoImpl) Delete(ctx echo.Context, id int) error {
	var data DmChucVu

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

func (u *dmChucVuRepoImpl) Insert(ctx echo.Context, item DmChucVu) (*DmChucVu, error) {
	err := u.db.DbSql01.Create(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return &item, nil
}

func (u *dmChucVuRepoImpl) Update(ctx echo.Context, item DmChucVu) error {
	err := u.db.DbSql01.Save(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}
