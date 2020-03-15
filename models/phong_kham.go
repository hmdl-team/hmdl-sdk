package models

import "hmdl-user-service/models/core"

type PhongKham struct {
	core.Model
	MaPhongKham  string `json:"ma_phong_kham"`
	TenPhongKham string `json:"ten_phong_kham"`
	DiaChi       string `json:"dia_chi"`
	SoDienThoai  string `json:"so_dien_thoai"`
}
