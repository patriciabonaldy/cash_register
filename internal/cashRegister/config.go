package cashRegister

import (
	_ "embed"
	"fmt"

	"github.com/patriciabonaldy/cash_register/internal/models"

	"gopkg.in/yaml.v3"
)

//go:embed rules.yml
var data []byte

// Config represents the structure to store all about limit configuration.
type Config struct {
	Rules rules `yaml:"rules"`
}

type (
	ruleName string
	rules    map[ruleName]Rule
)

// Rule represents the structure to store the details of a rule by default.
type Rule struct {
	Name     ruleName `yaml:"name"`
	Desc     string   `yaml:"desc"`
	Product  string   `yaml:"product"`
	Quantity int      `yaml:"quantity"`
	NewPrice float64  `yaml:"newPrice,omitempty"`
	fn       func(item models.Item, rule Rule) models.Item
}

// configRules are by default
var configRules Config

// LoadRulesConfig function load configuration of rules through yaml file
func LoadRulesConfig() error {
	err := yaml.Unmarshal(data, &configRules)
	if err != nil {
		return fmt.Errorf("couldn't parse yaml file.: %s", err)
	}

	return nil
}
