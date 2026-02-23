package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"notifications/internal/model"
	"time"

	"gorm.io/gorm"
)

type NotificationService struct {
	DB *gorm.DB
}

func NewNotificationService(db *gorm.DB) *NotificationService {
	return &NotificationService{DB: db}
}

func (s *NotificationService) Create(ctx context.Context, req model.CreateNotificationRequest) (*model.Notification, error) {
	defaultStatus := 0

	notification := &model.Notification{
		UserID:        req.UserID,
		Title:         req.Title,
		Body:          req.Body,
		Status:        defaultStatus,
		SourceService: req.SourceService,
		AuthorID:      req.AuthorID,
		CreatedAt:     time.Now(),
	}

	if len(req.Payload) > 0 {
		notification.Payload, _ = json.Marshal(req.Payload)
	}

	if err := s.DB.WithContext(ctx).Create(notification).Error; err != nil {
		log.Printf("DB error: %v", err)
		return nil, fmt.Errorf("failed to create notification: %w", err)
	}

	return notification, nil
}
