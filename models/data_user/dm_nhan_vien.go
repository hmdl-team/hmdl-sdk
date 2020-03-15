package data_user

type NhanVien struct {
	DM_NhanVienId     int    `gorm:"column:DM_NhanVienID; primary_key;AUTO_INCREMENT" json:"dm_nhan_vien_id"`
	MaNhanVien        string `gorm:"column:MaNhanVien;" json:"ma_nhan_vien"`
	MaChamCong        string `gorm:"column:MaChamCong;" json:"ma_cham_cong"`
	TenNhanVien       string `gorm:"column:TenNhanVien;" json:"ten_nhan_vien"`
	FullName          string `gorm:"column:FullName;" json:"full_name"`
	LastName          string `gorm:"column:LastName;" json:"last_name"`
	MidName           string `gorm:"column:MidName;" json:"mid_name"`
	FirstName         string `gorm:"column:FirstName;" json:"first_name"`
	GioiTinh          bool   `gorm:"column:GioiTinh;"  json:"gioi_tinh"`
	TinhTrang         bool   `gorm:"column:TinhTrang;"  json:"tinh_trang"`
	NhanVienChinhThuc bool   `gorm:"column:NhanVienChinhThuc;" json:"nhan_vien_chinh_thuc"`
	PhongID           int    `gorm:"column:PhongID;" json:"phong_id"`
	ChucDanh          string `gorm:"column:ChucDanh;" json:"chuc_danh"`
	ChucVu            string `gorm:"column:ChucVu;" json:"chuc_vu"`
	TenPhongBan       string `gorm:"column:TenPhongBan;" json:"ten_phong_ban"`
	Email             string `gorm:"column:Email;" json:"email"`
	QuanLyTrucTiepID  int    `gorm:"column:QuanLyTrucTiepID;" json:"quan_ly_truc_tiep_id"`
	SoTheChamCong     string `gorm:"column:SoTheChamCong;" json:"so_the_cham_cong"`
	DM_ChucDanhID     int    `gorm:"column:DM_ChucDanhID;" json:"dm_chuc_danh_id"`
	DM_ChucVuID       int    `gorm:"column:DM_ChucVuID;" json:"dm_chuc_vu_id"`
	TaiKhoanDomain    string `gorm:"column:TaiKhoanDomain;" json:"tai_khoan_domain"`
	SoDienThoai       string `gorm:"column:SoDienThoai;" json:"so_dien_thoai"`
	NhanThongBao      bool   `gorm:"column:NhanThongBao;" json:"nhan_thong_bao"`
}

func (NhanVien) TableName() string {
	return "DM_NhanVien"
}
