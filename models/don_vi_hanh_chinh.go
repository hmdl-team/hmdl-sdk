package models

import "hmdl-user-service/models/core"

type DonViHanhChinh struct {
	core.Model
	MaDonVi     string
	TenDonVi    string
	Cap         int
	ParentId    *int
	TenTiengAnh string
	GhiChu      string
}
