package cashRegister

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	err := loadConfig()
	assert.NoError(t, err)
}