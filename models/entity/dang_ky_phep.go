package entity

import "time"

type DangKyPhep struct {
	NhanVienID  int       `json:"nhan_vien_id"`
	NgayDangKy  time.Time `json:"ngay_dang_ky"`
	CaLamViecID string    `json:"ca_lam_viec_id"`
	LyDo        string    `json:"ly_do"`
}
