package main

import (
	"github.com/gin-gonic/gin"
	"github.com/syahjamal/mccs_server_mysql/config"
	"github.com/syahjamal/mccs_server_mysql/controller"
	"gorm.io/gorm"
)

//variable global untuk controller
var (
	db             *gorm.DB                  = config.SetupDatabaseConn()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConn(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	r.Run()
}
