package logger

import (
	"fmt"
	"github.com/pangud/pangud/pkg/shared/config"
	"go.uber.org/zap"
)

func New() *zap.Logger {
	var cfg zap.Config
	if config.Env() == "prod" {
		cfg = zap.NewProductionConfig()
	} else {
		cfg = zap.NewDevelopmentConfig()
	}
	cfg.Encoding = config.Logger().Encoding
	cfg.Level = config.Logger().Level
	cfg.OutputPaths = config.Logger().OutputPaths
	cfg.ErrorOutputPaths = config.Logger().ErrorOutputPaths
	log, err := cfg.Build()
	if err != nil {
		fmt.Println(err)
	}
	return log
}
