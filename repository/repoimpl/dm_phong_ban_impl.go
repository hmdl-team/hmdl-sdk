package repoimpl

import (
	"context"
	"errors"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	. "hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
)

func NewDM_PhongBanRepo(db *gorm.DB) repository.DM_PhongBanRepo {
	return &dmPhongbanrepoimpl{db: db}
}

type dmPhongbanrepoimpl struct {
	db *gorm.DB
}

func (u *dmPhongbanrepoimpl) GetCayPhongBanByUserId(ctx context.Context, userId int) ([]DM_PhongBan, error) {
	var data []DM_PhongBan

	var dsPhong []int

	err := u.db.Raw(`
		SELECT
		  dpb.DM_PhongBanID
		FROM DM_PhongBan dpb
		
		WHERE dpb.DM_PhongBanID NOT IN (SELECT
			CAST(dtsnd.GiaTriThamSo AS INT)
		  FROM DM_ThamSoNguoiDung dtsnd
		  WHERE dtsnd.MaThamSo = 'PHONGBAN'
		  AND dtsnd.DM_NhanVien_DM_NhanVienID = (SELECT
			  dtk.DM_NhanVienID
			FROM DM_TaiKhoan dtk
    WHERE dtk.DM_TaiKhoanID = ?))
	`, userId).Pluck("PhongBanId", &dsPhong).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}
	//"DM_PhongBanID not in (?)", dsPhong
	err = u.db.
		Preload("Children", "DM_PhongBanID not in (?)", dsPhong).
		Preload("Children.Children", "DM_PhongBanID not in (?)", dsPhong).
		Preload("Children.Children.Children", "DM_PhongBanID not in (?)", dsPhong).
		Where("ParentKeyId is null").Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *dmPhongbanrepoimpl) GetCayPhongBan(c echo.Context) ([]DM_PhongBan, error) {
	var data []DM_PhongBan

	err := u.db.Set("gorm:auto_preload", true).Where("ParentKeyId is null").Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *dmPhongbanrepoimpl) GetAll(ctx echo.Context) ([]DM_PhongBan, error) {
	var data []DM_PhongBan

	err := u.db.Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *dmPhongbanrepoimpl) GetPhongBanComBobox(ctx echo.Context) ([]DM_PhongBan, error) {
	var data []DM_PhongBan

	err := u.db.Where(&DM_PhongBan{TinhTrang: true}).Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *dmPhongbanrepoimpl) GetById(ctx echo.Context, id int) (*DM_PhongBan, error) {
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

func (u *dmPhongbanrepoimpl) Delete(ctx echo.Context, id int) error {
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

func (u *dmPhongbanrepoimpl) Insert(ctx echo.Context, item DM_PhongBan) (*DM_PhongBan, error) {
	err := u.db.Create(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return &item, nil
}

func (u *dmPhongbanrepoimpl) Update(ctx echo.Context, item DM_PhongBan) error {
	err := u.db.Save(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}
