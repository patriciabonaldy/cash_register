package handler

import (
	"net/http"

	"github.com/patriciabonaldy/cash_register/internal/models"

	"github.com/gin-gonic/gin"

	"github.com/patriciabonaldy/cash_register/internal/cashRegister"
)

type Handler struct {
	service cashRegister.Service
}

func New(service cashRegister.Service) Handler {
	return Handler{
		service: service,
	}
}

// CreateBasketHandler create a new basket.
// return 201 if this could be created.
// Otherwise, it will return 500
// Create godoc
// @Summary      Create a new basket.
// @Description  return 201 if this could be created. Otherwise, it will return 500
// @Tags         basket
// @Accept       json
// @Produce      plain
// @Success      201  {string}  string         "success"
// @Failure      400  {string}  string         "bad Request"
// @Failure      500  {string}  string         "fail"
// @Router       /baskets [post]
func (h *Handler) CreateBasketHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		basket, err := h.service.CreateBasket(ctx)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}

		ctx.JSON(http.StatusCreated, toResponse(basket))
	}
}

// GetBasketHandler return basket.
// require a basket id and
// return 200 if this is ok.
// otherwise will return 400
// GetBasketHandler godoc
// @Summary      Show all products of basket
// @Description  requires a basket ID example:"0bfce8da-bdc9-11ec-b9f3-acde48001122"
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      500  {object}  Response
// @Router       /baskets/{id} [get]
func (h *Handler) GetBasketHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" || id == "0" {
			ctx.Status(http.StatusBadRequest)
			return
		}

		basket, err := h.service.GetBasket(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}

		ctx.JSON(http.StatusOK, toResponse(basket))
	}
}

// RemoveBasketHandler remove a basket.
// require a basket id.
// it will return 200 if this is ok.
// otherwise will return 400
// RemoveBasketHandler godoc
// @Summary      remove a basket
// @Description  requires a basket ID example:"0bfce8da-bdc9-11ec-b9f3-acde48001122"
// @Tags         basket
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "ID"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /baskets/{id} [DELETE]
func (h *Handler) RemoveBasketHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" || id == "0" {
			ctx.Status(http.StatusBadRequest)
			return
		}

		err := h.service.RemoveBasket(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}

		ctx.Status(http.StatusOK)
	}
}

// CheckoutBasketHandler close a basket.
// require a basket id.
// it will return 200 if this is ok.
// otherwise will return 400
// CheckoutBasketHandler godoc
// @Summary      close a basket
// @Description  requires a basket id, close of basket and will show details of order.
// @Tags         basket
// @Accept       json
// @Produce      plain
// @Param        id   path      string  true  "ID"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      500  {object}  Response
// @Router       /baskets/{id}/checkout [post]
func (h *Handler) CheckoutBasketHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" || id == "0" {
			ctx.Status(http.StatusBadRequest)
			return
		}

		basket, err := h.service.CheckoutBasket(ctx, id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, toResponse(basket))
	}
}

// AddProductHandler add a new product to basket.
// return 201 if this could be created.
// Otherwise, it will return 500
// AddProductHandler godoc
// @Summary      add a new product to basket.
// @Description  requires a basket id, and a product code. if product/code not exists then return "product does not exist"
// @Tags         basket
// @Accept       json
// @Produce      plain
// @Param        id     path      string  true  "ID"
// @Param        code   path      string  true  "CODE"
// @Success      200  {object}  Response
// @Failure      400  {object}  Response
// @Failure      500  {object}  Response
// @Router       /baskets/{id}/products/{code} [post]
func (h *Handler) AddProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}

		code := ctx.Param("code")
		if code == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}
		req := ProductRequest{
			BasketID:    id,
			ProductCode: code,
		}
		_, err := h.service.AddProduct(ctx, req.BasketID, req.ProductCode)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.Status(http.StatusCreated)
	}
}

// RemoveProductHandler remove a product inside a basket.
// require a basket id and product id.
// it will return 200 if this is ok.
// otherwise will return 400
// RemoveProductHandler godoc
// @Summary      remove a product in the basket.
// @Description  requires a basket id, and a product code. if product/code not exists then return "product does not exist"
// @Tags         basket
// @Accept       json
// @Produce      plain
// @Param        id     path      string  true  "ID"
// @Param        code   path      string  true  "CODE"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /baskets/{id}/products/{code} [delete]
func (h *Handler) RemoveProductHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		if id == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}

		code := ctx.Param("code")
		if code == "" {
			ctx.Status(http.StatusBadRequest)
			return
		}
		req := ProductRequest{
			BasketID:    id,
			ProductCode: code,
		}
		basket, err := h.service.RemoveProduct(ctx, req.BasketID, req.ProductCode)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}

		ctx.JSON(http.StatusOK, toResponse(basket))
	}
}

func toResponse(basket models.Basket) Response {
	resp := Response{
		ID:   basket.Code,
		Item: []Item{},
	}

	for _, v := range basket.Items {
		item := Item{
			Product: Product{
				Code:  v.Product.Code,
				Name:  v.Product.Name,
				Price: v.Product.Price,
			},
			Quantity: v.Quantity,
			Total:    v.Total,
		}
		resp.Item = append(resp.Item, item)
		resp.Total += item.Total
	}
	return resp
}
