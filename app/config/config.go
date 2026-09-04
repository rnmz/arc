package config

import (
	"os"
	"path"

	"github.com/bytedance/sonic"
)

type Config struct {
	LogConfig    Log    `json:"log"`
	FiberConfig  Fiber  `json:"fiber"`
	DbConfig     DB     `json:"db"`
	ValkeyConfig Valkey `json:"valkey"`
	MasterKey    string `json:"master"`
}

type Log struct {
	Level       string `json:"level"`
	LogFilePath string `json:"file"`
}

type Fiber struct {
	TrustedIPv4 []string `json:"trusted_ipv4"`
	TrustedIPv6 []string `json:"trusted_ipv6"`
}

type DB struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Login    string `json:"login"`
	Password string `json:"password"`
	SSLMode  bool   `json:"ssl_mode"`
}

type Valkey struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
}

func LoadConfig() (Config, error) {
	file := path.Join(".", "config", "config.json")
	data, err := os.ReadFile(file)
	if err != nil {
		return Config{}, err
	}
	var config Config
	if err := sonic.Unmarshal(data, &config); err != nil {
		return Config{}, err
	}
	return config, nil
}
