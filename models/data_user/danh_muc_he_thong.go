package data_user

import "hmdl-user-service/models/core"

type DanhMucHeThong struct {
	core.Model
	DanhMucCode   int    `gorm:"column:DanhMucCode; default:null" json:"danh_muc_code,omitempty"`
	LoaiDanhMuc   string `gorm:"column:LoaiDanhMuc;default:null" json:"loai_danh_muc,omitempty"`
	NhomDanhMucId int    `gorm:"column:NhomDanhMucId;default:null" json:"nhom_danh_muc_id,omitempty"`
	TenDanhMuc    string `gorm:"column:TenDanhMuc;default:null" json:"ten_danh_muc,omitempty"`
	VietTat       string `gorm:"column:VietTat;default:null" json:"viet_tat,omitempty"`
	GhiChu        string `gorm:"column:GhiChu;default:null" json:"ghi_chu,omitempty"`
}

func (DanhMucHeThong) TableName() string {
	return "DanhMucHeThong"
}
