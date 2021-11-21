package config

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

type ConfigSuite struct {
	suite.Suite
}

func (suite *ConfigSuite) TestConfig() {
	config, err := ReadConfig("broquiz.example")
	suite.NoError(err)

	suite.Equal("0.0.0.0", config.API.Host)
	suite.Equal(8080, config.API.Port)
}
