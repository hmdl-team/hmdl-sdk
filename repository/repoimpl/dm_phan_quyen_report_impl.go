package repoimpl

import (
	"errors"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	. "hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
)

func NewDM_PhanQuyen_ReportRepo(db *gorm.DB) repository.DM_PhanQuyen_ReportRepo {
	return &DM_PhanQuyen_ReportRepoImpl{db: db}
}

type DM_PhanQuyen_ReportRepoImpl struct {
	db *gorm.DB
}

func (u *DM_PhanQuyen_ReportRepoImpl) GetAll(ctx echo.Context) ([]DM_PhanQuyen_Report, error) {
	var data []DM_PhanQuyen_Report

	err := u.db.Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *DM_PhanQuyen_ReportRepoImpl) GetById(ctx echo.Context, id int) (*DM_PhanQuyen_Report, error) {
	var dsPhongKha DM_PhanQuyen_Report
	err := u.db.Where("DM_PhanQuyen_ReportId = ?", id).Find(&dsPhongKha).Error

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

func (u *DM_PhanQuyen_ReportRepoImpl) Delete(ctx echo.Context, id int) error {
	var data DM_PhanQuyen_Report

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

func (u *DM_PhanQuyen_ReportRepoImpl) Insert(ctx echo.Context, item DM_PhanQuyen_Report) (*DM_PhanQuyen_Report, error) {
	err := u.db.Create(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return &item, nil
}

func (u *DM_PhanQuyen_ReportRepoImpl) Update(ctx echo.Context, item DM_PhanQuyen_Report) error {
	err := u.db.Save(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}
