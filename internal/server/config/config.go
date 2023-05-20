package config

import (
	"encoding/json"
	"net"
	"os"
)

type IConfig interface {
	GetHTTPAddress() string
}

type Config struct {
	HTTP HTTP `json:"http"`
}

type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func NewConfig(path string) (*Config, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	config := &Config{}
	err = json.Unmarshal(file, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func (c *Config) GetHTTPAddress() string {
	return net.JoinHostPort(c.HTTP.Host, c.HTTP.Port)
}
