package cashRegister

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/patriciabonaldy/cash_register/internal/models"
	"github.com/patriciabonaldy/cash_register/internal/platform/storage/storagemocks"
)

func TestService_CreateBasketBasket_Success(t *testing.T) {
	repositoryMock := new(storagemocks.Repository)
	basketExpected := models.Basket{
		Code:  "99999",
		Total: 0,
	}
	repositoryMock.On("CreateBasket", mock.Anything, mock.Anything).Return(basketExpected, nil)

	service := NewService(nil, repositoryMock)
	basket, err := service.CreateBasket(context.Background())

	repositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, basketExpected, basket)
}

func TestService_GetBasketBasket_RepositoryError(t *testing.T) {
	repositoryMock := new(storagemocks.Repository)
	repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(models.Basket{}, models.ErrBasketNotFound)

	service := NewService(nil, repositoryMock)
	_, err := service.GetBasket(context.Background(), "99999")

	repositoryMock.AssertExpectations(t)
	assert.Error(t, err)
}

func TestService_GetBasketBasket_Success(t *testing.T) {
	repositoryMock := new(storagemocks.Repository)
	basketExpected := models.Basket{
		Code:  "99999",
		Total: 0,
	}
	repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(basketExpected, nil)

	service := NewService(nil, repositoryMock)
	basket, err := service.GetBasket(context.Background(), "99999")

	repositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, basketExpected, basket)
}

func TestService_Remove_Basket_Success(t *testing.T) {
	repositoryMock := new(storagemocks.Repository)
	repositoryMock.On("RemoveBasket", mock.Anything, mock.Anything).Return(nil)

	service := NewService(nil, repositoryMock)
	err := service.RemoveBasket(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256")

	repositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestService_Remove_Basket_Unsuccess(t *testing.T) {
	repositoryMock := new(storagemocks.Repository)
	repositoryMock.On("RemoveBasket", mock.Anything, mock.Anything).Return(models.ErrBasketNotFound)

	service := NewService(nil, repositoryMock)
	err := service.RemoveBasket(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256")

	repositoryMock.AssertExpectations(t)
	assert.Equal(t, models.ErrBasketNotFound, err)
}

func TestService_AddProduct_First_Time_Success(t *testing.T) {
	repositoryMock := new(storagemocks.Repository)
	itemMock := models.Item{
		Product: models.Product{
			Code:  "TSHIRT",
			Name:  "Summer T-Shirt",
			Price: 20,
		},
	}

	itemMock.Quantity = 1
	itemMock.Total = 20
	basketExpected := models.Basket{
		Code: "4200f350-4fa5-11ec-a386-1e003b1e5256",
		Items: map[string]models.Item{
			"TSHIRT": itemMock,
		},
		Total: 20,
	}
	basketmock := models.Basket{
		Code:  "4200f350-4fa5-11ec-a386-1e003b1e5256",
		Items: make(map[string]models.Item),
		Total: 0,
	}

	repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(basketmock, nil).Once()
	repositoryMock.On("GetItem", mock.Anything, mock.Anything, mock.Anything).Return(models.Item{}, models.ErrItemNotFound).Once()
	repositoryMock.On("CreateItem", mock.Anything, mock.Anything, mock.Anything).Return(itemMock, nil)
	repositoryMock.On("UpdateBasket", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(basketExpected, nil).Once()

	service := NewService(nil, repositoryMock)
	basket, err := service.AddProduct(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256", "TSHIRT")

	assert.NoError(t, err)
	assert.Equal(t, basketExpected, basket)
}

func TestService_AddProduct_Success(t *testing.T) {
	repositoryMock := new(storagemocks.Repository)
	itemMock := models.Item{
		Product: models.Product{
			Code:  "TSHIRT",
			Name:  "Summer T-Shirt",
			Price: 20,
		},
		Quantity: 2,
		Total:    40,
	}
	basketExpected := models.Basket{
		Code: "4200f350-4fa5-11ec-a386-1e003b1e5256",
		Items: map[string]models.Item{
			"TSHIRT": itemMock,
			"PANTS": {
				Product: models.Product{
					Code:  "PANTS",
					Name:  "Summer Pants",
					Price: 7.5,
				},
				Quantity: 1,
				Total:    7.5,
			},
		},
		Total: 20,
	}
	repositoryMock.On("GetItem", mock.Anything, mock.Anything, mock.Anything).Return(itemMock, nil)
	repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(basketExpected, nil)

	basketExpected = models.Basket{
		Code: "4200f350-4fa5-11ec-a386-1e003b1e5256",
		Items: map[string]models.Item{
			"TSHIRT": {
				Product: models.Product{
					Code:  "TSHIRT",
					Name:  "Summer T-Shirt",
					Price: 20,
				},
				Quantity: 3,
				Total:    45,
			},
			"PANTS": {
				Product: models.Product{
					Code:  "PANTS",
					Name:  "Summer Pants",
					Price: 7.5,
				},
				Quantity: 1,
				Total:    7.5,
			},
		},
		Total: 52.5,
	}
	repositoryMock.On("UpdateBasket", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(basketExpected, nil)

	service := NewService(nil, repositoryMock)
	basket, err := service.AddProduct(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256", "TSHIRT")
	assert.NoError(t, err)
	assert.Equal(t, basketExpected, basket)
}

func TestService_Remove_Product_Success(t *testing.T) {
	basketExpected := models.Basket{
		Code:  "4200f350-4fa5-11ec-a386-1e003b1e5256",
		Items: make(map[string]models.Item),
		Total: 0,
	}
	repositoryMock := new(storagemocks.Repository)
	repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(basketExpected, nil)
	repositoryMock.On("RemoveProduct", mock.Anything, mock.Anything, mock.Anything).Return(basketExpected, nil).Once()

	service := NewService(nil, repositoryMock)
	basket, err := service.RemoveProduct(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256", "TSHIRT")
	repositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, basketExpected, basket)
}

func TestService_Remove_Product_UnSuccess(t *testing.T) {
	repositoryMock := new(storagemocks.Repository)
	repositoryMock.On("RemoveProduct", mock.Anything, mock.Anything, mock.Anything).Return(models.Basket{}, models.ErrItemNotFound)

	service := NewService(nil, repositoryMock)
	_, err := service.RemoveProduct(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256", "DRESS")
	assert.Equal(t, models.ErrProductNotFound, err)

	// basket is closed and we try remove one product
	basketExpected := models.Basket{
		Code:  "4200f350-4fa5-11ec-a386-1e003b1e5256",
		Items: make(map[string]models.Item),
		Total: 0,
		Close: true,
	}
	repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(basketExpected, nil)
	repositoryMock.On("RemoveProduct", mock.Anything, mock.Anything, mock.Anything).Return(basketExpected, nil)
	_, err = service.RemoveProduct(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256", "PANTS")
	assert.EqualError(t, err, "basket is closed")
}

func TestService_CheckoutBasket(t *testing.T) {
	basketMock := models.Basket{
		Code: "4200f350-4fa5-11ec-a386-1e003b1e5256",
		Items: map[string]models.Item{
			"VOUCHER": {
				Product: models.Product{
					Code:  "VOUCHER",
					Name:  "Gift Card",
					Price: 5,
				},
				Quantity: 2,
				Total:    10,
			},
			"TSHIRT": {
				Product: models.Product{
					Code:  "TSHIRT",
					Name:  "Summer T-Shirt",
					Price: 20,
				},
				Quantity: 3,
				Total:    60,
			},
			"PANTS": {
				Product: models.Product{
					Code:  "PANTS",
					Name:  "Summer Pants",
					Price: 7.5,
				},
				Quantity: 1,
				Total:    7.5,
			},
		},
		Total: 20,
	}
	basketExpected := models.Basket{
		Code: "4200f350-4fa5-11ec-a386-1e003b1e5256",
		Items: map[string]models.Item{
			"VOUCHER": {
				Product: models.Product{
					Code:  "VOUCHER",
					Name:  "Gift Card",
					Price: 5,
				},
				Quantity: 2,
				Total:    10,
			},
			"TSHIRT": {
				Product: models.Product{
					Code:  "TSHIRT",
					Name:  "Summer T-Shirt",
					Price: 20,
				},
				Quantity: 3,
				Total:    45,
			},
			"PANTS": {
				Product: models.Product{
					Code:  "PANTS",
					Name:  "Summer Pants",
					Price: 7.5,
				},
				Quantity: 1,
				Total:    7.5,
			},
		},
		Total: 74.5,
		Close: true,
	}

	repositoryMock := new(storagemocks.Repository)
	repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(basketMock, nil)
	repositoryMock.On("UpdateBasket", mock.Anything, mock.Anything).Return(basketExpected, nil)

	service := NewService(RulesEngine, repositoryMock)
	basketID := "4200f350-4fa5-11ec-a386-1e003b1e5256"

	err := LoadRulesConfig()
	require.NoError(t, err)

	basket, err := service.CheckoutBasket(context.Background(), basketID)
	assert.NoError(t, err)
	assert.Equal(t, basketExpected, basket)
}
