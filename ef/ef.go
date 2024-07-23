package ef

import (
	"log/slog"
	"net/http"
)

type EF struct {
	mux        *http.ServeMux
	logger     *slog.Logger
	middleware []Middleware
}

type Middleware func(http.Handler) http.Handler

// New creates a new EF instance and adds some core middleware.
func New(logger *slog.Logger) *EF {
	ef := &EF{
		mux:    http.NewServeMux(),
		logger: logger,
	}
	ef.Use(ef.WithLogger)
	ef.Use(ef.WithTimer)
	return ef
}

// Use appends middleware to the chain.
func (ef *EF) Use(mw Middleware) {
	ef.middleware = append(ef.middleware, mw)
}

// Handle calls the underlying http.ServeMux
// and adds the verb in front of the pattern (1.22 and up)
func (ef *EF) Handle(verb, pattern string, h Handler) {
	ef.mux.Handle(verb+" "+pattern, h)
}

// Get is a convenience method that calls ef.Handle
// with "GET" as the first argument.
func (ef *EF) Get(pattern string, h Handler) {
	ef.Handle(http.MethodGet, pattern, h)
}

// ServeHTTP calls the middleware chain and http.ServeMux.
func (ef *EF) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ef.logger.Info("something")
	wrappedMux(ef.mux, ef.middleware).ServeHTTP(w, r)
}

// wrappedMux calls each middleware in reverse order (inside out)
// and returns the wrapped mux.
func wrappedMux(mux http.Handler, chain []Middleware) http.Handler {
	wrapped := mux
	lastIndex := len(chain) - 1

	for i := lastIndex; i >= 0; i-- {
		wrapped = chain[i](wrapped)
	}
	return wrapped
}
