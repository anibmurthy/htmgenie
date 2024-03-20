package logger

import (
	"context"
	"os"
	"runtime/debug"
	"strconv"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const CtxCorrelationKey = "CorrelationID"

var once sync.Once

var logger *zap.Logger

func Get() *zap.Logger {
	once.Do(func() {
		stdout := zapcore.AddSync(os.Stdout)

		// Uncomment following to write logs to a file
		// [Note: uncomment corresponding  2 lines below for encoding and plugging to core]
		// lumberjack writes logs to rolling files: https://github.com/natefinch/lumberjack
		file := zapcore.AddSync(&lumberjack.Logger{
			// Move the name to be read from configuration
			Filename:   "htmgenie.log",
			MaxSize:    5,
			MaxBackups: 10,
			MaxAge:     14,
			Compress:   true,
		})

		logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
		if err != nil {
			logLevel = int(zap.InfoLevel)
		}

		level := zap.NewAtomicLevelAt(zapcore.Level(logLevel))

		productionCfg := zap.NewProductionEncoderConfig()
		productionCfg.TimeKey = "timestamp"
		productionCfg.EncodeTime = zapcore.ISO8601TimeEncoder

		developmentCfg := zap.NewDevelopmentEncoderConfig()
		developmentCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder

		consoleEncoder := zapcore.NewConsoleEncoder(developmentCfg)

		// Uncomment following to write logs to a file
		fileEncoder := zapcore.NewJSONEncoder(productionCfg)

		core := zapcore.NewCore(consoleEncoder, stdout, level)
		if os.Getenv("APP_ENV") != "development" {
			core = zapcore.NewTee(
				zapcore.NewCore(consoleEncoder, stdout, level),
				zapcore.NewCore(fileEncoder, file, level),
			)
		}

		// It is useful in to include Go version in the long run
		goVersion := "unknown"
		if buildInfo, ok := debug.ReadBuildInfo(); ok {
			goVersion = buildInfo.GoVersion
		}

		logger = zap.New(core).
			With(zap.String("go_version", goVersion))
	})

	return logger
}

// FromCtx returns the Logger associated with the ctx. If no logger
// is associated, the default logger is returned, unless it is nil
// in which case a disabled logger is returned.
func FromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(CtxCorrelationKey).(*zap.Logger); ok {
		return l
	} else if l := logger; l != nil {
		return l
	}

	return zap.NewNop()
}

// WithContext returns a copy of ctx with the Logger attached.
func WithCtx(ctx context.Context, l *zap.Logger) context.Context {
	if lp, ok := ctx.Value(CtxCorrelationKey).(*zap.Logger); ok {
		if lp == l {
			// Do not store same logger.
			return ctx
		}
	}

	return context.WithValue(ctx, CtxCorrelationKey, l)
}
