package cashRegister

import (
	"context"
	"testing"

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

	service := NewService(repositoryMock)
	basket, err := service.CreateBasket(context.Background())

	repositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, basketExpected, basket)
}

func TestService_GetBasketBasket_RepositoryError(t *testing.T) {
	repositoryMock := new(storagemocks.Repository)
	repositoryMock.On("FindBasketByID", mock.Anything, mock.Anything).Return(models.Basket{}, models.ErrBasketNotFound)

	service := NewService(repositoryMock)
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

	service := NewService(repositoryMock)
	basket, err := service.GetBasket(context.Background(), "99999")

	repositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, basketExpected, basket)
}

func TestService_Remove_Basket_Success(t *testing.T) {
	repositoryMock := new(storagemocks.Repository)
	repositoryMock.On("RemoveBasket", mock.Anything, mock.Anything).Return(nil)

	service := NewService(repositoryMock)
	err := service.RemoveBasket(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256")

	repositoryMock.AssertExpectations(t)
	assert.NoError(t, err)
}

func TestService_Remove_Basket_Unsuccess(t *testing.T) {
	repositoryMock := new(storagemocks.Repository)
	repositoryMock.On("RemoveBasket", mock.Anything, mock.Anything).Return(models.ErrBasketNotFound)

	service := NewService(repositoryMock)
	err := service.RemoveBasket(context.Background(), "4200f350-4fa5-11ec-a386-1e003b1e5256")

	repositoryMock.AssertExpectations(t)
	assert.Equal(t, models.ErrBasketNotFound, err)
}
