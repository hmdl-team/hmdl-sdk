package repoimpl

import (
	"database/sql"
	"github.com/getsentry/raven-go"
	"github.com/jinzhu/gorm"
	"hmdl-user-service/models/data_user"
	"hmdl-user-service/models/entity"
	"hmdl-user-service/repository"
)

//NewTaiKhoanRepo : khởi tạo
func NewDanhMucHeThongRePo(dbsql *gorm.DB, dbpos *gorm.DB) repository.DanhMucHeThongRepository {
	return &DanhMucHeThongImpl{
		DbSql: dbsql,
		DbPos: dbpos,
	}
}

//khởi tạo
type DanhMucHeThongImpl struct {
	DbSql *gorm.DB
	DbPos *gorm.DB
}

func (u *DanhMucHeThongImpl) GetAllChucDanh() ([]data_user.DanhMucHeThong, error) {
	return u.GetAllDanhMucHeThongByLoaiDanhMuc("ChucDanh")
}

func (u *DanhMucHeThongImpl) GetAllChucVu() ([]data_user.DanhMucHeThong, error) {
	return u.GetAllDanhMucHeThongByLoaiDanhMuc("ChucVu")
}

func (u *DanhMucHeThongImpl) GetAllDanhMucHeThongByLoaiDanhMuc(LoaiDanhMuc string) ([]data_user.DanhMucHeThong, error) {
	data := make([]data_user.DanhMucHeThong, 0)
	err := u.DbPos.Where("loai_danh_muc = ?", LoaiDanhMuc).Find(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}
	if data != nil {
		return data, nil
	}
	return nil, nil
}

func (u *DanhMucHeThongImpl) GetDanhMucHeThongByDanhMucCode(DanhMucCode int, LoaiDanhMuc string) (*data_user.DanhMucHeThong, error) {
	data := &data_user.DanhMucHeThong{}
	err := u.DbPos.Where("danh_muc_code = ? and loai_danh_muc = ?", DanhMucCode, LoaiDanhMuc).First(&data).Error

	if err != nil && !gorm.IsRecordNotFoundError(err) {
		raven.CaptureErrorAndWait(err, nil)
		return nil, err
	}
	if data != nil {
		return data, nil
	}
	return nil, nil
}

func (u *DanhMucHeThongImpl) DongBoChucdanh() error {
	data := make([]entity.DmChucdanh, 0)
	err := u.DbSql.Raw(`
	SELECT
	  dcd.DM_ChucDanhId AS chucdanhid
	 ,dcd.TenChucDanh AS ten_chuc_danh
	 ,dcd.TenChucDanhVietTat AS ten_chuc_danh_viet_tat
	 ,dcd.GhiChu AS ghi_chu
	FROM DM_ChucDanh dcd
	`).Scan(&data).Error

	if err != nil && err != sql.ErrNoRows {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	if len(data) == 0 {
		return nil
	}

	for _, item := range data {

		danhMucCu, _ := u.GetDanhMucHeThongByDanhMucCode(item.Chucdanhid, "ChucDanh")

		if danhMucCu == nil {
			dm := &data_user.DanhMucHeThong{
				DanhMucCode: item.Chucdanhid,
				LoaiDanhMuc: "ChucDanh",
				TenDanhMuc:  item.TenChucDanh,
				VietTat:     item.TenChucDanhVietTat,
				GhiChu:      item.GhiChu,
			}
			u.Insert(dm)
		}

	}

	return nil
}

func (u *DanhMucHeThongImpl) DongBoChucVu() error {
	data := make([]entity.DmChucVu, 0)
	err := u.DbSql.Raw(`
	SELECT
		  dcv.DM_ChucVuId AS chuc_vu_id
		 ,dcv.TenChucVu AS ten_chuc_vu
		 ,dcv.TenChucVuVietTat AS ten_viet_tat
		 ,dcv.GhiChu AS ghi_chu
	FROM DM_ChucVu dcv 
	`).Scan(&data).Error

	if err != nil && err != sql.ErrNoRows {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}

	if len(data) == 0 {
		return nil
	}

	for _, item := range data {

		data, _ := u.GetDanhMucHeThongByDanhMucCode(item.ChucVuId, "ChucVu")

		if data == nil {
			dm := &data_user.DanhMucHeThong{
				DanhMucCode: item.ChucVuId,
				LoaiDanhMuc: "ChucVu",
				TenDanhMuc:  item.TenChucVu,
				VietTat:     item.TenVietTat,
				GhiChu:      item.GhiChu,
			}
			_ = u.Insert(dm)
		}

	}

	return nil
}

func (u *DanhMucHeThongImpl) Insert(item *data_user.DanhMucHeThong) error {
	err := u.DbPos.Create(&item).Error
	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		return err
	}
	return nil
}
