package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"helloGinAndGorm/config"
	"helloGinAndGorm/controller"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
	authController controller.AuthController = controller.NewAuthController()
)

func main() {
	defer config.CloseDatabaseConnection(db)
	router := gin.Default()

	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	//router.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "Hello World",
	//	})
	//})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}