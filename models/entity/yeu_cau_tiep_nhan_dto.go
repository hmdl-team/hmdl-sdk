package entity

import (
	"github.com/jinzhu/gorm/dialects/postgres"
	"hmdl-user-service/models/data_user"

	"time"
)

type YeuCauTiepNhanDTO struct {
	TaiKhoanTiepNhan    string           `json:"tai_khoan_tiep_nhan"`
	TenNhanVienTiepNhan string           `json:"ten_nhan_vien_tiep_nhan"`
	Id                  int              `json:"id"`
	TenBenhNhan         string           `json:"ten_benh_nhan"`
	SoVaoVien           string           `json:"so_vao_vien"`
	SoTiepNhan          string           `json:"so_tiep_nhan"`
	NgaySinh            *time.Time       `json:"ngay_sinh"`
	GioiTinh            *bool               `json:"gioi_tinh"`
	TinhThanhId         *int                `json:"tinh_thanh_id"`
	QuanHuyenId         *int                `json:"quan_huyen_id"`
	PhuongXaId          *int                `json:"phuong_xa_id"`
	DiaChi              string              `json:"dia_chi"`
	SoDienThoai         string              `json:"so_dien_thoai"`
	Email               string              `json:"email"`
	TenNguoiLienHe      string              `json:"ten_nguoi_lien_he"`
	TamNgung            *bool               `json:"tam_ngung"`
	NguoiYeuCauId       *int                `json:"nguoi_yeu_cau_id"`
	NguoiYeuCau         *data_user.NhanVien `json:"nguoi_yeu_cau" gorm:"foreignkey:NguoiYeuCauId"`
	NhanVienTiepNhanId  *int                `json:"nhan_vien_tiep_nhan_id"`
	NhanVienTiepNhan    *data_user.NhanVien `json:"nhan_vien_tiep_nhan" gorm:"foreignkey:NhanVienTiepNhanId"`
	NoiDung             string              `json:"noi_dung"`
	ChanDoan            string              `json:"chan_doan"`
	NgayYeuCau          *time.Time          `json:"ngay_yeu_cau"`
	NgayTiepNhan        *time.Time          `json:"ngay_tiep_nhan"`
	HoanThanh           *bool               `json:"hoan_thanh"`
	DaTiepNhan          *bool               `json:"da_tiep_nhan"`
	IpTiepNhan          string              `json:"-"`
	NamSinh             *int                `json:"nam_sinh"`
	BacSiYeuCau         postgres.Jsonb      `json:"bac_si_yeu_cau"`
}
