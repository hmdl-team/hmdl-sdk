package core

import "time"

type Model struct {
	Id        int        `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	CreatedAt *time.Time `sql:"DEFAULT:current_timestamp" json:"-"`
	UpdatedAt *time.Time `json:"-"`
}

type Model2 struct {
	UserCreate string     `json:"-"`
	CreatedAt  *time.Time `sql:"DEFAULT:current_timestamp" json:"-"`
	UserUpdate string     `json:"-"`
	UpdatedAt  *time.Time `json:"-"`
}
