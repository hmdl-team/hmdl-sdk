package data_user

import "time"

type DM_MenuWeb struct {
	Id         int           `gorm:"column:Id; primary_key;AUTO_INCREMENT" json:"id"`
	MenuName   string        `gorm:"column:MenuName;" json:"menu_name"`
	ParentId   *int          `gorm:"column:ParentId;" json:"parent_id"`
	Children   []*DM_MenuWeb `gorm:"foreignkey:ParentId;" json:"children"`
	Path       string        `gorm:"column:Path;" json:"path"`
	Icon       string        `gorm:"column:Icon;" json:"icon"`
	Enable     bool          `gorm:"column:Enable;" json:"enable"`
	GroupId    *int          `gorm:"column:GroupId;" json:"group_id"`
	OrderBy    *int          `gorm:"column:OrderBy;" json:"order_by"`
	ImageLink  *int          `gorm:"column:ImageLink;" json:"image_link"`
	Mota       string        `gorm:"column:Mota;" json:"mota"`
	Ismobile   bool          `gorm:"column:Ismobile;" json:"ismobile"`
	MobileLink string        `gorm:"column:MobileLink;" json:"mobile_link"`
	CreatedAt  *time.Time    `gorm:"column:CreatedAt;" sql:"DEFAULT:current_timestamp" json:"-"`
	UpdatedAt  *time.Time `gorm:"column:UpdatedAt;" json:"-"`
}

func (DM_MenuWeb) TableName() string {
	return "DM_MenuWeb"
}
