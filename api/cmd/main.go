package main

import (
	"github.com/patriciabonaldy/cash_register/api/cmd/bootstrap"
	"log"
)

// @title API document title
// @version version(1.0)
// @description Description of specifications
// @Precautions when using termsOfService specifications

// @host 0.0.0.0:8080
// @BasePath /
func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal(err)
	}
}
