package model

import "time"

type Notification struct {
	ID            uint   `gorm:"primaryKey"`
	UserID        int64  `gorm:"index"`
	Title         string `gorm:"size:255"`
	Body          string
	Status        int       `gorm:"default:0"`
	SourceService string    `gorm:"size:100"`
	AuthorID      string    `gorm:"size:100"`
	Payload       []byte    `gorm:"type:jsonb"`
	CreatedAt     time.Time `gorm:"index;not null"`
}
