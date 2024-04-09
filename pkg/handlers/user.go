package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/logger"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/model"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/service"
	"go.uber.org/zap"
)

// UserHandler holds the user service
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler creates a new UserHandler with the provided UserService
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (uh *UserHandler) UpsertUserHandler(w http.ResponseWriter, r *http.Request) {
    var newUser models.User
    err := json.NewDecoder(r.Body).Decode(&newUser)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    ctx := r.Context()

    existingUser, err := uh.userService.GetUserByEmail(ctx, newUser.Email)
    if err != nil && err.Error() != service.ErrUserNotFound.Error() {
        fmt.Printf("Error: %v\n", err)
        fmt.Printf("service Error: %v\n", service.ErrUserNotFound)
        fmt.Printf("compare: %v\n", errors.Is(err, service.ErrUserNotFound))
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if existingUser != nil {
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(existingUser)
        logger.Log.Info("user is updated", zap.Uint("id", existingUser.ID))
    } else {
        err = uh.userService.CreateUser(ctx, &newUser)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Set status code for successful creation
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newUser)
        logger.Log.Info("user is created", zap.Uint("id", newUser.ID))
    }
}

