package tracer

import "go.uber.org/zap"

var (
	Logger    *zap.Logger
	crLoggErr error
)

func init() {
	config := zap.NewDevelopmentConfig()
	config.Level.SetLevel(zap.InfoLevel)
	if Logger, crLoggErr = config.Build(); crLoggErr != nil {
		panic("error when setup the log")
	}
}
