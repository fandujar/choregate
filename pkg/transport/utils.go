package transport

import (
	"net/http"
)

func RbacMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fake implementation
		next.ServeHTTP(w, r)
	})
}
