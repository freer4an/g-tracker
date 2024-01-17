package handlers

import (
	"bytes"
	"errors"
	"net/http"

	"github.com/freer4an/groupie-tracker/internal/models"
	"github.com/freer4an/groupie-tracker/pkg/router"
)

func (h *Handler) BandHandler(w http.ResponseWriter, req *http.Request) {
	buf := bytes.Buffer{}
	id := router.GetParamInt(req.Context(), "id")
	data, err := models.GetBandData(id)
	if err != nil {
		h.errorResponse(w, err, http.StatusInternalServerError)
		return
	}
	if err := temp.ExecuteTemplate(&buf, "band.html", data); err != nil {
		h.errorResponse(w, errors.Join(err, ErrTemplate), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	buf.WriteTo(w)
}
