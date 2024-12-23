package main

import (
	"ticket-online/middleware"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Product Routes - Authentication Required
	categoryRoutes := router.Group("/category")
	categoryRoutes.Use(middleware.AuthMiddleware())
	categoryRoutes.Any("/*proxyPath", reverseProxy("http://localhost:8081", "/product"))

	router.Run(":8080")
}

func reverseProxy(target, prefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxyPath := strings.TrimPrefix(c.Param("proxyPath"), prefix)
		targetURL := target + proxyPath
		http.Redirect(c.Writer, c.Request, targetURL, http.StatusTemporaryRedirect)
	}
}