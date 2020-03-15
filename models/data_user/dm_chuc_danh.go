package data_user

type DmChucDanh struct {
	DmChucDanhId       int    `gorm:"column:DM_ChucDanhId; primary_key;AUTO_INCREMENT" json:"dm_chuc_danh_id"`
	TenChucDanh        string `gorm:"column:TenChucDanh;" json:"ten_chuc_danh"`
	TenChucDanhVietTat string `gorm:"column:TenChucDanhVietTat;" json:"TenChucDanhVietTat"`
	GhiChu             string `gorm:"column:GhiChu;" json:"GhiChu"`
}

func (DmChucDanh) TableName() string {
	return "DM_ChucDanh"
}
