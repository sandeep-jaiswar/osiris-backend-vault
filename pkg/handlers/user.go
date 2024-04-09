package handlers

import (
	"context"
	"encoding/json"
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

// UpsertUserHandler handles the creation or update of a user
func (uh *UserHandler) UpsertUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()

	existingUser, err := uh.userService.GetUserByEmail(ctx, newUser.Email)
	if err != nil {
		if err.Error() != service.ErrUserNotFound.Error() {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if existingUser != nil {
		err = uh.updateUser(w, existingUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		err = uh.createUser(w, ctx, &newUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// updateUser updates an existing user
func (uh *UserHandler) updateUser(w http.ResponseWriter, existingUser *models.User) error {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(existingUser)
	logger.Log.Info("user is updated", zap.Uint("id", existingUser.ID))
	return nil
}

// createUser creates a new user
func (uh *UserHandler) createUser(w http.ResponseWriter, ctx context.Context, newUser *models.User) error {
	err := uh.userService.CreateUser(ctx, newUser)
	if err != nil {
		return err
	}

	// Set status code for successful creation
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
	logger.Log.Info("user is created", zap.Uint("id", newUser.ID))
	return nil
}
