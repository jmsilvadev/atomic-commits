package services

import (
	"context"

	"github.com/jmsilvadev/atomic-commits/internal/entities"
)

type Service interface {
	CreateTransaction(ctx context.Context, data *entities.Transaction) (*entities.Transaction, error)
}
