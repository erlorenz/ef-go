package ef

import (
	"net/http"
)

const contentTypeJSON = "application/json"
const contentTypeHTML = "text/html"
const defaultType = contentTypeJSON

// HandlerFunc returns an error that is transformed into a standardized response.
// It implements the http.Handler interface.
type HandlerFunc func(w http.ResponseWriter, r *http.Request) error

// ServeHTTP calls the ef.HandlerFunc then uses the
// ef.Error to write the response.
// If the error is not an ef.Error it writes a default JSON response.
func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	err := h(w, r)

	if err != nil {
		switch e := err.(type) {
		case Error:
			w.Header().Set("Content-Type", e.Type)
			w.WriteHeader(e.Code)
			w.Write(e.Body)
			return
		default:
			body := `{"error": "Internal server error."}`
			w.Header().Set("Content-Type", defaultType)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(body))
			return
		}

	}
}
