package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	URL          string `mapstructure:"url"`
	Token        string `mapstructure:"token"`
	Organization string `mapstructure:"organization"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName(".sq-cli")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(os.Getenv("HOME"))
	viper.AddConfigPath(".")

	viper.SetEnvPrefix("SONAR")
	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("url", "https://sonarcloud.io")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
		// Config file not found is okay if env vars are set
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unable to decode into struct: %w", err)
	}

	return &cfg, nil
}

func SaveConfig(cfg *Config) error {
	viper.Set("url", cfg.URL)
	viper.Set("token", cfg.Token)
	viper.Set("organization", cfg.Organization)

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := filepath.Join(home, ".sq-cli.yaml")
	return viper.WriteConfigAs(configPath)
}
