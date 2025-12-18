package models

type Projects struct {
	ID          uint `gorm:"primary key"`
	Title       string
	Description string
	URL         string
}
