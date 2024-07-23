package ef

import (
	"context"
	"log/slog"
	"net/http"
	"time"
)

type ContextKey string

const LoggerContextKey = ContextKey("logger")
const TimerContextKey = ContextKey("timer")

// WithLogger injects the ef.logger into the context.
func (ef *EF) WithLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), LoggerContextKey, ef.logger)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

// GetLogger returns the ef.logger from the context.
// If it doesn't exist it returns the default slog.Logger.
func getLogger(ctx context.Context) *slog.Logger {
	logger, ok := ctx.Value(LoggerContextKey).(*slog.Logger)
	if !ok {
		return slog.Default()
	}
	return logger
}

// WithTimer adds a request start time into the context.
// It logs the duration after the handler returns.
func (ef *EF) WithTimer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger := getLogger(r.Context())
		start := time.Now()
		ctx := context.WithValue(r.Context(), TimerContextKey, start)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
		since := time.Since(start).Milliseconds()
		logger.Info("Request complete.", "duration", since)
	})
}
