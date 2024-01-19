package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

// key типа для уникального идентификатора запроса
type contextKey string

const requestIDKey contextKey = "requestID"

// generateRequestID генерирует короткий уникальный идентификатор запроса (ULID)
func generateRequestID() string {
	return uuid.New().String()
}

// withRequestID добавляет уникальный идентификатор запроса в контекст
func withRequestID(ctx context.Context) context.Context {
	requestID := generateRequestID()
	return context.WithValue(ctx, requestIDKey, requestID)
}

// getRequestID извлекает уникальный идентификатор запроса из контекста
func getRequestID(ctx context.Context) string {
	requestID, ok := ctx.Value(requestIDKey).(string)
	if !ok {
		return "N/A"
	}
	return requestID
}

// handler обработчик HTTP запроса
func handler(w http.ResponseWriter, r *http.Request) {
	// Извлекаем текущий контекст
	ctx := r.Context()

	// Извлекаем уникальный идентификатор запроса из контекста
	requestID := getRequestID(ctx)
	fmt.Printf("Received request with ID: %s\n", requestID)

	// Выполняем какую-то логику обработки запроса

	// Продолжаем выполнение обработчика с обновленным контекстом
	// ...

	// Отправляем ответ клиенту
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Пример использования: создаем HTTP сервер с обработчиком
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
