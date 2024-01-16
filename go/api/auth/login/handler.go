package login

import (
	"net/http"
	"toko/database"
	"toko/session"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func loginHandler(c *gin.Context, db *gorm.DB, user database.User) {
	// Hapus sesi yang ada sebelumnya (jika ada)
	session.ClearSession(c.Writer, c.Request)

	token, err := session.GenerateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menghasilkan token!"})
		return
	}

	tokenString, ok := token["token"].(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mendapatkan token!"})
		return
	}
	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	// c.SetCookie("token", tokenString, 3600, "/", "xhaeffer.me:11092", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login Berhasil!",
		"token":   token,
	})
}
