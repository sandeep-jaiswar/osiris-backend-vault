package dotenv

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
)

type Config struct {
	Go struct {
		Server struct {
			Port        int    `yaml:"port"`
			ContextPath string `yaml:"context-path"`
		} `yaml:"server"`
		Session struct {
			Lifetime struct {
				Unit  string `yaml:"unit"`
				Value int    `yaml:"value"`
			} `yaml:"lifetime"`
		} `yaml:"session"`
	} `yaml:"go"`
}

func Enable() {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}

	// Set the path to the YAML configuration file
	yamlConfigPath := filepath.Join(wd, "config", "application.yml")
	fmt.Println(yamlConfigPath)
	// Load configuration from YAML file
	config.AddDriver(yaml.Driver)
	if err := config.LoadFiles(yamlConfigPath); err != nil {
		fmt.Println("Error loading YAML configuration file:", err)
		return
	}

	fmt.Println("Configuration loaded successfully")

	// Create a new Config instance and bind configuration values to it
	var cfg Config
	if err := config.BindStruct("", &cfg); err != nil {
		fmt.Println("Error binding configuration:", err)
		return
	}

	// Set environment variables based on configuration values
	os.Setenv("GO_SERVER_PORT", fmt.Sprintf("%d", cfg.Go.Server.Port))
	os.Setenv("GO_SERVER_CONTEXT_PATH", cfg.Go.Server.ContextPath)
	os.Setenv("GO_SESSION_LIFETIME_UNIT", cfg.Go.Session.Lifetime.Unit)
	os.Setenv("GO_SESSION_LIFETIME_VALUE", fmt.Sprintf("%d", cfg.Go.Session.Lifetime.Value))

	// Print environment variables for verification
	fmt.Println("Environment variables set successfully:")
	fmt.Println("GO_SERVER_PORT:", os.Getenv("GO_SERVER_PORT"))
	fmt.Println("GO_SERVER_CONTEXT_PATH:", os.Getenv("GO_SERVER_CONTEXT_PATH"))
	fmt.Println("GO_SESSION_LIFETIME_UNIT:", os.Getenv("GO_SESSION_LIFETIME_UNIT"))
	fmt.Println("GO_SESSION_LIFETIME_VALUE:", os.Getenv("GO_SESSION_LIFETIME_VALUE"))
}
