package rest

import (
	"encoding/json"
	"fmt"
	"github.com/fadyat/cloud/internal/persistence"
	"net/http"
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
	fmt.Printf("events: %+v \n error: %+v", events, err)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
		return
	}

	err = json.NewEncoder(w).Encode(events)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
		return
	}
}

func (eh *EventServiceHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
}

func (eh *EventServiceHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event := persistence.Event{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
		return
	}

	id, err := eh.dbHandler.CreateEvent(event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
		return
	}
	fmt.Fprintf(w, "{\"id\": \"%s\"}", id)
}
