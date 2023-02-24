package app

import (
	"depeche/internal/auth/delivery/http"
	httpserver "depeche/internal/server"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := initRouter()
	server := httpserver.NewServer(router)

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}

func initRouter() *gin.Engine {
	router := gin.Default()

	authHandler := http.NewAuthHandler()

	// вешаем авторизационную миддлвару на все api
	authMiddleware := http.NewAuthMiddleware(authHandler.Service)

	apiEndpointsGroup := router.Group("/api", authMiddleware)
	// тестовый эндпоинт
	apiEndpointsGroup.GET("/test", func(context *gin.Context) {
		context.Writer.WriteString("TEST")
	})

	// регистриурем урлы авторизации и регистрации
	http.BindAuthEndpoints(router, authHandler)

	return router
}
