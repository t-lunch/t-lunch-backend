package transport

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func AuthMiddleware(protectedURL []string) runtime.Middleware {
	return func(next runtime.HandlerFunc) runtime.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			if !slices.Contains(protectedURL, r.URL.Path) {
				authHeader := r.Header.Get("Authorization")
				fmt.Println(authHeader)
				if authHeader == "" {
					http.Error(w, "missing auth token", http.StatusUnauthorized)
					return
				}
			}
			next(w, r, pathParams)
		}
	}
}
