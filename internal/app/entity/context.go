package entity

import (
	"context"

	"github.com/google/uuid"
)

// key типа для уникального идентификатора запроса
type ctxKeyReqID string

const reqIDKey ctxKeyReqID = "reqID"

// withRequestID добавляет уникальный идентификатор запроса в контекст
func WithNewReqID(ctx context.Context) context.Context {
	reqID := uuid.New().String()
	return context.WithValue(ctx, reqIDKey, reqID)
}
