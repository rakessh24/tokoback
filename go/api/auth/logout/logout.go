package logout

import (
	"net/http"
	"toko/session"

	"github.com/gin-gonic/gin"
)

func logoutHandler(c *gin.Context) {
	if _, ok := session.GetSession(c.Request)["user"]; !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Permintaan tidak valid!"})
		return
	}

	// Hapus session dan token
	session.ClearSession(c.Writer, c.Request)
	session.ClearToken(c)

	c.JSON(http.StatusOK, gin.H{"message": "Logout Berhasil!"})
}

func Logout(r *gin.Engine) {
	r.GET("/logout", logoutHandler)
}
