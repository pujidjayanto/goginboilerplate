package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

/*
Initiate a minimal zap production logger
using zap.Must to panic if error happens
*/
func Init() {
	cfg := zap.NewProductionConfig()

	/*
		Customize the encoder configuration
		use key timestamp instead of ts
		use iso time format instead of default
	*/
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.TimeKey = "timestamp"

	// Build the logger from the modified configuration
	logger = zap.Must(cfg.Build())
}

/*
ConfigureLogger is configure minimal logger
to match with environment setup
*/
func ConfigureLogger(env string) {
	var cfg zap.Config

	switch env {
	case "production", "staging":
		cfg = zap.NewProductionConfig()
		cfg.DisableCaller = true
	default:
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.EncoderConfig.TimeKey = "timestamp"

	l := zap.Must(cfg.Build())

	logger = l
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// Info logs an informational message.
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Warn logs a warning message.
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// Error logs an error message.
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// Panic logs a message and then panics.
func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

// Fatal logs a message and then exits the program.
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func GetInstance() *zap.Logger {
	return logger
}

func SyncLogger() {
	if logger != nil {
		_ = logger.Sync()
	}
}
