package cash_register

import (
	_ "embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

//go:embed rules.yml
var data []byte

// Config represents the structure to store all about limit configuration.
type config struct {
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
	Quantity int64    `yaml:"quantity"`
	NewPrice float64  `yaml:"newPrice,omitempty"`
}

// configRules are by default
var configRules config

// loadConfig function load configuration of rules through yaml file
func loadConfig() error {
	var source config
	err := yaml.Unmarshal(data, &source)
	if err != nil {
		return fmt.Errorf("couldn't parse yaml file.: %s", err)
	}

	return nil
}
