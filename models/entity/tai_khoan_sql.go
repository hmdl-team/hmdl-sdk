package entity

type TaiKhoanSQL struct {
	ID          int    `json:"id"`
	TenTaiKhoan string `json:"ten_tai_khoan"`
	MatKhau     string `json:"mat_khau"`
	TenNhanVien string `json:"ten_nhan_vien"`
	NhanVienID  int    `json:"nhan_vien_id"`
	PhanQuyenID int    `json:"phan_quyen_id"`
	CapDo       string `json:"cap_do"`
	TinhTrang   bool   `json:"tinh_trang"`
}
