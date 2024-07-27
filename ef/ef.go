package ef

import (
	"net/http"
)

type EF struct {
	Mux *http.ServeMux
}

// New embeds the mux and returns a new *EF.
func New(mux *http.ServeMux) *EF {
	return &EF{mux}
}

// Handle registers the handler to the mux
func (ef *EF) Handle(pattern string, h HandlerFunc) {
	ef.Mux.Handle(pattern, h)
}

// Get prepends GET to the pattern and calls ef.Handle.
func (ef *EF) Get(pattern string, h HandlerFunc) {
	ef.Handle("GET "+pattern, h)
}

// Post prepends POST to the pattern and registers the handler.
func (ef *EF) Post(mux *http.ServeMux, pattern string, h HandlerFunc) {
	ef.Handle("POST "+pattern, h)
}

// Patch prepends PATCH to the pattern and registers the handler.
func (ef *EF) Patch(mux *http.ServeMux, pattern string, h HandlerFunc) {
	ef.Handle("PATCH "+pattern, h)
}

func (ef *EF) Put(mux *http.ServeMux, pattern string, h HandlerFunc) {
	ef.Handle("PUT "+pattern, h)
}

func (ef *EF) Delete(mux *http.ServeMux, pattern string, h HandlerFunc) {
	ef.Handle("DELETE "+pattern, h)
}
