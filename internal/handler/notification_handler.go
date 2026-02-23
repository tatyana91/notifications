package handler

import (
	"notifications/internal/model"
	"notifications/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type NotificationHandler struct {
	Service *service.NotificationService
}

func NewNotificationHandler(db *gorm.DB) *NotificationHandler {
	return &NotificationHandler{
		Service: service.NewNotificationService(db),
	}
}

func (h *NotificationHandler) Create(c *gin.Context) {
	var req model.CreateNotificationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	notification, err := h.Service.Create(c.Request.Context(), req)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "failed to create notification"})
		return
	}

	c.JSON(201, notification)
}
