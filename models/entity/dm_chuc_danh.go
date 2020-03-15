package entity

type DmChucdanh struct {
	Chucdanhid         int    `json:"chucdanhid"`
	TenChucDanh        string `json:"ten_chuc_danh"`
	TenChucDanhVietTat string `json:"ten_chuc_danh_viet_tat"`
	GhiChu             string `json:"ghi_chu"`
}
