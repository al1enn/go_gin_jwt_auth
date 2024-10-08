package routes

import (
	"auth/pkg/handlers"
	"auth/pkg/middlewares"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(server *gin.Engine) {
	// Группа для публичных маршрутов
	public := server.Group("/auth")
	{
		public.POST("/regsiter", handlers.Register)
		public.POST("/login", handlers.Login)
		public.POST("/logout", middlewares.AuthMiddleware, handlers.Logout)
	}

	// Swagger route
	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Private routes
	private := server.Group("/api")
	{
		private.GET("/profile", middlewares.AuthMiddleware, handlers.Profile)
	}
}
