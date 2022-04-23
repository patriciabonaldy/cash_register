package storage

import (
	"context"

	"github.com/patriciabonaldy/cash_register/internal/models"
)

// Repository defines the expected behaviour from a storage.
type Repository interface {
	FindBasketByID(ctx context.Context, id string) (models.Basket, error)
	CreateBasket(ctx context.Context, id string) (models.Basket, error)
	GetItem(ctx context.Context, basketID string, productCode string) (models.Item, error)
	UpdateBasket(ctx context.Context, basketID models.Basket) (models.Basket, error)
	RemoveProduct(ctx context.Context, basketID, productCode string) (models.Basket, error)
	RemoveBasket(ctx context.Context, id string) error
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=storagemocks --name=Repository
