package rest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/fadyat/cloud/internal/persistence"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type EventServiceHandler struct {
	dbHandler persistence.DatabaseHandler
}

func NewEventHandler(dh persistence.DatabaseHandler) *EventServiceHandler {
	return &EventServiceHandler{
		dbHandler: dh,
	}
}

func (eh *EventServiceHandler) GetAllEvents(w http.ResponseWriter, r *http.Request) {
	events, err := eh.dbHandler.FindAll()
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}

	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}
}

func (eh *EventServiceHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	criteria, value := vars["criteria"], vars["value"]

	var event *persistence.Event
	var err error
	switch strings.ToLower(criteria) {
	case "id":
		id, e := hex.DecodeString(value)
		fmt.Printf("id: %s, err: %v\n", id, err)
		if e == nil {
			event, err = eh.dbHandler.FindEvent(id)
			fmt.Printf("event: %v, err: %v\n", event, err)
		}
	case "name":
		event, err = eh.dbHandler.FindEventByName(value)
	default:
		event, err = nil, fmt.Errorf("invalid criteria: %s", criteria)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}

	err = json.NewEncoder(w).Encode(event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}
}

func (eh *EventServiceHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event := persistence.Event{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}

	id, err := eh.dbHandler.CreateEvent(event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"error": "%s"}`, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, `{"id": "%s"}`, id)
}
