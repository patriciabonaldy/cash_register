package bootstrap

import (
	"fmt"
	handler2 "github.com/patriciabonaldy/cash_register/api/cmd/bootstrap/handler"
	"github.com/patriciabonaldy/cash_register/api/cmd/docs"
	"log"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine
	handler  handler2.Handler
}

func New(port uint, handler handler2.Handler) Server {
	srv := Server{
		engine:   gin.New(),
		httpAddr: fmt.Sprintf(":%d", port),
		handler:  handler,
	}

	srv.engine.TrustedPlatform = gin.PlatformGoogleAppEngine

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

// Middleware is a gin.HandlerFunc that set CORS
func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (s *Server) registerRoutes() {
	s.engine.Use(Middleware())
	s.engine.GET("/health", handler2.CheckHandler())
	basket := s.engine.Group("/baskets")
	{
		basket.POST("", s.handler.CreateBasketHandler())
		basket.GET("/:id", s.handler.GetBasketHandler())
		basket.DELETE(":id", s.handler.RemoveBasketHandler())
		basket.POST("/:id/checkout", s.handler.CheckoutBasketHandler())
		basket.POST("/:id/products/:code", s.handler.AddProductHandler())
		basket.DELETE("/:id/products/:code", s.handler.RemoveProductHandler())
	}

	docs.SwaggerInfo.Title = "Swagger Documentation API"
	docs.SwaggerInfo.Description = "API Documentation."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "0.0.0.0:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}

	// use ginSwagger middleware to serve the API docs
	s.engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
