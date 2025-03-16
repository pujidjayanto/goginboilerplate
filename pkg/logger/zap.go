package logger

import (
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	sugar *zap.SugaredLogger
	once  sync.Once
)

// Init logger regardless their environment. Make same behavior all environment
func Init() {
	once.Do(func() {
		cfg := zap.Config{
			Level:         zap.NewAtomicLevelAt(zapcore.InfoLevel),
			Development:   false,
			Encoding:      "console",
			DisableCaller: false,
			EncoderConfig: zapcore.EncoderConfig{
				TimeKey:        "timestamp",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				FunctionKey:    zapcore.OmitKey,
				MessageKey:     "msg",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.CapitalLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.FullCallerEncoder,
			},
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
		}

		logger := zap.Must(cfg.Build())
		sugar = logger.Sugar()
	})
}

func Debug(msg string, args ...interface{}) {
	sugar.Debugw(msg, args...)
}

func Info(msg string, args ...interface{}) {
	sugar.Infow(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	sugar.Warnw(msg, args...)
}

func Error(msg string, args ...interface{}) {
	sugar.Errorw(msg, args...)
}

func Panic(msg string, args ...interface{}) {
	sugar.Panicw(msg, args...)
}

func Fatal(msg string, args ...interface{}) {
	sugar.Fatalw(msg, args...)
}

func SyncLogger() {
	if sugar != nil {
		_ = sugar.Sync()
	}
}
