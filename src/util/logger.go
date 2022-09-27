package util

import (
	"context"
	"go-micro.dev/v4/logger"
	"go-micro.dev/v4/metadata"
)

const (
	traceHeader = "Micro-Trace-Id"
	spanHeader  = "Micro-Span-Id"
)

func LoggerHelper(ctx context.Context) *logger.Helper {
	md, _ := metadata.FromContext(ctx)
	tracefield := map[string]interface{}{
		traceHeader: md[traceHeader],
		spanHeader:  md[spanHeader],
	}
	return logger.NewHelper(logger.DefaultLogger.Fields(tracefield))
}
