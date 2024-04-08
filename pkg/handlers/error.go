package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/model"
)

func ErrorHandler(errorCode int, errorMessage string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if r := recover(); r != nil {
					// Handle panic and format error response
					jsonError := models.Error{
						ErrorCode:    errorCode,
						ErrorMessage: errorMessage,
					}
					jsonResponse(w, errorCode, jsonError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func jsonResponse(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    json.NewEncoder(w).Encode(data)
}
