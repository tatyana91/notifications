package main

import (
	"notifications/internal/config"
	"notifications/internal/handler"
	"notifications/internal/middleware"
	"notifications/pkg/db"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func setupRouter() *gin.Engine {
	r := gin.Default()

	//r.Use(middleware.Logger())
	r.Use(middleware.APIKeyAuth(DB))

	v1 := r.Group("/api/v1")
	{
		notifications := handler.NewNotificationHandler(DB)

		v1.POST("/notifications", notifications.Create)
		//v1.PATCH("/notifications/:id", notifications.Update)
		//v1.GET("/notifications", notifications.List)
	}

	return r
}

func main() {
	cfg := config.Load()

	var err error
	DB, err = db.Connect(cfg.Database)
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}

	sqlDB, _ := DB.DB()
	if err := sqlDB.Ping(); err != nil {
		panic("database ping failed: " + err.Error())
	}

	r := setupRouter()
	r.Run(":8080")
}
