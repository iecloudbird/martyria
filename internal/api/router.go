package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"
)

// Router sets up all API routes and middleware.
func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()

	// Health
	mux.HandleFunc("GET /health", h.Health)

	// V1 API
	mux.HandleFunc("GET /v1/authors", h.ListAuthors)
	mux.HandleFunc("GET /v1/authors/{slug}", h.GetAuthor)
	mux.HandleFunc("GET /v1/authors/{slug}/quotes", h.GetAuthorQuotes)
	mux.HandleFunc("GET /v1/quotes", h.ListQuotes)
	mux.HandleFunc("GET /v1/quotes/random", h.RandomQuote)
	mux.HandleFunc("GET /v1/quotes/daily", h.DailyQuote)
	mux.HandleFunc("GET /v1/quotes/{id}", h.GetQuote)
	mux.HandleFunc("GET /v1/topics", h.ListTopics)
	mux.HandleFunc("GET /v1/topics/{slug}/quotes", h.GetTopicQuotes)

	// Wrap with middleware chain
	var handler http.Handler = mux
	handler = CORSMiddleware(handler)
	handler = LoggingMiddleware(handler)
	handler = RecoveryMiddleware(handler)

	return handler
}

// --- Middleware ---

// RecoveryMiddleware catches panics and returns 500.
func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("PANIC: %v", err)
				writeJSON(w, http.StatusInternalServerError, map[string]string{
					"error": "internal server error",
				})
			}
		}()
		next.ServeHTTP(w, r)
	})
}

// LoggingMiddleware logs each request.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		sw := &statusWriter{ResponseWriter: w, status: 200}
		next.ServeHTTP(sw, r)
		log.Printf("%s %s %d %s", r.Method, r.URL.Path, sw.status, time.Since(start).Round(time.Millisecond))
	})
}

// CORSMiddleware allows cross-origin requests.
func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-API-Key")
		w.Header().Set("Access-Control-Max-Age", "86400")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// RateLimitMiddleware (placeholder for Redis-backed limiter later).
func RateLimitMiddleware(requestsPerMin int) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// TODO: Implement Redis-backed rate limiting
			next.ServeHTTP(w, r)
		})
	}
}

// --- Helpers ---

type statusWriter struct {
	http.ResponseWriter
	status int
}

func (sw *statusWriter) WriteHeader(code int) {
	sw.status = code
	sw.ResponseWriter.WriteHeader(code)
}

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("JSON encode error: %v", err)
	}
}

func queryParam(r *http.Request, key, defaultVal string) string {
	v := r.URL.Query().Get(key)
	v = strings.TrimSpace(v)
	if v == "" {
		return defaultVal
	}
	return v
}
