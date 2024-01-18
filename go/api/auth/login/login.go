package login

import (
	"net/http"
	"toko/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Identifier string `form:"identifier" binding:"required"`
	Password   string `form:"password" binding:"required"`
}

func Login(r *gin.Engine, db *gorm.DB) {
	r.POST("/login", func(c *gin.Context) {
		if isUserLoggedIn(c) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Permintaan tidak valid!"})
			return
		}

		var loginRequest LoginRequest
		if err := c.ShouldBind(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Permintaan tidak valid!"})
			return
		}

		var user database.User
		// 	if err := db.Where("username = ?", loginRequest.Identifier).First(&user).Error; err != nil {
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "Akun tidak terdaftar!"})
		// return
		if err := db.Where("username = ?", loginRequest.Identifier).First(&user).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Akun tidak terdaftar!"})
			return
		}

		if !verifyPassword(loginRequest.Password, user.Password) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Password anda salah!"})
			return
		}

		loginHandler(c, db, user)
	})

}
