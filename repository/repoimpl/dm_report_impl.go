package repoimpl

import (
	"errors"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	. "hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
)

func NewDM_ReportRepo(db *gorm.DB) repository.DM_ReportRepo {
	return &DM_ReportRepoImpl{db: db}
}

type DM_ReportRepoImpl struct {
	db *gorm.DB
}

func (u *DM_ReportRepoImpl) GetReportTrangBaoCaoByPhanQuyenId(phanQuyenId int) (data []DM_Report, err error) {
	var dsReportId []DM_PhanQuyen_Report

	err = u.db.Model(&DM_PhanQuyen_Report{}).Where(&DM_PhanQuyen_Report{
		DM_PhanQuyenID: phanQuyenId,
	}).Find(&dsReportId).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return data, err
	}

	var dsIdWhere []int

	for _, item := range dsReportId {
		dsIdWhere = append(dsIdWhere, item.DM_ReportId)
	}

	var dataReport []DM_Report

	err = u.db.Where("DM_ReportId in (?)", dsIdWhere).Find(&dataReport).Error
	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return data, err
	}

	data = u.getReportChildren(dataReport, 0)

	return data, nil
}

func (u *DM_ReportRepoImpl) getReportChildren(dataInput []DM_Report, parentId int) (data []DM_Report) {

	for _, item := range dataInput {
		if *item.ParentId == parentId {
			resul := u.getReportChildren(dataInput, item.DM_ReportId)
			item.Children = resul
			data = append(data, item)
		}
	}

	return data
}

func (u *DM_ReportRepoImpl) GetReportPhanQuyenId(phanQuyenId int) (data []DM_Report, err error) {

	var dsReportId []DM_PhanQuyen_Report

	err = u.db.Model(&DM_PhanQuyen_Report{}).Where(&DM_PhanQuyen_Report{
		DM_PhanQuyenID: phanQuyenId,
	}).Find(&dsReportId).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return data, err
	}
	var dsIdWhere []int
	for _, item := range dsReportId {
		dsIdWhere = append(dsIdWhere, item.DM_ReportId)
	}

	err = u.db.Set("gorm:auto_preload", true).Where("DM_ReportId in (?)", dsIdWhere).Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return data, err
	}
	return data, nil
}

func (u *DM_ReportRepoImpl) GetAll(ctx echo.Context) ([]DM_Report, error) {
	var data []DM_Report

	err := u.db.Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *DM_ReportRepoImpl) GetById(ctx echo.Context, id int) (*DM_Report, error) {
	var dsPhongKha DM_Report
	err := u.db.Where("DM_ReportId = ?", id).Find(&dsPhongKha).Error

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

func (u *DM_ReportRepoImpl) Delete(ctx echo.Context, id int) error {
	var data DM_Report

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

func (u *DM_ReportRepoImpl) Insert(ctx echo.Context, item DM_Report) (*DM_Report, error) {
	err := u.db.Create(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return &item, nil
}

func (u *DM_ReportRepoImpl) Update(ctx echo.Context, item DM_Report) error {
	err := u.db.Save(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}
