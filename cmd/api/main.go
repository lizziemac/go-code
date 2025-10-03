package main

import (
	"net/http"

	"myapp/internal/api/server"
	"myapp/internal/logger"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: server.Setup(),
		// TLS, ReadTimeout, WriteTimeout, etc. can be configured here
	}

	logger.Info("server listening",
		"port", srv.Addr,
	)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatal("server error",
			"error", err,
		)
	}
}
