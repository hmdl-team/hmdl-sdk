package repoimpl

import (
	"database/sql"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"

	"hmdl-user-service/models/data_user"
	"hmdl-user-service/repository"
)

func NewPhanQuyenMenuRepo(db *gorm.DB) repository.PhanQuyenMenuRepository {
	return &PhanQuyenMenuRepoImpl{
		Db: db,
	}
}

type PhanQuyenMenuRepoImpl struct {
	Db *gorm.DB
}

func (u *PhanQuyenMenuRepoImpl) GetMenuByPhanQuyenId(phanQuyenId int, duAnId int) ([]data_user.DM_MenuWeb, error) {
	var data []data_user.DM_MenuWeb

	err := u.Db.Table("DM_PhanQuyenMenu").
		Select("DM_MenuWeb.*").
		Joins("left join DM_MenuWeb on DM_PhanQuyenMenu.menu_id = DM_MenuWeb.id").
		Where("DM_DuAnId = ? and DM_MenuWeb.enable = true and DM_PhanQuyenMenu.phan_quyen_id = ? and parent_id is null", duAnId, phanQuyenId).
		Order("OrderBy asc").
		Find(&data).Error

	for i, item := range data {
		var chilData []*data_user.DM_MenuWeb

		err := u.Db.Table("DM_PhanQuyenMenu").
			Select("DM_MenuWeb.*").
			Joins("left join DM_MenuWeb on DM_PhanQuyenMenu.menu_id = DM_MenuWeb.id").
			Where("DM_DuAnId = ? and DM_MenuWeb.enable = true and DM_PhanQuyenMenu.phan_quyen_id = ? and  DM_MenuWeb.parent_id = ? ", duAnId, phanQuyenId, item.Id).
			Order("order_by asc").
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

func (u *PhanQuyenMenuRepoImpl) GetAllPhanQuyenMenu() ([]data_user.DM_PhanQuyenMenu, error) {
	var data []data_user.DM_PhanQuyenMenu

	err := u.Db.Preload("DM_MenuWeb").Preload("DM_PhanQuyen").Find(&data).Error

	if err != nil && err != sql.ErrNoRows {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}
	if data != nil {
		return data, nil
	}
	return nil, nil
}

func (u *PhanQuyenMenuRepoImpl) Delete(id int) error {
	data := &data_user.DM_PhanQuyenMenu{}
	err := u.Db.Where("id = ?", id).Find(&data).Error

	if err != nil && err != sql.ErrNoRows {
		raven.CaptureErrorAndWait(err, nil)
		return err
	} else if data != nil {
		err := u.Db.Delete(&data).Error
		if err != nil {
			raven.CaptureErrorAndWait(err, nil)
			return err
		}

	}

	return nil
}

func (u *PhanQuyenMenuRepoImpl) GetById(id int) (*data_user.DM_PhanQuyenMenu, error) {
	data := &data_user.DM_PhanQuyenMenu{}
	err := u.Db.First(&data, id).Error
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}
	if data != nil {
		return data, nil
	}
	return nil, nil
}

func (u *PhanQuyenMenuRepoImpl) Insert(Menu *data_user.DM_PhanQuyenMenu) error {

	data := &data_user.DM_PhanQuyenMenu{}
	err := u.Db.Where("DM_PhanQuyenID = ? and DM_MenuWebId = ?", Menu.DM_PhanQuyenID, Menu.DM_MenuWebId).Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}
	if data.DM_PhanQuyenID > 0 {
		return nil
	}

	err = u.Db.Create(&Menu).Error

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}
	return nil
}

func (u *PhanQuyenMenuRepoImpl) Update(Menu *data_user.DM_PhanQuyenMenu) error {
	err := u.Db.Save(&Menu).Error
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}
	return nil
}
