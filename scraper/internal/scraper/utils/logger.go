package utils

import (
	"context"
	"go.uber.org/zap"
)

func getMdcFromContext(ctx context.Context) map[string]string {
	mdc, ok := ctx.Value("mdc").(map[string]string)
	if ok && mdc != nil {
		return mdc
	}
	return make(map[string]string)
}

func UpdateContextMdc(ctx context.Context, mdc map[string]string) context.Context {
	existingMdcs := getMdcFromContext(ctx)
	newMdc := make(map[string]string)
	for key, value := range existingMdcs {
		newMdc[key] = value
	}
	for key, value := range mdc {
		newMdc[key] = value
	}
	ctx = context.WithValue(ctx, "mdc", newMdc)
	return ctx
}

func GetLogger(ctx context.Context, logger *zap.Logger) *zap.Logger {
	mdc := getMdcFromContext(ctx)
	return getLoggerWithMdc(logger, mdc)
}

func getLoggerWithMdc(logger *zap.Logger, mdc map[string]string) *zap.Logger {
	for key, value := range mdc {
		logger = logger.With(zap.String(key, value))
	}
	return logger
}
