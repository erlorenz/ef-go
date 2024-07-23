package ef

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON returns an Output and nil error for a JSON response.
func JSON(code int, body any) (Output, error) {
	envelope := map[string]any{"data": body}
	b, err := json.Marshal(envelope)
	if err != nil {
		return Output{}, Error{
			error:   err,
			Code:    http.StatusInternalServerError,
			Type:    contentTypeJSON,
			Message: fmt.Sprintf("Internal error, unable to marshal output: %s", err.Error()),
		}
	}

	return Output{
		ContentType: contentTypeJSON,
		Code:        code,
		Body:        b,
	}, nil
}

// HTML returns the Output and nil error for an HTML response.
func HTML(code int, body string) (Output, error) {
	return Output{
		ContentType: contentTypeHTML,
		Code:        code,
		Body:        []byte(body),
	}, nil
}
