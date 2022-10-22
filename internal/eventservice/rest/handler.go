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
}

func (eh *EventServiceHandler) GetEvent(w http.ResponseWriter, r *http.Request) {
}

func (eh *EventServiceHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	event := persistence.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
		return
	}

	id, err := eh.dbHandler.AddEvent(event)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", err)
		return
	}
	fmt.Fprintf(w, "{\"id\": \"%s\"}", id)
}
