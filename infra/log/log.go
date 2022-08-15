package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.SugaredLogger
var ZapLogger *zap.Logger

func init() {
	cfg := zap.Config{
		Encoding:         "json",
		Level:            zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stdout"},
		EncoderConfig:    zap.NewProductionEncoderConfig(),
	}

	ZapLogger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	logger = ZapLogger.Sugar()
	logger.Sync()
}

type LogContext map[string]interface{}

func Info(msg string, ctx *LogContext) {
	logger.Infow(msg, zap.Any("event", ctx))
}

func Warn(msg string, ctx *LogContext) {
	logger.Warnw(msg, zap.Any("event", ctx))
}

func Error(msg string, err error, ctx *LogContext) {
	logger.Errorw(msg, zap.Error(err), zap.Any("event", ctx))
}

func Fatal(msg string, err error, ctx *LogContext) {
	logger.Fatalw(msg, zap.Error(err), zap.Any("event", ctx))
}
