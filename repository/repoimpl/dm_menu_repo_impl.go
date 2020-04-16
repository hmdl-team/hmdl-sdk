package repoimpl

import (
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
)

func NewMenuWebRepo(db *gorm.DB) repository.MenuRepository {
	return &menuWebRepoImpl{db: db}
}

type menuWebRepoImpl struct {
	db *gorm.DB
}

func (u *menuWebRepoImpl) getMenuChildren(dataInput []data_user.DM_MenuWeb, parentId int, cap int) (duLieu []data_user.DM_MenuWeb) {

	for _, item := range dataInput {
		if item.ParentId == parentId {
			resul := u.getMenuChildren(dataInput, item.Id,cap +1)
			item.Children = resul
			item.Cap = cap
			duLieu = append(duLieu, item)
		}
	}

	return duLieu
}

func (u *menuWebRepoImpl) GetMenuByPhanQuyenId(phanQuyenId int) (data []data_user.DM_MenuWeb, err error) {

	var dsReportId []data_user.DM_PhanQuyenMenu

	err = u.db.Model(&data_user.DM_PhanQuyenMenu{}).Where(&data_user.DM_PhanQuyenMenu{
		DM_PhanQuyenID: phanQuyenId,
	}).Find(&dsReportId).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return data, err
	}
	var dsIdWhere []int
	for _, item := range dsReportId {
		dsIdWhere = append(dsIdWhere, item.DM_MenuWebId)
	}

	err = u.db.Where("Id in (?)", dsIdWhere).Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return data, err
	}
	return data, nil
}

func (u *menuWebRepoImpl) GetMenuByPhanQuyenIdAndDuAnId(phanQuyenId int, duAnId int) (data []*data_user.DM_MenuWeb, err error) {
	var dataMenu []data_user.DM_PhanQuyenMenu

	err = u.db.Model(&data_user.DM_PhanQuyenMenu{}).Where(&data_user.DM_PhanQuyenMenu{
		DM_PhanQuyenID: phanQuyenId,
	}).Find(&dataMenu).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return data, err
	}

	var dsIdWhere []int

	for _, item := range dataMenu {
		dsIdWhere = append(dsIdWhere, item.DM_MenuWebId)
	}

	var dataMenuTheoPhanQuyen []data_user.DM_MenuWeb
	var dataParentId []*data_user.DM_MenuWeb

	err = u.db.Where("Id in (?) and DM_DuAnId = ?", dsIdWhere, duAnId).Find(&dataMenuTheoPhanQuyen).Error
	err = u.db.Where("Id in (?) and DM_DuAnId = ? and (ParentId is null or ParentId =0)", dsIdWhere, duAnId).Find(&dataParentId).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	for _, item := range dataParentId {
		item.Children = u.getMenuChildren(dataMenuTheoPhanQuyen, item.Id,item.Cap+1)
		data = append(data, item)
	}

	return data, nil

}

func (u *menuWebRepoImpl) GetAll(ctx echo.Context) ([]data_user.DM_MenuWeb, error) {
	var data []data_user.DM_MenuWeb

	err := u.db.Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return data, err
}

func (u *menuWebRepoImpl) GetById(ctx echo.Context, id int) (*data_user.DM_MenuWeb, error) {
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

func (u *menuWebRepoImpl) Delete(ctx echo.Context, id int) error {
	err := u.db.Delete(data_user.DM_MenuWeb{}, data_user.DM_MenuWeb{
		Id: id,
	}).Error

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}

func (u *menuWebRepoImpl) Insert(ctx echo.Context, item data_user.DM_MenuWeb) (*data_user.DM_MenuWeb, error) {
	err := u.db.Create(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}

	return &item, nil
}

func (u *menuWebRepoImpl) Update(ctx echo.Context, item data_user.DM_MenuWeb) error {
	err := u.db.Save(&item).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	return nil
}
