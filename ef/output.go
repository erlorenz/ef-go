package ef

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON writes a JSON response and returns a nil error.
func JSON(w http.ResponseWriter, code int, body any) error {
	envelope := map[string]any{"data": body}
	b, err := json.Marshal(envelope)
	if err != nil {
		return Error{
			error:   err,
			Code:    http.StatusInternalServerError,
			Type:    contentTypeJSON,
			Message: fmt.Sprintf("Internal error, unable to marshal output: %s", err.Error()),
		}
	}

	w.Header().Set("Content-Type", contentTypeJSON)
	w.WriteHeader(code)
	w.Write(b)
	return nil
}

// HTML writes an HTML resoponse and returns a nil error.
func HTML(w http.ResponseWriter, code int, body string) error {

	w.Header().Set("Content-Type", contentTypeHTML)
	w.WriteHeader(code)
	w.Write([]byte(body))

	return nil
}
