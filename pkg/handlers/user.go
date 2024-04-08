package handlers

import (
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

func (uh *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := r.Context()
    err = uh.userService.CreateUser(ctx, &newUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set status code for successful creation
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
	logger.Log.Info("user is created", zap.Uint("id",newUser.ID))
}