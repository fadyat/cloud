package rest

import (
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
}
