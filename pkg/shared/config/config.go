package config

import "go.uber.org/zap"

type Config struct {
	Env    string
	Logger *zap.Config
}

var config *Config

func Logger() *zap.Config {
	if config != nil {
		return config.Logger
	}
	return nil
}
func Env() string {
	if config != nil {
		return config.Env
	}
	return ""
}
func Set(cfg *Config) {
	config = cfg
}
