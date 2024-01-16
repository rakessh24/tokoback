package bata

import (
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"toko/database"
)

func Bata(r *gin.Engine, db *gorm.DB) {
	r.GET("/api/bata", func(c *gin.Context) {
		var data []database.Bata

		if err := db.Find(&data).Error; err != nil {
			fmt.Println("Error fetching data:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil data!"})
			return
		}

		for i := range data {
			encodedImage := base64.StdEncoding.EncodeToString(data[i].Foto)
			data[i].EncodedFoto = encodedImage
		}

		c.JSON(http.StatusOK, data)
	})
}
