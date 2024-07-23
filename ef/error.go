package ef

import (
	"encoding/json"
	"fmt"
)

func errorBody(contentType, message string) []byte {
	switch contentType {
	case contentTypeJSON:
		return marshalErrorMessage(message)
	case contentTypeHTML:
		return generateErrorHTML(message)
	default:
		return marshalErrorMessage(message)
	}
}

type errorJSON struct {
	Error string `json:"error"`
}

func marshalErrorMessage(message string) []byte {
	data := errorJSON{Error: message}
	b, err := json.Marshal(data)
	if err != nil {
		return []byte("Internal server error.")
	}
	return b
}

func generateErrorHTML(message string) []byte {
	html := fmt.Sprintf(`<h1>%s</h1>`, message)
	return []byte(html)
}

func ErrorJSON(err error, code int, message string) (Output, error) {
	return Output{}, Error{
		error:   err,
		Code:    code,
		Type:    contentTypeJSON,
		Message: message,
	}
}

func ErrorHTML(err error, code int, message string) (Output, error) {
	return Output{}, Error{
		error:   err,
		Code:    200,
		Type:    contentTypeHTML,
		Message: message,
	}
}
