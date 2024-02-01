package ctxlogger

import (
	"context"

	"log/slog"

	"github.com/lithammer/shortuuid/v4"
)

type ctxLogger struct{}

// ContextWithLogger добавляет логгер в контекст
func ContextWithLogger(ctx context.Context, l *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxLogger{}, l)
}

// LoggerFromContext извлекает логгер из контекста
func LoggerFromContext(ctx context.Context) *slog.Logger {
	if l, ok := ctx.Value(ctxLogger{}).(*slog.Logger); ok {
		return l
	}
	slog.Error("Отсутсвует логгер в контексте", ctx)
	return slog.Default()
}

// key типа для уникального идентификатора запроса
type ctxKeyReqID string

const reqIDKey ctxKeyReqID = "reqID"

// withRequestID добавляет уникальный идентификатор запроса в контекст
func CtxWithNewReqID(ctx context.Context) context.Context {
	reqID := shortuuid.New()
	return context.WithValue(ctx, reqIDKey, reqID)
}

func GetReqID(ctx context.Context) string {
	requestID, ok := ctx.Value(reqIDKey).(string)
	if !ok {
		return "N/A"
	}
	return requestID
}
