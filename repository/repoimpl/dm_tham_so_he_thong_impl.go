package repoimpl

import (
	"context"
	"errors"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"hmdl-user-service/core"
	. "hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
)

func NewDmThamSoHeThongRepo(db *core.DbData) repository.DmThamSoHeThongRepo {
	return &dmThamSoHeThongRepoImpl{
		db: db,
	}
}

type dmThamSoHeThongRepoImpl struct {
	db *core.DbData
}

func (u *dmThamSoHeThongRepoImpl) GetThamSoValueByCode(ctx context.Context, code string) (*DmThamSoHeThong, error) {
	var data DmThamSoHeThong

	err := u.db.DbSql01.Raw("select top 1 ts.* From  DM_ThamSoHeThong ts where ts.MaThamSo = ?",code).Scan(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}
	return &data, err
}

func (u *dmThamSoHeThongRepoImpl) GetThamSoByCode(ctx context.Context, code string) ([]DmThamSoHeThong, error) {
	var data []DmThamSoHeThong

	err := u.db.DbSql01.Where(&DmThamSoHeThong{
		MaThamSo: code,
	}).Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *dmThamSoHeThongRepoImpl) GetAll(ctx echo.Context) ([]DmThamSoHeThong, error) {
	var data []DmThamSoHeThong

	err := u.db.DbSql01.Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *dmThamSoHeThongRepoImpl) GetById(ctx echo.Context, id int) (*DmThamSoHeThong, error) {
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

func (u *dmThamSoHeThongRepoImpl) Delete(ctx echo.Context, id int) error {
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

func (u *dmThamSoHeThongRepoImpl) Insert(ctx echo.Context, item DmThamSoHeThong) (*DmThamSoHeThong, error) {
	err := u.db.DbSql01.Create(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return &item, nil
}

func (u *dmThamSoHeThongRepoImpl) Update(ctx echo.Context, item DmThamSoHeThong) error {
	err := u.db.DbSql01.Save(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}
