package handlers

import (
	"html/template"

	"github.com/freer4an/groupie-tracker/pkg/logger"
)

type Handler struct {
	log  *logger.Logger
	temp *template.Template
}

func NewHandler(log *logger.Logger, temp *template.Template) *Handler {
	return &Handler{
		log:  log,
		temp: temp,
	}
}
