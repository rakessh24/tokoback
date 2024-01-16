package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"toko/api/auth/login"
	"toko/api/auth/logout"
	"toko/api/auth/register"
	"toko/api/auth/user"
)

func Auth(r *gin.Engine, db *gorm.DB) {
	login.Login(r, db)
	register.Register(r, db)
	logout.Logout(r)
	user.User(r, db)
}
