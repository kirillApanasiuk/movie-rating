package http

import (
	"encoding/json"
	"errors"
	"log"
	"metadata.com/internal/controller"
	"metadata.com/model"
	"net/http"
	"strconv"
)

type HttpHandler struct {
	ctrl *controller.Controller
}

func New(ctrl *controller.Controller) *HttpHandler {
	return &HttpHandler{ctrl: ctrl}
}

func (h *HttpHandler) Handle(w http.ResponseWriter, req *http.Request) {
	recordID := model.RecordID(req.FormValue("id"))
	if recordID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	recordType := model.RecordType(req.FormValue("type"))
	if recordType == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAggregatedRating(req.Context(), recordID, recordType)
		if err != nil && errors.Is(err, controller.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("failed to encode response: %v\n", err)
		}

	case http.MethodPut:
		userID := model.UserId(req.FormValue("userId"))
		v, err := strconv.ParseFloat(req.FormValue("value"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := h.ctrl.PutRating(req.Context(), recordID, recordType, &model.Rating{
			UserID: userID,
			Value:  model.RatingValue(v),
		}); err != nil {
			log.Printf("failed to encode response: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

}
