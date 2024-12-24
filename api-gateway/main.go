package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/visaramadhan/project-golang_microservice_tiicket-online/api-gateway/middleware"
)

func main() {
	router := gin.Default()

	// Product Routes - Authentication Required
	categoryRoutes := router.Group("/category")
	categoryRoutes.Use(middleware.AuthMiddleware())
	categoryRoutes.Any("/*proxyPath", reverseProxy("http://localhost:8087", "/product"))

	router.Run(":8087")
}

func reverseProxy(target, prefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxyPath := strings.TrimPrefix(c.Param("proxyPath"), prefix)
		targetURL := target + proxyPath
		http.Redirect(c.Writer, c.Request, targetURL, http.StatusTemporaryRedirect)
	}
}
