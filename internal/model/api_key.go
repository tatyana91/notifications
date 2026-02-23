package model

import "time"

type APIKey struct {
	KeyHash     string    `gorm:"primaryKey;size:64"`
	ServiceName string    `gorm:"size:100;not null"`
	IsActive    bool      `gorm:"default:true"`
	CreatedAt   time.Time `gorm:"default:now()"`
}
