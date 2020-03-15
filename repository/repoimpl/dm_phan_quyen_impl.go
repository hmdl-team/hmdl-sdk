package repoimpl

import (

	"errors"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
)

func NewDmPhanQuyenRepo(db *gorm.DB) repository.DmPhanQuyenRepo {
	return &dmPhanQuyenRepoImpl{db: db}
}

type dmPhanQuyenRepoImpl struct {
	db *gorm.DB
}

func (u *dmPhanQuyenRepoImpl) GetAll(ctx echo.Context) ([]data_user.DM_PhanQuyen, error) {
	var data []data_user.DM_PhanQuyen

	err := u.db.Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *dmPhanQuyenRepoImpl) GetById(ctx echo.Context, id int) (*data_user.DM_PhanQuyen, error) {
	var dsPhongKha data_user.DM_PhanQuyen
	err := u.db.Find(&dsPhongKha, id).Error

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

func (u *dmPhanQuyenRepoImpl) Delete(ctx echo.Context, id int) error {
	var data data_user.DM_PhanQuyen

	err := u.db.First(&data, id).Error
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

func (u *dmPhanQuyenRepoImpl) Insert(ctx echo.Context, item data_user.DM_PhanQuyen) (*data_user.DM_PhanQuyen, error) {
	err := u.db.Create(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return &item, nil
}

func (u *dmPhanQuyenRepoImpl) Update(ctx echo.Context, item data_user.DM_PhanQuyen) error {
	err := u.db.Save(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}
