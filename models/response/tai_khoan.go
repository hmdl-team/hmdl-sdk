package response

import (
	. "hmdl-user-service/models/data_user"
)

type ResTaiKhoanToken struct {
	Token        string `json:"token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

type ResTaiKhoan struct {
	DM_NhanVienId  int           `json:"dm_nhan_vien_id,omitempty"`
	NhanVienSuDung *NhanVien     `json:"nhan_vien_su_dung,omitempty"`
	PhanQuyenId    *int          `json:"phan_quyen_id,omitempty"`
	DonVi          string        `json:"don_vi,omitempty"`
	DM_PhanQuyenID *int          `json:"dm_phan_quyen_id,omitempty"`
	DM_PhanQuyen   *DM_PhanQuyen `json:"dm_phan_quyen,omitempty"`
}
