package cashRegister

import (
	"context"

	"github.com/google/uuid"

	"github.com/patriciabonaldy/cash_register/internal/models"
	"github.com/patriciabonaldy/cash_register/internal/platform/storage"
)

// Service is the default Service interface
// implementation returned byNewService.
type Service struct {
	repository storage.Repository
}

// NewService returns the default Service interface implementation.
func NewService(repository storage.Repository) Service {
	return Service{repository: repository}
}

// CreateBasket create a basket.
// it will return a new basket if this is ok.
// otherwise will return error
func (s Service) CreateBasket(ctx context.Context) (models.Basket, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return models.Basket{}, err
	}

	basket, err := s.repository.CreateBasket(ctx, id.String())
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}

// GetBasket return a basket.
// require a basket id
// it will return a basket if this is ok.
// otherwise will return  error
func (s Service) GetBasket(ctx context.Context, id string) (models.Basket, error) {
	basket, err := s.repository.FindBasketByID(ctx, id)
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}
