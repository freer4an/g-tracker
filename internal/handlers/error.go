package handlers

import (
	"errors"
	"net/http"
)

func (h *Handler) errorResponse(w http.ResponseWriter, err error, code int) {
	w.WriteHeader(code)
	h.log.Errorf("%s", err)
	http.Error(w, errors.Unwrap(err).Error(), code)
}
