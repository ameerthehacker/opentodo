package config

import (
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/suite"
	"os"
	"testing"
)

type ConfigTestSuite struct {
	suite.Suite
}

func (s *ConfigTestSuite) TestValidConfig() {
	_ = os.Setenv("PORT", "8000")

	assert.Equal(s.T(), GetConfig(Config{
		Name:         "PORT",
		Required:     false,
	}), "8000")
}

func (s *ConfigTestSuite) TestDefaultConfig() {
	assert.Equal(s.T(), GetConfig(Config{
		Name:         "DB_HOST",
		Required:     false,
		DefaultValue: "localhost",
	}), "localhost")
}

func (s *ConfigTestSuite) TestInValidConfig() {
	assert.PanicMatches(s.T(), func() {
		GetConfig(Config{
			Name:         "WRONG",
			Required:     true,
		})
	}, "Config WRONG is required")
}

func TestGetConfig(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
