package data_user

type DM_PhanQuyenMenu struct {
	DM_PhanQuyenMenuId int          `gorm:"column:DM_PhanQuyenMenuId; primary_key;AUTO_INCREMENT" json:"phan_quyen_menu_id"`
	DM_PhanQuyenID     int          `gorm:"column:DM_PhanQuyenID;" json:"dm_phan_quyen_id"`
	DM_PhanQuyen       DM_PhanQuyen `gorm:"foreignkey:DM_PhanQuyenID" json:"dm_phan_quyen"`
	DM_MenuWebId       int          `gorm:"column:DM_MenuWebId;" json:"menu_web_id"`
	DM_MenuWeb         *DM_MenuWeb  `gorm:"foreignkey:DM_MenuWebId" json:"menu_web"`
}

func (DM_PhanQuyenMenu) TableName() string {
	return "DM_PhanQuyenMenu"
}
