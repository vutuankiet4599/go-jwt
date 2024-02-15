package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vutuankiet4599/go-jwt/app/http/controllers"
	"github.com/vutuankiet4599/go-jwt/app/http/middlewares"
	"github.com/vutuankiet4599/go-jwt/app/repository"
	"github.com/vutuankiet4599/go-jwt/app/service"
	"github.com/vutuankiet4599/go-jwt/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB = config.SetupDatabaseConnection()
	userRepository repository.UserRepository = repository.NewUserRepository(db)
	jwtService service.JwtService = service.NewJwtService()
	authService service.AuthService = service.NewAuthService(userRepository)
	authController controllers.AuthController = controllers.NewAuthController(authService, jwtService)

	bookRepository = repository.NewBookRepository(db)
	bookService = service.NewBookService(bookRepository)
	bookController controllers.BookController = controllers.NewBookController(bookService)
)

func InitApiRouter() *gin.Engine {
	routes := gin.Default()

	authRoutes := routes.Group("/api/auth")
	{
		authRoutes.POST("/register", authController.Register)
		authRoutes.POST("/login", authController.Login)
		authRoutes.Use(middlewares.AuthorizeJwt(jwtService))
		{
			authRoutes.GET("/user", authController.User)
		}
	}

	bookRoutes := routes.Group("/api/book", middlewares.AuthorizeJwt(jwtService))
	{
		bookRoutes.GET("/", bookController.GetAll)
		bookRoutes.GET("/:id", bookController.GetOneById)
		bookRoutes.POST("/", bookController.Insert)
		bookRoutes.DELETE("/", bookController.DeleteAll)
		bookRoutes.Use(middlewares.CheckBookBelongToUser(db))
		{
			bookRoutes.PUT("/:id", bookController.Update)
			bookRoutes.DELETE("/:id", bookController.DeleteOneById)
		}
	}

	return routes
}
