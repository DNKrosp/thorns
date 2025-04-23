package http

import (
	"github.com/gin-gonic/gin"
	"thorns/health"
)

func RegisterHTTPEndpoints(router *gin.Engine, uc health.UseCase) {
	h := NewHandler(uc)

	router.GET("/health", h.HealthAction)
}
