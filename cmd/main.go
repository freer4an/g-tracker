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
	port = "8000"
)

func main() {
	logger := logger.InitLogger(os.Stdout)
	err := RunApp(logger)
	if err != nil {
		logger.Fatalf("RunApp: %v", err)
	}
}

func RunApp(logger *logger.Logger) error {
	handler := handlers.NewHandler(logger)
	r := initRoutes(handler)
	logger.Infof("Launching http://localhost:" + port)
	go func() {
		err := http.ListenAndServe(fmt.Sprintf(":%s", port), r)
		if err != nil {
			logger.Fatalf("ListenAndServe error: %v", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	logger.Infof("Shutdown Server ...")
	return nil
}

func initRoutes(h *handlers.Handler) *router.Router {
	r := router.NewRouter()
	fs := http.FileServer(http.Dir("front/static"))
	r.ServeStatic("/static/", http.StripPrefix("/static", fs))
	r.Get("/", h.Home)
	r.Get("/band/:id", h.BandHandler)
	return r
}
