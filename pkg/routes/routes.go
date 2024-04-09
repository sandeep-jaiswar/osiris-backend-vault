package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/klauspost/compress/gzhttp"
	"github.com/rs/cors"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/handlers"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/logger"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/repository"
	"github.com/sandeep-jaiswar/osiris-backend-vault/pkg/service"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB) error {

	serverPort := os.Getenv("GO_SERVER_PORT")
	serverHost := os.Getenv("GO_SERVER_HOST")
	if serverPort == "" {
		serverPort = "8080"
	}

	serverAddr := fmt.Sprintf("%s:%s", serverHost, serverPort)

	router := mux.NewRouter()

	userRepo := repository.NewUserRepository(db)
  	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	router.HandleFunc("/api/v1/user", userHandler.UpsertUserHandler).Methods("POST")

	// Add middleware
	handler := http.Handler(router)
	handler = cors.Default().Handler(handler)
	handler = gzhttp.GzipHandler(handler)

	// Create HTTP server
	httpServer := &http.Server{
		Addr:         serverAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      handler,
	}

	// Start HTTP server in a goroutine
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Error("server error", zap.Error(err))
		}
	}()

	logger.Log.Info("Press CTRL+C to exit...")

	// Wait for exit signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // Blocking operation to wait for signal

	logger.Log.Info("Received shutdown signal...")

	// Graceful shutdown
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		logger.Log.Error("error shutting down server", zap.Error(err))
		return err
	}

	logger.Log.Info("Server shutting down...")

	return nil
}