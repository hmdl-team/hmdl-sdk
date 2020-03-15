package data_user

type DM_PhongBan_NhanVien struct {
	PhongBanNhanVienID int  `json:"phong_ban_nhan_vien_id" gorm:"column:PhongBanNhanVienID; primary_key;AUTO_INCREMENT"`
	DmNhanvienId       int  `json:"dm_nhan_vien_id" gorm:"column:DM_NhanVienId"`
	DmPhongBanId       int  `json:"dm_phong_ban_id" gorm:"column:DM_PhongBanID"`
	TinhTrang          bool `json:"TinhTrang"`
	TruongBoPhan       bool `json:"TruongBoPhan"`
	TruongNhanSu       bool `json:"TruongNhanSu"`
}

func (DM_PhongBan_NhanVien) TableName() string {
	return "DM_PhongBan_NhanVien"
}
