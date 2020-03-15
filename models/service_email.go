package models

import "hmdl-user-service/models/core"

type ServiceEmail struct {
	core.Model
	ServiceId  int    `json:"service_id"`
	Email      string `json:"email"`
	NhanVienId int    `json:"nhan_vien_id"`
	TinhTrang  bool   `json:"tinh_trang"`
}
