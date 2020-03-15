package models

import "hmdl-user-service/models/core"

type Service struct {
	core.Model
	TenService    string          `json:"ten_service"`
	MoTa          string          `json:"mo_ta"`
	TinhTrang     bool            ` json:"tinh_trang"`
	ServiceEmails []*ServiceEmail `gorm:"foreignkey:ServiceId" json:"service_id"`
}
