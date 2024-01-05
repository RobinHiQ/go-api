package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"github.com/RobinHiQ/go-api/docs"
)

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	docs.SwaggerInfo.Title = "API Documentation for Job Description Generator"
	docs.SwaggerInfo.Description = "Job description generator API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "https://available-work.fly.dev/"
	docs.SwaggerInfo.BasePath = "/v2"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	r := gin.New()

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
