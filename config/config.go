package config

import (
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

var env *string

func init() {
	env = flag.String("env", "stage", "application runtime environment")
}

type Config struct {
	Environment map[string]*Environment `yaml:"environment"`
}

type Environment struct {
	GRPCServerAddr string `yaml:"GRPCServerAddr"`
	LogLevel       string `yaml:"LogLevel"`
	DbUrl          string `yaml:"database_url"`
	WebSocketAddr  string `yaml:"websocket_server_addr"`
}

func (c *Config) Env() *Environment {
	return c.Environment[*env]
}

// LoadConfig загружает конфиг из указанного файла
func LoadConfig(filePath string) (*Config, error) {
	configFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	godotenv.Load(filepath.Join(filepath.Dir(filePath), ".env"))

	return loadConfigFile(configFile)
}

func LoadDefaultConfig() (*Config, error) {
	configFile, err := findDefaultConfigFiles()
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	return loadConfigFile(configFile)
}

func loadConfigFile(file io.Reader) (*Config, error) {
	config := new(Config)
	err := yaml.NewDecoder(file).Decode(config)

	if err != nil {
		return nil, err
	}

	for i := range config.Environment {
		config.Environment[i].DbUrl = os.ExpandEnv(config.Environment[i].DbUrl)
	}

	return config, nil
}

func findDefaultConfigFiles() (*os.File, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	var configFile *os.File
	for range 5 {
		configFile, err = os.Open(filepath.Join(wd, "config.yaml"))
		if err != nil {
			wd = filepath.Join(wd, "..")
			continue
		}
		break
	}

	if configFile == nil {
		return nil, fmt.Errorf("config file isn't exist")
	}

	godotenv.Load(filepath.Join(wd, ".env"))
	return configFile, nil
}

func ConfigureLogger(config *Config) error {
	switch config.Env().LogLevel {
	case "error":
		slog.SetLogLoggerLevel(slog.LevelError)

	case "debug":
		slog.SetLogLoggerLevel(slog.LevelDebug)

	case "info":
		slog.SetLogLoggerLevel(slog.LevelInfo)

	default:
		return fmt.Errorf("undefined log level: %s", config.Env().LogLevel)
	}

	return nil
}
