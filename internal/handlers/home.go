package handlers

import (
	"bytes"
	"context"
	"errors"
	"net/http"

	"github.com/freer4an/groupie-tracker/internal/models"
)

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	var data models.HomeData
	var buf bytes.Buffer
	err := data.FillArists()
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}
	if err := temp.ExecuteTemplate(&buf, "index.html", data); err != nil {
		h.errorResponse(w, errors.Join(err, ErrTemplate), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}

func (h *Handler) Context(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "key", "value")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
