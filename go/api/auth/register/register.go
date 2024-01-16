package register

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"toko/database"
	"toko/hash"
)

type RegisterRequest struct {
	// IdUser   int    `form:"id_user"`
	Username string `form:"username" binding:"required"`
	Nama     string `form:"nama_user" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Alamat   string `form:"alamat" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Register(r *gin.Engine, db *gorm.DB) {
	r.POST("/register", func(c *gin.Context) {
		var registerRequest RegisterRequest

		if err := c.ShouldBind(&registerRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Permintaan tidak valid!"})
			return
		}

		// Gunakan fungsi hash untuk menyimpan kata sandi
		hashedPassword, err := hash.HashPassword(registerRequest.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password!"})
			return
		}
		newUser := database.User{
			Username: registerRequest.Username,
			Nama:     registerRequest.Nama,
			Email:    registerRequest.Email,
			Alamat:   registerRequest.Alamat,
			Password: hashedPassword,
		}

		if err := db.Create(&newUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat akun!"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Registrasi Berhasil!"})
	})
}
