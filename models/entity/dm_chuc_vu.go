package entity

type DmChucVu struct {
	ChucVuId   int    `json:"chuc_vu_id"`
	TenChucVu  string `json:"ten_chuc_vu"`
	TenVietTat string `json:"ten_viet_tat"`
	GhiChu     string `json:"ghi_chu"`
}
