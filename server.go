package main

import (
	"github.com/gin-gonic/gin"
	"github.com/syahjamal/mccs_server_mysql/config"
	"github.com/syahjamal/mccs_server_mysql/controller"
	"github.com/syahjamal/mccs_server_mysql/repository"
	"github.com/syahjamal/mccs_server_mysql/service"
	"gorm.io/gorm"
)

//variable global untuk controller
var (
	db             *gorm.DB                  = config.SetupDatabaseConn()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService     service.JWTService        = service.NewJWTService()
	authService    service.AuthService       = service.NewAuthService(userRepository)
	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConn(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	// r.Run(":5050")
	r.Run(":5050")
}
