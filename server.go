package main

import (
	"os"

	"github.com/fabiosebastiano/graphql-server/http"
	"github.com/fabiosebastiano/graphql-server/middleware"
	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	server := gin.Default()

	server.Use(middleware.BasicAuth())

	server.GET("/", http.PlaygroundHandler())
	server.POST("/query", http.GraphQLHandler())
	server.Run(defaultPort)

}
