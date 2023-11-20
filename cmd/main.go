package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/freer4an/groupie-tracker/internal/handlers"
	"github.com/freer4an/groupie-tracker/pkg/logger"
	"github.com/freer4an/groupie-tracker/pkg/router"
)

const (
	url  = "https://groupietrackers.herokuapp.com/api/artists"
	port = "8000"
)

func main() {
	logger := logger.InitLogger(os.Stderr)
	err := RunApp(logger)
	if err != nil {
		logger.Fatalf("RunApp: %v", err)
	}
}

func RunApp(logger *logger.Logger) error {
	handler := handlers.NewHandler(logger, nil)
	mux := initRoutes(handler)
	fmt.Println("Launching http://localhost:" + port)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", port), mux)
		if err != nil {
			logger.Fatalf("ListenAndServe error: %v", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	return nil
}

func initRoutes(h *handlers.Handler) *router.Router {
	r := router.NewRouter()
	r.Add("GET", "/", h.Home).Use(h.Context)
	r.Routes()
	// fs := http.FileServer(http.Dir("static"))
	// mux.Handle("/static/", http.StripPrefix("/static", fs))

	return r
}
