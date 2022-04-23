package main

import (
	"log"

	"github.com/patriciabonaldy/cash_register/cmd/api/bootstrap"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
