package handlers

import (
	"errors"
	"net/http"
)

var (
	ErrTemplate = errors.New("package error: 'template'")
)

func (h *Handler) errorResponse(w http.ResponseWriter, err error, code int) {
	if err == nil {
		h.log.Errorf("Unexpected error")
		http.Error(w, http.StatusText(code), code)
		return
	}

	h.log.Errorf("%s", err)

	switch {
	case errors.Is(err, ErrTemplate):
		http.Error(w, "Client error", code)
	default:
		http.Error(w, http.StatusText(code), code)
	}
}
