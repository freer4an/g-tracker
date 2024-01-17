package handlers

import (
	"html/template"

	"github.com/freer4an/groupie-tracker/pkg/logger"
)

var (
	temp = template.Must(template.ParseGlob("front/templates/*.html"))
)

type Handler struct {
	log *logger.Logger
}

func NewHandler(log *logger.Logger) *Handler {
	return &Handler{
		log: log,
	}
}
