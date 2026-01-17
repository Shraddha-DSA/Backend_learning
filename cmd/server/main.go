package main

import (
	"go_assignment/app"
	"go_assignment/middleware"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	//r := gin.Default()
	r.Use(middleware.RequestID())
	r.Use(middleware.Logger())

	r.Use(gin.Recovery())
	r.Use(middleware.RateLimit(5, 10*time.Second))
	r.Use(middleware.Timeout(2 * time.Second))
	r.Use(middleware.Errorhandler())

	r.Use(middleware.CORSMiddleware())

	app.RegisterRoutes(r)
	r.Run()

	err := r.RunTLS(
		":8443",
		"cert.pem",
		"key.pem",
	)
	if err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}
}
