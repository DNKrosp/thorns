package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"thorns/health"
)

type Handler struct {
	useCase health.UseCase
}

func NewHandler(useCase health.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

func (h Handler) HealthAction(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		h.useCase.Health(context),
	)
}
