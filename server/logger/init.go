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
	case "debug":
		logLevel = zap.DebugLevel
	case "warn":
		logLevel = zap.WarnLevel
	case "error":
		logLevel = zap.ErrorLevel
	}

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	// lumberjack.Logger is already safe for concurrent use, so we don't need to
	// lock it.
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   loggerFileName,
		MaxSize:    5, // megabytes
		MaxBackups: 3,
		MaxAge:     15, // days
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
