package entity

type NhanVien struct {
	DM_NhanVienID     int    `json:"dm_nhan_vien_id"`
	MaNhanVien        string `json:"ma_nhan_vien"`
	MaChamCong        string `json:"ma_cham_cong"`
	TenNhanVien       string `json:"ten_nhan_vien"`
	GioiTinh          bool   `json:"gioi_tinh"`
	TinhTrang         bool   `json:"tinh_trang"`
	NhanVienChinhThuc bool   `json:"nhan_vien_chinh_thuc"`
	PhongID           int    `json:"phong_id"`
	ChucDanh          string `json:"chuc_danh"`
	ChucVu            string `json:"chuc_vu"`
	TenPhongBan       string `json:"ten_phong_ban"`
	Email             string `json:"email"`
	QuanLyTrucTiepID  int    `json:"quan_ly_truc_tiep_id"`
	SoTheChamCong     string `json:"so_the_cham_cong"`
	DM_ChucDanhID     int    `json:"dm_chuc_danh_id"`
	DM_ChucVuID       int    `json:"dm_chuc_vu_id"`
	TaiKhoanDomain    string `json:"tai_khoan_domain"`
	FullName          string `json:"full_name"`
	LastName          string `json:"last_name"`
	MidName           string `json:"mid_name"`
	FirstName         string `json:"first_name"`
	NhanThongBao      bool   `json:"nhan_thong_bao"`
}
