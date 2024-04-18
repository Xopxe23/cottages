package middleware

import (
	"log/slog"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", "method", r.Method, "url", r.RequestURI, "user", r.Context().Value(ctxUserID))
		next.ServeHTTP(w, r)
	})
}
