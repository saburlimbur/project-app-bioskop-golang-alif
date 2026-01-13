package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type LoggerMiddleware struct {
	Logger *zap.Logger
}

func NewLoggerMiddleware(logger *zap.Logger) LoggerMiddleware {
	return LoggerMiddleware{logger}
}

func (m *LoggerMiddleware) LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)

		m.Logger.Info("HTTP Request",
			zap.String("method: ", r.Method),
			zap.String("path: ", r.URL.Path),
			zap.Duration("duration: ", duration),
		)
	})
}
