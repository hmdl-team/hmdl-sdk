package request

import "hmdl-user-service/models/data_user"

type PhanQuyenBaoCaoReq struct {
	DM_PhanQuyenID int                   `json:"dm_phan_quyen_id"`
	DanhSachReport []data_user.DM_Report `json:"danh_sach_report"`
}

type PhanQuyenMenuReq struct {
	DM_PhanQuyenID int                    `json:"dm_phan_quyen_id"`
	DanhSachMenu   []data_user.DM_MenuWeb `json:"danh_sach_menu"`
}
