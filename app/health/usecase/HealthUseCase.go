package usecase

import (
	"context"
	"github.com/spf13/viper"
	"thorns/models"
)

type HealthUseCase struct {
}

func NewHealthUseCase() *HealthUseCase {
	return &HealthUseCase{}
}

func (l HealthUseCase) Health(ctx context.Context) *models.Health {
	return &models.Health{
		Status:  true,
		Version: viper.GetString("version"),
	}
}
