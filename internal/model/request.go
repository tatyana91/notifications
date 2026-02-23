package model

type CreateNotificationRequest struct {
	UserID        int64          `json:"user_id" binding:"required"`
	Title         string         `json:"title" binding:"required,max=255"`
	Body          string         `json:"body" binding:"required"`
	SourceService string         `json:"source_service" binding:"required,max=100"`
	AuthorID      string         `json:"author_id" binding:"required,max=100"`
	Payload       map[string]any `json:"payload,omitempty"`
}

type UpdateNotificationRequest struct {
	Status int `json:"status" binding:"required,oneof=0 1"`
}

type GetNotificationsRequest struct {
	Limit    int    `form:"limit,default=20" binding:"min=1,max=100"`
	Offset   int    `form:"offset,default=0"`
	DateFrom string `form:"date_from"`
	DateTo   string `form:"date_to"`
	AuthorID string `form:"author_id"`
	Service  string `form:"service"`
	Status   *int   `form:"status"`
}
