package rest

import (
	"github.com/fadyat/cloud/internal/persistence"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func ServeAPI(httpAddr, httpsAddr string, dbHandler persistence.DatabaseHandler) (httpErrChan, httpsErrChan chan error) {
	r := mux.NewRouter()
	h := NewEventHandler(dbHandler)
	er := r.PathPrefix("/api/v1/events").Subrouter()
	er.Path("").Methods("GET").HandlerFunc(h.GetAllEvents)
	er.Path("").Methods("POST").HandlerFunc(h.CreateEvent)
	er.Path("/{criteria}/{value}").Methods("GET").HandlerFunc(h.GetEvent)

	_http := &http.Server{
		Addr:              httpAddr,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	https := &http.Server{
		Addr:              httpsAddr,
		Handler:           r,
		ReadHeaderTimeout: 5 * time.Second,
	}

	httpErrChan = make(chan error)
	httpsErrChan = make(chan error)

	go func() {
		httpErrChan <- _http.ListenAndServe()
	}()

	go func() {
		// Generate a self-signed certificate with command:
		// 		go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
		//
		// The browser will complain about the self-signed certificate.
		httpsErrChan <- https.ListenAndServeTLS("cert.pem", "key.pem")
	}()

	return httpErrChan, httpsErrChan
}
