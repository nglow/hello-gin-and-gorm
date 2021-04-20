package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"helloGinAndGorm/config"
	"helloGinAndGorm/controller"
	"helloGinAndGorm/middleware"
	"helloGinAndGorm/repository"
	"helloGinAndGorm/service"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
	userRepository repository.UserRepository =  repository.NewUserRepository(db)
	bookRepository repository.BookRepository = repository.NewBookRepository(db)

	jwtService service.JwtService = service.NewJwtService()
	userService service.UserService = service.NewUserService(userRepository)
	authService service.AuthService = service.NewAuthService(userRepository)
	bookService service.BookService = service.NewBookService(bookRepository)

	authController controller.AuthController = controller.NewAuthController(authService, jwtService)
	userController controller.UserController = controller.NewUserController(userService, jwtService)
	bookController controller.BookController = controller.NewBookController(bookService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	router := gin.Default()

	authRoutes := router.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	userRoutes := router.Group("api/user", middleware.AuthorizeJwt(jwtService))
	{
		userRoutes.GET("/profile", userController.Profile)
		userRoutes.PUT("/profile", userController.Update)
	}

	bookRoutes := router.Group("api/books", middleware.AuthorizeJwt(jwtService))
	{
		bookRoutes.GET("/", bookController.All)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.GET("/:id", bookController.FindById)
		bookRoutes.PUT("/:id", bookController.Update)
		bookRoutes.DELETE("/:id", bookController.Delete)
	}
	//router.GET("/", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "Hello World",
	//	})
	//})
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}