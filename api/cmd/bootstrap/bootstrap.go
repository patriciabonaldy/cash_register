package bootstrap

import (
	"github.com/patriciabonaldy/cash_register/api/cmd/bootstrap/handler"
	"log"

	"github.com/patriciabonaldy/cash_register/internal/cashRegister"
	"github.com/patriciabonaldy/cash_register/internal/platform/storage/memory"
)

const (
	port = 8080
)

// Run application
func Run() error {
	err := cashRegister.LoadRulesConfig()
	if err != nil {
		log.Fatal(err)
	}

	repository := memory.NewRepository()
	service := cashRegister.NewService(cashRegister.RulesEngine, repository)
	handler := handler.New(service)
	srv := New(port, handler)
	return srv.Run()
}
