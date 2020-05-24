package data_user

type DmPhanquyenReport struct {
	DM_PhanQuyen_ReportId int `gorm:"column:DM_PhanQuyen_ReportId; primary_key;AUTO_INCREMENT" json:"dm_phan_quyen_report_id"`
	DM_PhanQuyenID        int `gorm:"column:DM_PhanQuyenID" json:"dm_phan_quyen_id"`
	DM_ReportId           int `gorm:"column:DM_ReportId" json:"dm_report_id"`
}

func (DmPhanquyenReport) TableName() string {
	return "DM_PhanQuyen_Report"
}
