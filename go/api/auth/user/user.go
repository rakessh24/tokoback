package user

import (
	"net/http"
	"toko/session"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func User(r *gin.Engine, db *gorm.DB) {
	r.GET("/api/userData/byID", session.VerifyToken(), func(c *gin.Context) {
		// Hapus sesi yang ada sebelumnya (jika ada)
		session.ClearSession(c.Writer, c.Request)

		sessionData := session.GetSession(c.Request)

		if sessionData == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Sesi tidak ditemukan"})
			return
		}

		user, ok := sessionData["user"].(map[string]interface{})
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Format sesi tidak valid"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"user": user})
	})
}
