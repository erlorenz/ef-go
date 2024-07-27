package ef

import (
	"encoding/json"
	"fmt"
)

// Error implements the Error interface and can be used by the ef.HandlerFunc
// to write an error response.
type Error struct {
	error   error
	Body    []byte
	Code    int
	Type    string
	Message string
}

func (e Error) Error() string {
	return fmt.Sprintf("status %d: %s", e.Code, e.error.Error())
}

// ErrorJSON builds an Error that will be sent as JSON.
// It will be in the format {"error": "<message>"}.
func ErrorJSON(err error, code int, message string) Error {
	bodyData := map[string]string{"error": message}
	b, err := json.Marshal(bodyData)
	if err != nil {
		b = []byte(`{"error":"Internal server error"}`)
	}

	return Error{
		error:   err,
		Body:    b,
		Code:    code,
		Type:    contentTypeJSON,
		Message: message,
	}
}

// ErrorHTML builds an Error that will be sent as an HTML page.
func ErrorHTML(err error, code int, message string) Error {
	html := fmt.Sprintf(`<html><body><h1>%s</h1></body></html>`, message)
	return Error{
		error:   err,
		Body:    []byte(html),
		Code:    200,
		Type:    contentTypeHTML,
		Message: message,
	}
}
