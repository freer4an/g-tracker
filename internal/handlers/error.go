package handlers

import (
	"errors"
	"net/http"
)

var (
	ErrTemplate = errors.New("package error: 'template'")
)

func (h *Handler) errorResponse(w http.ResponseWriter, err error, code int) {
	h.log.Errorf("%s", err)
	switch {
	case errors.Is(err, ErrTemplate):
		http.Error(w, ErrTemplate.Error(), code)
	default:
		http.Error(w, errors.Unwrap(err).Error(), code)
	}
}
