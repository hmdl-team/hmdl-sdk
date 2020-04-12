package data_user

import "time"

type DM_PhongBan struct {
	DM_PhongBanId     int           `gorm:"column:DM_PhongBanID; primary_key;AUTO_INCREMENT" json:"dm_phong_ban_id,omitempty"`
	ParentKeyId       *int          `gorm:"column:ParentKeyId;" json:"parent_key_id"`
	Children          []DM_PhongBan `gorm:"foreignkey:ParentKeyId; column:Children" json:"children"`
	TenPhongBan       string        `gorm:"column:TenPhongBan;" json:"ten_phong_ban"`
	TruongKhoaPhongID int           `gorm:"column:TenPhongBan;" json:"truong_khoa_phong_id"`
	NgayTao           *time.Time    `gorm:"column:NgayTao" json:"ngay_tao"`
	TinhTrang         bool          `gorm:"column:TinhTrang" json:"tinh_trang"`
	Parent            bool          `gorm:"column:Parent" json:"parent"`
	TenTiengAnh       string        `gorm:"column:TenTiengAnh"json:"ten_tieng_anh"`
}

func (DM_PhongBan) TableName() string {
	return "DM_PhongBan"
}
