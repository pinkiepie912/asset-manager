package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/health", func(context *gin.Context) {
		context.String(http.StatusOK, "OK!")
	})

	// Add middlewares
	return router
}
