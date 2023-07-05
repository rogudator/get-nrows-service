package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rogudator/get-nrows-service/internal/repository"
	"github.com/rogudator/get-nrows-service/internal/service"
	"github.com/rogudator/get-nrows-service/internal/transport"
)

// @title n Rows Service
// @version 1.0
// @description service to get n rows from table

// @host localhost:8080
func main() {
	// Layer holding the information in table
	repo := repository.NewRepository()
	// Layer of business logic
	serv := service.NewService(repo)
	// Layer for rest communication
	rest := transport.NewTransport(serv)

	// Start server in another goroutine
	server := new(Server)
	go func() {
		if err := server.Run("8080", rest.InitRoutes()); err != nil {
			log.Fatalf("server: error occured while running http server: %s", err.Error())
		}

	}()
	log.Print("server: get n rows service started")
	// Shutdown when the service is being closed
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Print("server: get n rows service stopped")
}

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}
