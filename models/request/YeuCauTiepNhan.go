package request

import "time"

type YeuCauTiepNhan struct {
	TuNgay           *time.Time `json:"tu_ngay,omitempty"`
	DenNgay          *time.Time `json:"den_ngay,omitempty"`
	NguoiYeuCau      *int       `json:"nguoi_yeu_cau,omitempty"`
	TinhTrangXacNhan bool       `json:"tinh_trang_xac_nhan,omitempty"`
}
