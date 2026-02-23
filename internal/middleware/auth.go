package middleware

import (
	"crypto/sha256"
	"encoding/hex"
	"notifications/internal/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func APIKeyAuth(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")
		if key == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "missing API key"})
			return
		}

		hash := sha256.Sum256([]byte(key))
		keyHash := hex.EncodeToString(hash[:])

		var count int64
		db.Model(&model.APIKey{}).
			Where("key_hash = ? AND is_active = true", keyHash).
			Count(&count)

		if count == 0 {
			c.AbortWithStatusJSON(403, gin.H{"error": "invalid API key"})
			return
		}

		c.Next()
	}
}
