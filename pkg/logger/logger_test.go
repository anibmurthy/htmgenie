package logger_test

import (
	"context"
	"testing"

	"github.com/anibmurthy/htmgenie/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestFromCtx(t *testing.T) {

	t.Run("NewNop", func(t *testing.T) {
		ctx := context.Background()

		got := logger.FromCtx(ctx)
		want := zap.NewNop()
		if *got != *want {
			t.Errorf("Expected %v; got %v", *want, *got)
		}
	})

	t.Run("Default logger", func(t *testing.T) {
		ctx := context.Background()
		want := logger.Get()
		got := logger.FromCtx(ctx)
		if want != got {
			t.Errorf("Expected %v; got %v", want, got)
		}
	})

	t.Run("Default logger", func(t *testing.T) {
		want := &zap.Logger{}
		ckey := zap.String(logger.CtxCorrelationKey, "randomvalue")
		ctx := setContextWithCorrelation(ckey, want)
		got := logger.FromCtx(ctx)
		if &want == &got {
			t.Errorf("Expected %v; got %v", &want, &got)
		}
	})
}

func TestGet(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		l := logger.Get()

		if l == nil {
			t.Error("logger instance not retrieved")
		}
	})
}

func setContextWithCorrelation(ckey zapcore.Field, l *zap.Logger) context.Context {
	ctx := context.WithValue(
		context.Background(),
		ckey,
		l,
	)
	ctx = logger.WithCtx(ctx, l)
	return ctx
}
