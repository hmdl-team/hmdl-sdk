package data_user

type DM_NhanVienQuanLy struct {
	DM_NhanVienQuanLyId int `gorm:"column:DM_NhanVienQuanLy;primary_key;AUTO_INCREMENT" json:"danh_sach_nhan_vien_quan_ly_id"`
	NhanVienId          int `gorm:"column:NhanVienId" json:"nhan_vien_id"`
	QuanLyCapTrenId     int `gorm:"column:QuanLyCapTrenId" json:"quan_ly_cap_tren_id"`
	Cap                 int `gorm:"column:Cap" json:"cap"`
}

func (DM_NhanVienQuanLy) TableName() string {
	return "DM_NhanVienQuanLy"
}
