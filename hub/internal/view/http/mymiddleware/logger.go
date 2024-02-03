package mymiddleware

import (
	"bytes"
	"hub/internal/ctxlogger"
	"io"
	"log/slog"

	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

// Добавляет в контекст хэндлера логгер
func Logger(logger *slog.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			// Встраиваем свой ReqID в контекст
			// Для получения reqID := ctxlogger.GetReqID(ctx)
			ctx := ctxlogger.CtxWithNewReqID(r.Context())

			// Встраиваем в логгер поле request_id
			l := logger.With(
				slog.String("req_id", ctxlogger.GetReqID(ctx)),
			)

			//Считываем body, что бы вывести его в лог
			rawBody, err := io.ReadAll(r.Body)
			if err != nil {
				l.Error("Ошбика чтения body", "error", err, "func",
					"mymiddleware.Logger")
			}
			// Restore the io.ReadCloser to it's original state
			r.Body = io.NopCloser(bytes.NewBuffer(rawBody))

			l.Info("new request",
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("method", r.Method),
				slog.String("remote_addr", r.RemoteAddr),
				slog.String("path", r.URL.Path),
				slog.String("req_body", string(rawBody)),
			)

			// Встраиваем логгер в контекст
			ctx = ctxlogger.ContextWithLogger(ctx, l)

			ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

			t1 := time.Now()
			defer func() {
				if time.Since(t1) > 10*time.Millisecond {
					l.Warn("Slow request complete",
						slog.Int("status", ww.Status()),
						slog.String("duration", time.Since(t1).String()),
					)
				}

				if ww.Status() != http.StatusOK {
					l.Warn("Result not 200 OK",
						slog.Int("status", ww.Status()),
						slog.String("duration", time.Since(t1).String()),
					)
				} else {
					l.Info("request completed",
						slog.Int("status", ww.Status()),
						slog.String("duration", time.Since(t1).String()),
					)
				}

			}()

			next.ServeHTTP(ww, r.WithContext(ctx))
		}

		return http.HandlerFunc(fn)
	}
}
