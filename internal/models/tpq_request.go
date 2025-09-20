package models

import (
	"time"
)

type TPQRequestItem struct {
	RequestID  string   `gorm:"primaryKey"`
	ArtifactID string   `gorm:"primaryKey"`
	Comment    string   // Другие поля, если есть: Quantity int, ItemOrder int, Main bool
	Artifact   Artifact `gorm:"foreignKey:ArtifactID"`
}

type TPQRequest struct {
	ID          string `gorm:"primaryKey"`
	Status      string
	CreatedAt   time.Time
	CreatorID   uint
	FormedAt    *time.Time
	CompletedAt *time.Time
	ModeratorID *uint
	Excavation  string
	Result      string
	TPQItems    []TPQRequestItem `gorm:"foreignKey:RequestID"`
}
