package repoimpl

import (
	"errors"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"hmdl-user-service/models/data_user"
	"hmdl-user-service/models/request"
	"hmdl-user-service/repository"
	"strings"
)

func NewDmTaiKhoanRepo(db *gorm.DB) repository.TaiKhoanRepository {
	return &dmTaiKhoanRepoImpl{db: db}
}

type dmTaiKhoanRepoImpl struct {
	db *gorm.DB
}

func (u *dmTaiKhoanRepoImpl) Login(ctx echo.Context, loginReq request.ReqSignIn) (data_user.DM_TaiKhoan, error) {
	data := data_user.DM_TaiKhoan{}
	err := u.db.
		Where(" LOWER(TenTaiKhoan) like ?", strings.ToLower(loginReq.UserName)).
		Set("gorm:auto_preload", true).
		Find(&data).
		Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return data, err
	}
	return data, nil
}

func (u *dmTaiKhoanRepoImpl) GetAll(ctx echo.Context) ([]data_user.DM_TaiKhoan, error) {
	var data []data_user.DM_TaiKhoan

	err := u.db.Set("gorm:auto_preload", true).Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *dmTaiKhoanRepoImpl) GetById(ctx echo.Context, id int) (*data_user.DM_TaiKhoan, error) {
	var dsPhongKha data_user.DM_TaiKhoan
	err := u.db.Set("gorm:auto_preload", true).Find(&dsPhongKha, id).Error

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

func (u *dmTaiKhoanRepoImpl) Delete(ctx echo.Context, id int) error {
	var data data_user.DM_TaiKhoan

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

func (u *dmTaiKhoanRepoImpl) Insert(ctx echo.Context, item data_user.DM_TaiKhoan) (*data_user.DM_TaiKhoan, error) {
	err := u.db.Create(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return &item, nil
}

func (u *dmTaiKhoanRepoImpl) Update(ctx echo.Context, item data_user.DM_TaiKhoan) error {
	err := u.db.Save(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}
