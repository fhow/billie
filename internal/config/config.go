package config

import (
	"fmt"
	"github.com/vrischmann/envconfig"
	"log"
)

type (
	Config struct {
		Server *Server
	}
	Server struct {
		Port string
	}
)

func InitConfig(prefix string) (*Config, error) {

	config := &Config{}

	if err := envconfig.InitWithPrefix(&config, prefix); err != nil {
		return nil, fmt.Errorf("init config failed: %w", err)
	}

	log.Printf("logs %v", config.Server)
	log.Printf("prefix %v", prefix)

	return config, nil
}
