package data_user

type DM_DuAn struct {
	DM_DuAnId int    `gorm:"column:DM_DuAnId" json:"dm_du_an_id"`
	TenDuAn   string `gorm:"column:TenDuAn" json:"ten_du_an"`
}

func (DM_DuAn) TableName() string {
	return "DM_DuAn"
}
