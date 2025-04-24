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

// HealthAction @Summary Метод статуса работы сервиса
// @Tags health
// @Description Проверка работы сервиса и его компонент
// @ID health
// @Accept json
// @Produce json
// @Success 200 {integer} integer 1
// @Router /health [get]
func (h Handler) HealthAction(context *gin.Context) {
	context.JSON(
		http.StatusOK,
		h.useCase.Health(context),
	)
}
