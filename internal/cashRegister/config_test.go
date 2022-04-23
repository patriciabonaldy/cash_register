package cashRegister

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	err := loadConfig()
	assert.NoError(t, err)
}
