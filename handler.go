package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type handler struct {
	router mux.Router
}

// NewHandler - creates a new http.Handler with the appropriate function
// handlers
func NewHandler() http.Handler {
	h := handler{
		router: *mux.NewRouter(),
	}

	h.router.HandleFunc("/word_count_per_sentence", h.wordCountPerSentence).Methods("POST")
	h.router.HandleFunc("/total_letter_count", h.totalLetterCount).Methods("POST")

	return h
}

// ServeHTTP - Ensures we implement the http.Handler interface
func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.router.ServeHTTP(w, r)
}

func (h handler) wordCountPerSentence(w http.ResponseWriter, r *http.Request) {
	var req map[string]string

	if err := readRequest(r, &req); err != nil {
		writeResponse(w, http.StatusBadRequest, err)
		return
	}

	// count the words per sentence

	var err error
	if err != nil {
		writeResponse(w, http.StatusOK, err)
		return
	} else {
		writeResponse(w, http.StatusOK, "")
		return
	}
}

func (h handler) totalLetterCount(w http.ResponseWriter, r *http.Request) {
}

func readRequest(r *http.Request, obj interface{}) error {
	if r.Header.Get("Content-Type") != "application/json" {
		return errors.New("error: invalid content-type")
	}

	return json.NewDecoder(r.Body).Decode(&obj)
}

func writeResponse(w http.ResponseWriter, code int, obj interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	return json.NewEncoder(w).Encode(obj)
}
