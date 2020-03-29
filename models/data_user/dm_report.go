package data_user

type DM_Report struct {
	DM_ReportId int    `gorm:"column:DM_ReportId; primary_key;AUTO_INCREMENT" json:"dm_report_id"`
	Name        string `gorm:"column:Name" json:"name"`
	ParentId    int    `gorm:"column:ParentId" json:"parent_id"`
	ReporCode   string `gorm:"column:ReporCode" json:"repor_code"`
	TamNgung    bool   `gorm:"column:TamNgung,default:0" json:"tam_ngung"`
}
