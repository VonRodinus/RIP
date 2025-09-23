package models

import _ "gorm.io/gorm"

// Artifact представляет артефакт (услугу)
type Artifact struct {
	ID          string `gorm:"primaryKey"`
	Name        string
	Description string
	Status      string `gorm:"default:active"`
	ImageURL    string `gorm:"type:varchar(255)"`
	TPQ         int
	StartDate   int
	EndDate     int
	Epoch       string
}
