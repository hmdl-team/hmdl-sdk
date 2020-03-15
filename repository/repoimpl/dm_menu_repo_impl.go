package repoimpl

import (

	"errors"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
)

func NewMenuWebRepo(db *gorm.DB) repository.MenuRepository {
	return &MenuWebRepoImpl{db: db}
}

type MenuWebRepoImpl struct {
	db *gorm.DB
}

func (u *MenuWebRepoImpl) GetMenuByPhanQuyenId(phanQuyenId int,duAnId int) ([]data_user.DM_MenuWeb, error) {
	var data []data_user.DM_MenuWeb

	err := u.db.Table("DM_PhanQuyenMenu").
		Select("DM_MenuWeb.*").
		Joins("left join DM_MenuWeb  on DM_PhanQuyenMenu.DM_MenuWebId = DM_MenuWeb.Id").
		Where("DM_DuAnId = ? and DM_MenuWeb.enable = 1 and DM_PhanQuyenMenu.DM_PhanQuyenID = ? and DM_MenuWeb.ParentId is null", duAnId,phanQuyenId).
		Order("OrderBy ASC").
		Find(&data).Error

	for i, item := range data {
		var chilData []*data_user.DM_MenuWeb

		err := u.db.Table("DM_PhanQuyenMenu").
			Select("DM_MenuWeb.*").
			Joins("left join DM_MenuWeb on DM_PhanQuyenMenu.DM_MenuWebId = DM_MenuWeb.Id").
			Where("DM_DuAnId = ? and DM_MenuWeb.Enable = 1 and DM_PhanQuyenMenu.DM_PhanQuyenID = ? and  DM_MenuWeb.ParentId = ? ", duAnId,phanQuyenId, item.Id).
			Order("OrderBy ASC").
			Find(&chilData).Error

		if err != nil && !gorm.IsRecordNotFoundError(err) {
			raven.CaptureErrorAndWait(err, nil)
			return nil, err
		}

		data[i].Children = chilData

	}

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, nil
}

func (u *MenuWebRepoImpl) GetAll(ctx echo.Context) ([]data_user.DM_MenuWeb, error) {
	var data []data_user.DM_MenuWeb

	err := u.db.Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *MenuWebRepoImpl) GetById(ctx echo.Context, id int) (*data_user.DM_MenuWeb, error) {
	var dsPhongKha data_user.DM_MenuWeb
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

func (u *MenuWebRepoImpl) Delete(ctx echo.Context, id int) error {
	var data data_user.DM_MenuWeb

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

func (u *MenuWebRepoImpl) Insert(ctx echo.Context, item data_user.DM_MenuWeb) (*data_user.DM_MenuWeb, error) {
	err := u.db.Create(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return &item, nil
}

func (u *MenuWebRepoImpl) Update(ctx echo.Context, item data_user.DM_MenuWeb) error {
	err := u.db.Save(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}
