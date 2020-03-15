package data_user

import "time"

type DM_TaiKhoan struct {
	DM_TaiKhoanId    int           `gorm:"column:DM_TaiKhoanID;primary_key;AUTO_INCREMENT" json:"dm_tai_khoan_id"`
	TenTaiKhoan      string        `gorm:"column:TenTaiKhoan;" json:"ten_tai_khoan, omitempty"`
	MatKhau          string        `gorm:"column:MatKhau;" json:"_, omitempty"`
	DM_NhanVienId    int           `gorm:"column:DM_NhanVienID;" json:"dm_nhan_vien_id, omitempty"`
	NhanVienSuDung   *NhanVien     `gorm:"ForeignKey:DM_NhanVienID;AssociationForeignKey:DM_NhanVienID" json:"nhan_vien_su_dung"`
	NgayTao          *time.Time    `gorm:"column:NgayTao;" json:"ngay_tao, omitempty"`
	ThoiGianCapNhat  *time.Time    `gorm:"column:ThoiGianCapNhat;" json:"thoi_gian_cap_nhat, omitempty"`
	NguoiCapNhatId   *int          `gorm:"column:NguoiCapNhatId;" json:"nguoi_cap_nhat_id, omitempty"`
	TinhTrang        bool          `gorm:"column:TinhTrang;" json:"tinh_trang, omitempty"`
	PhanQuyenId      *int          `gorm:"column:PhanQuyenId;" json:"phan_quyen_id, omitempty"`
	DonVi            string        `gorm:"column:DonVi;" json:"don_vi, omitempty"`
	DM_PhanQuyenID   *int          `gorm:"column:DM_PhanQuyenID;" json:"dm_phan_quyen_id, omitempty"`
	DM_PhanQuyen     *DM_PhanQuyen `gorm:"ForeignKey:DM_PhanQuyenID;AssociationForeignKey:DM_PhanQuyenID" json:"dm_phan_quyen"`
	DomainLogin      bool          `gorm:"column:DomainLogin;" json:"domain_login, omitempty"`
	ThoiGianDangNhap *time.Time    `gorm:"column:ThoiGianDangNhap;" json:"thoi_gian_dang_nhap, omitempty"`
	ThoiGianDangXuat *time.Time    `gorm:"column:ThoiGianDangXuat;" json:"thoi_gian_dang_xuat, omitempty"`
	TuDongDangNhap   bool          `gorm:"column:TuDongDangNhap;" json:"tu_dong_dang_nhap, omitempty"`
	MatKhauWeb       string        `gorm:"column:MatKhauWeb;" json:"_, omitempty"`
}

func (DM_TaiKhoan) TableName() string {
	return "DM_TaiKhoan"
}
