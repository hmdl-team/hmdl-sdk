package data_user

type DM_PhanQuyen struct {
	DM_PhanQuyenID int    `gorm:"column:DM_PhanQuyenID; primary_key;AUTO_INCREMENT" json:"dm_phan_quyen_id"`
	CapDo          string `gorm:"column:CapDo;" json:"cap_do"`
	TinhTrang      bool   `gorm:"column:TinhTrang;" json:"tinh_trang"`
}

func (DM_PhanQuyen) TableName() string {
	return "DM_PhanQuyen"
}
