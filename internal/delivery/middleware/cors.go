package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Origin"}
	corsConfig.AllowOrigins = []string{"http://95.163.212.121:8000", "http://95.163.212.121:8080", "http://95.163.212.121:80"}

	return cors.New(corsConfig)
}
