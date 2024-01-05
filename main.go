package main

import (
	"log"
	"net/http"
	"os"

	api "github.com/RobinHiQ/go-api/api"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Helloworld PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

// @title           API Documentation for Job Description Generator
// @version         1.0
// @description     Job description generator API.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      https://available-work.fly.dev/
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/", func(c *gin.Context) {
		c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>Available jobs API</title>
			</head>
			<body>
				<h1>Welcome to the available jobs API written in Go!</h1>
				<p><a href="/swagger/index.html">View the API documentation</a></p>
			</body>
			</html>
		`))
	})

	r.GET("/helloworld", Helloworld)

	// Use gin's routing functions to add routes
	r.GET("/job-description", gin.HandlerFunc(api.GetJobDescription))

	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":" + port)

	log.Println("listening on", port)
}
