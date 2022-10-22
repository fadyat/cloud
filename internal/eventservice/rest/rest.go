package rest

import (
	"github.com/fadyat/cloud/internal/persistence"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func ServeAPI(addr string, dbHandler persistence.DatabaseHandler) error {
	r := mux.NewRouter()
	h := NewEventHandler(dbHandler)
	er := r.PathPrefix("/api/v1/events").Subrouter()
	er.Path("").Methods("GET").HandlerFunc(h.GetAllEvents)
	er.Path("").Methods("POST").HandlerFunc(h.CreateEvent)
	er.Path("/{criteria}/{value}").Methods("GET").HandlerFunc(h.GetEvent)

	s := &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return s.ListenAndServe()
}
