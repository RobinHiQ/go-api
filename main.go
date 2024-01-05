package main

import (
	"log"
	"os"

	api "github.com/RobinHiQ/go-api/api"
	"github.com/RobinHiQ/go-api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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

	// Use gin's routing functions to add routes
	r.GET("/job-description", gin.HandlerFunc(api.GetJobDescription))

	r.Run(":" + port)

	log.Println("listening on", port)
}
