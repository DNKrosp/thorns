package config

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "thorns/docs"
	"thorns/health"
	"thorns/health/http"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc health.UseCase) {
	h := http.NewHandler(uc)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/health", h.HealthAction)
}
