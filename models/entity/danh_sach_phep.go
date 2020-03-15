package entity

import "time"

type DanhSachPhep struct {
	PhieuYeuCauID      int       `json:"phieu_yeu_cau_id"`
	NhanVienThucHienID int       `json:"nhan_vien_thuc_hien_id"`
	TenNhanVien        string    `json:"ten_nhan_vien"`
	TruongBoPhan       string    `json:"truong_bo_phan"`
	NgayPhep           time.Time `json:"ngay_phep"`
	TruongBoPhanDuyet  bool      `json:"truong_bo_phan_duyet"`
	NhanSuDuyet        bool      `json:"nhan_su_duyet"`
	DuyetBanGiamDoc    bool      `json:"duyet_ban_giam_doc"`
	TenCa              string    `json:"ten_ca"`
	LyDo               string    `json:"ly_do"`
	ThoiGianYeuCau     time.Time `json:"thoi_gian_yeu_cau"`
	ThoiGianDuyet      time.Time `json:"thoi_gian_duyet"`
	Cap                int       `json:"cap"`
}
