package transport

import (
	"context"
	"encoding/json"
	goerrors "errors"
	"net/http"
	"slices"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/t-lunch/t-lunch-backend/internal/errors"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func WriteError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(ErrorResponse{
		Message: err.Error(),
	})
}

func AuthMiddleware(protectedURL []string, secret string) runtime.Middleware {
	return func(next runtime.HandlerFunc) runtime.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			if !slices.Contains(protectedURL, r.URL.Path) {
				authToken := r.Header.Get("Authorization")
				if authToken == "" {
					WriteError(w, http.StatusUnauthorized, errors.ErrMissingAuthToken)
					return
				}

				jwtToken, err := jwt.ParseWithClaims(
					authToken,
					jwt.MapClaims{},
					func(t *jwt.Token) (interface{}, error) {
						return []byte(secret), nil
					},
				)
				if err != nil {
					WriteError(w, http.StatusUnauthorized, errors.ErrInvalidToken)
					return
				}

				payload, ok := jwtToken.Claims.(jwt.MapClaims)
				if !ok {
					WriteError(w, http.StatusUnauthorized, errors.ErrInvalidToken)
					return
				}

				if int64(payload["exp"].(float64)) < time.Now().Unix() {
					WriteError(w, http.StatusUnauthorized, errors.ErrTokenExpired)
					return
				}
			}
			next(w, r, pathParams)
		}
	}
}

func ErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	WriteError(w, StatusFromError(err), err)
}

func StatusFromError(err error) int {
	switch {
	case goerrors.Is(err, errors.ErrInvalidPassword):
		return http.StatusUnauthorized
	case goerrors.Is(err, errors.ErrInvalidToken):
		return http.StatusUnauthorized
	case goerrors.Is(err, errors.ErrTokenExpired):
		return http.StatusUnauthorized
	case goerrors.Is(err, errors.ErrInvalidRequest):
		return http.StatusBadRequest
	case goerrors.As(err, &errors.ErrUserWithEmailAlreadyExists{}):
		return http.StatusBadRequest
	case goerrors.As(err, &errors.ErrUserAndOwnerAreDifferent{}):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
