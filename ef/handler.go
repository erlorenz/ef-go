package ef

import (
	"fmt"
	"net/http"
)

const contentTypeJSON = "application/json"
const contentTypeHTML = "text/html"
const defaultType = contentTypeJSON

// Handler returns an ef.Output and an error that is transformed into a standardized response.
// The first argument is a context.Context as a convenience.
// It implements the http.Handler interface.
type Handler func(ctx Context, w http.ResponseWriter, r *http.Request) (Output, error)

// ServeHTTP calls the ef.Handler then uses the
// ef.Output and ef.Error to write the response.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := getLogger(r.Context())

	out, err := h(Context{r.Context()}, w, r)

	if err != nil {

		switch e := err.(type) {
		case Error:
			w.Header().Set("Content-Type", e.Type)
			w.WriteHeader(e.Code)
			w.Write(errorBody(e.Type, e.Message))
			logger.Error(e.Message, "error", e)
			return
		default:
			w.Header().Set("Content-Type", defaultType)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(errorBody(defaultType, e.Error()))
			logger.Error("Internal error.", "error", e)
			return
		}

	}

	w.Header().Set("Content-Type", out.ContentType)
	w.WriteHeader(out.Code)
	w.Write(out.Body)
}

type Output struct {
	ContentType string
	Code        int
	Body        []byte
}

type Error struct {
	error   error
	Code    int
	Type    string
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("status %d: %s", e.Code, e.error.Error())
}
