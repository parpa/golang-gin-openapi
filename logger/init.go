package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *zap.Logger

func Init() *zap.Logger {
	loggerFileName := os.Getenv("LOG_FILE")

	logLevel := zap.InfoLevel
	switch os.Getenv("LOG_LEVEL") {
	case "info":
		logLevel = zap.InfoLevel
		break
	case "debug":
		logLevel = zap.DebugLevel
		break
	case "warn":
		logLevel = zap.WarnLevel
		break
	case "error":
		logLevel = zap.ErrorLevel
		break
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// lumberjack.Logger is already safe for concurrent use, so we don't need to
	// lock it.
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   loggerFileName,
		MaxSize:    100, // megabytes
		MaxBackups: 3,
		MaxAge:     28, // days
	})
	swSugar := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), w)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		swSugar,
		logLevel,
	)
	logger = zap.New(core)

	return logger
}
