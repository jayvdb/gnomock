// Package gnomockd is an HTTP wrapper around Gnomock
package gnomockd

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/orlangure/gnomock/internal/errors"
)

// Handler returns an HTTP handler ready to serve incoming connections.
func Handler() http.Handler {
	log.Println("Creating http handler...")
	router := mux.NewRouter()
	router.HandleFunc("/start/{name}", startHandler()).Methods(http.MethodPost)
	router.HandleFunc("/stop", stopHandler()).Methods(http.MethodPost)
	router.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Println("Returning http handler...")

	return router
}

func respondWithError(w http.ResponseWriter, err error) {
	w.WriteHeader(errors.ErrorCode(err))

	err = json.NewEncoder(w).Encode(err)
	if err != nil {
		log.Println("can't respond with error:", err)
	}
}
