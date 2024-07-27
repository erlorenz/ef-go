package ef

import (
	"net/http"
)

func HF(efhf HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		efhf.ServeHTTP(w, r)
	}
}
