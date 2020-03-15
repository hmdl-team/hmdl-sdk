package data_user

type DmChucVu struct {
	DmChucVuId       int    `gorm:"column:DM_ChucVuId; primary_key;AUTO_INCREMENT" json:"dm_chuc_danh_id"`
	TenChucVu        string `gorm:"column:TenChucVu;" json:"ten_chuc_vu"`
	TenChucVuVietTat string `gorm:"column:TenChucVuVietTat;" json:"ten_chuc_vu_viet_tat"`
	GhiChu           string `gorm:"column:GhiChu;" json:"ghi_chu"`
}

func (DmChucVu) TableName() string {
	return "DM_ChucVu"
}
