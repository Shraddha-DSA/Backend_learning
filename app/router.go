package app

import (
	"go_assignment/handlers"
	"go_assignment/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", handlers.HealthCheck)

	r.GET("/users", handlers.GetUsers)
	r.POST("/users", handlers.CreateUser)
	r.PUT("/users/:id", handlers.UpdateUser)
	r.DELETE("/users/:id", handlers.DeleteUser)

	r.GET("/search", handlers.Search)

	r.GET("/users/:id/orders/:orderId", handlers.RouteParams)
	r.GET("/context", handlers.ContextData)

	r.GET(
		"/protected",
		middleware.ApiKeyAuth(),
		handlers.ProtectedRoute,
	)
	r.POST("/users/json", handlers.CreateUserJSON)

	r.GET("/profiles/:id", handlers.UserProfile)

	r.POST("/login", handlers.Login)
	r.GET("/profile", middleware.TokenAuth(), handlers.Profile)

	r.GET("/profile", middleware.JWTAuth(), handlers.Profile)

	r.GET("/context", handlers.ContextExample)
	r.GET("/health", handlers.Health)

}
