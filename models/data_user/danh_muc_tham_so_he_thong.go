package data_user

import "hmdl-user-service/models/core"

type DmThamSoHeThong struct {
	Id        int    `gorm:"column:Id;primary_key;AUTO_INCREMENT" json:"id"`
	MaThamSo  string `gorm:"column:MaThamSo;" json:"ma_tham_so, omitempty"`
	TenThamSo string `gorm:"column:TenThamSo;" json:"ten_tham_so, omitempty"`
	GiaTri    string `gorm:"column:GiaTri;" json:"gia_tri, omitempty"`
	core.Model2
}

func (DmThamSoHeThong) TableName() string {
	return "DM_ThamSoHeThong"
}
