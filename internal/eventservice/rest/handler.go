package rest

import (
	"encoding/json"
	"fmt"
	"github.com/fadyat/cloud/internal/persistence"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type EventService struct {
	db persistence.DatabaseHandler
}

func NewEventHandler(dh persistence.DatabaseHandler) *EventService {
	return &EventService{
		db: dh,
	}
}

func (es *EventService) GetAllEvents(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	events, err := es.db.FindAll()
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = encoder.Encode(&persistence.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = encoder.Encode(events)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = encoder.Encode(&persistence.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
}

func (es *EventService) GetEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	criteria, value := vars["criteria"], vars["value"]

	var event *persistence.Event
	var err error
	switch strings.ToLower(criteria) {
	case "id":
		event, err = es.db.FindEvent(value)
	case "name":
		event, err = es.db.FindEventByName(value)
	default:
		event, err = nil, fmt.Errorf("invalid criteria: %s", criteria)
	}

	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		_ = encoder.Encode(&persistence.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	err = json.NewEncoder(w).Encode(event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = encoder.Encode(&persistence.ErrorResponse{
			Error: err.Error(),
		})
		return
	}
}

func (es *EventService) CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	event := &persistence.Event{}
	err := json.NewDecoder(r.Body).Decode(event)
	encoder := json.NewEncoder(w)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = encoder.Encode(&persistence.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	id, err := es.db.CreateEvent(event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = encoder.Encode(&persistence.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = fmt.Fprintf(w, `{"id": "%x"}`, id)
}
