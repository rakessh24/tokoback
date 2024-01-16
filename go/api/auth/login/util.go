package login

import (
	"fmt"
	"toko/session"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func isUserLoggedIn(c *gin.Context) bool {
	// Hapus sesi yang ada sebelumnya (jika ada)
	session.ClearSession(c.Writer, c.Request)

	if _, ok := session.GetSession(c.Request)["user"]; ok {
		return true
	}
	return false
}

func checkIdentifier(identifier string) string {
	// Hapus bagian yang menghash NIK
	return identifier
}

func verifyPassword(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func convertMap(data map[interface{}]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range data {
		result[fmt.Sprintf("%v", key)] = value
	}
	return result
}
