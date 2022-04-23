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
	rulesEngine func(request models.Item) []Rule
	repository  storage.Repository
}

type ProductRequest struct {
	BasketID    string `json:"basket_id" binding:"required"`
	ProductCode string `json:"product_code" binding:"required"`
}

type BasketRequest struct {
	BasketID string `json:"basket_id" binding:"required"`
}

// NewService returns the default Service interface implementation.
func NewService(rules func(request models.Item) []Rule, repository storage.Repository) Service {
	return Service{rulesEngine: rules, repository: repository}
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

// RemoveBasket remove a basket.
// it will remove basket if this is ok.
// otherwise will return error
func (s Service) RemoveBasket(ctx context.Context, id string) error {
	err := s.repository.RemoveBasket(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// AddProduct add a new product into basket.
// require a basket id and product code
// it will return a basket if this is ok.
// otherwise will return  error
func (s Service) AddProduct(ctx context.Context, request ProductRequest) (models.Basket, error) {
	var basket, err = s.repository.FindBasketByID(ctx, request.BasketID)
	if err != nil {
		return models.Basket{}, err
	}

	if basket.Close {
		return models.Basket{}, models.ErrBasketIsClosed
	}

	item, err := s.repository.GetItem(ctx, request.BasketID, request.ProductCode)
	if err != nil {
		item, err = s.createItem(basket, request.ProductCode)
		if err != nil {
			return models.Basket{}, err
		}
	}

	item.Quantity++
	item.WithOutDiscount()
	code := item.Product.Code
	basket.Items[code] = item
	basket.CalculateTotal()

	basket, err = s.repository.UpdateBasket(ctx, basket)
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}

// RemoveProduct remove product inside basket.
// require a basket id and product code
// it will return a basket if this is ok.
// otherwise will return  error
func (s Service) RemoveProduct(ctx context.Context, request ProductRequest) (models.Basket, error) {
	_, ok := models.ProductMap[request.ProductCode]
	if !ok {
		return models.Basket{}, models.ErrProductNotFound
	}

	var basket, err = s.repository.FindBasketByID(ctx, request.BasketID)
	if err != nil {
		return models.Basket{}, err
	}

	if basket.Close {
		return models.Basket{}, models.ErrBasketIsClosed
	}

	basket, err = s.repository.RemoveProduct(ctx, request.BasketID, request.ProductCode)
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}

func (s Service) createItem(basket models.Basket, productCode string) (models.Item, error) {
	if basket.Close {
		return models.Item{}, models.ErrBasketIsClosed
	}

	product, ok := models.ProductMap[productCode]
	if !ok {
		return models.Item{}, models.ErrProductNotFound
	}

	item, ok := basket.Items[product.Code]
	if !ok {
		_item := models.Item{
			Product: product,
		}
		_item.WithOutDiscount()

		return _item, nil
	}

	return item, nil
}

// CheckoutBasket close a basket.
// require a basket id
// it will return a basket if this is ok.
// otherwise will return  error
func (s Service) CheckoutBasket(ctx context.Context, basketID string) (models.Basket, error) {
	var basket models.Basket
	basket, err := s.repository.FindBasketByID(ctx, basketID)
	if err != nil {
		return models.Basket{}, err
	}

	if basket.Close {
		return basket, nil
	}

	for _, item := range basket.Items {
		rulesItem := s.rulesEngine(item)
		for _, r := range rulesItem {
			item = r.fn(item, r)
		}

		basket.Items[item.Product.Code] = item
	}

	basket.CalculateTotal()
	basket.Close = true
	basket, err = s.repository.UpdateBasket(ctx, basket)
	if err != nil {
		return models.Basket{}, err
	}

	return basket, nil
}
