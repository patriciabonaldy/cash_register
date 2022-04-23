package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/patriciabonaldy/cash_register/internal/cashRegister"
	"github.com/patriciabonaldy/cash_register/internal/platform/server/handler"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
	service  cashRegister.Service
}

func New(port uint, service cashRegister.Service) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf(":%d", port),
		service:  service,
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.GET("/health", handler.CheckHandler())
	basket := s.engine.Group("/baskets")
	{
		basket.POST("", handler.CreateBasketHandler(s.service))
		basket.GET("/:id", handler.GetBasketHandler(s.service))
		basket.DELETE(":id", handler.RemoveBasketHandler(s.service))
		basket.POST("/:id/checkout", handler.CheckoutBasketHandler(s.service))
		basket.POST("/:id/products/:code", handler.AddProductHandler(s.service))
		basket.DELETE("/:id/products/:code", handler.RemoveProductHandler(s.service))
	}
}
