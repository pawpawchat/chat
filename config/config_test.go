package config

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	flag.Set("env", "testing")
	flag.Parse()

	wd, err := os.Getwd()
	assert.NoError(t, err)

	config, err := LoadConfig(wd + "/../config.yaml")
	assert.NoError(t, err)

	assert.NotNil(t, config)
	assert.NotEmpty(t, config.Env().GRPCServerAddr)
	assert.NotEmpty(t, config.Env().WebSocketAddr)
	assert.NotEmpty(t, config.Env().DbUrl)
}

func TestLoadDefaultConfig(t *testing.T) {
	flag.Set("env", "testing")
	flag.Parse()

	config, err := LoadDefaultConfig()
	assert.NoError(t, err)

	assert.NotNil(t, config)
	assert.NotEmpty(t, config.Env().GRPCServerAddr)
	assert.NotEmpty(t, config.Env().DbUrl)
}

func TestConfigureLogger(t *testing.T) {
	flag.Parse()

	wd, err := os.Getwd()
	assert.NoError(t, err)

	config, err := LoadConfig(wd + "/../config.yaml")
	assert.NoError(t, err)

	assert.NotNil(t, config, config.Env().GRPCServerAddr)

	assert.NoError(t, ConfigureLogger(config))
}
