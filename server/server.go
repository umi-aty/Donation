package server

import (
	"yesiamdonation/config"
	"yesiamdonation/controllers"
	"yesiamdonation/util"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.ConfigDatabase()
	authController controllers.AuthController = controllers.NewAuthController(util.ProvideUserAuthService(), util.ProvideUserJwtService())
)

func RegisterRoute(r *gin.Engine) {
	defer config.CloseDBConnection(db)

	r.POST("/register", authController.Register)

	r.Run()

}
