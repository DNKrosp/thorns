package health

import (
	"context"
	"thorns/models"
)

type UseCase interface {
	Health(ctx context.Context) *models.Health
}
