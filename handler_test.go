package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	h := NewHandler()
	assert.NotNil(t, h, "NewHandler returned nil handler")
}

func TestWordCounter(t *testing.T) {

	testCases := []struct {
		name    string
		method  string
		input   map[string]string
		expcode int
		expresp string
		sendhdr bool
	}{
		{
			name:    "there should be no GET",
			method:  "GET",
			expcode: http.StatusNotFound,
			expresp: "404 page not found\n",
			sendhdr: true,
		},
		{
			name:    "POST should return something",
			method:  "POST",
			input:   map[string]string{"somekey": "sentence1? sentence2."},
			expcode: http.StatusOK,
			expresp: "{\"sentence1?\":1,\"sentence2.\":1}\n",
			sendhdr: true,
		},
		{
			name:    "POST with invalid header should return nothing",
			method:  "POST",
			input:   map[string]string{"somekey": "sentence1? sentence2."},
			expcode: http.StatusBadRequest,
			expresp: "{}\n",
			sendhdr: false,
		},
	}

	for _, tc := range testCases {

		b, _ := json.Marshal(&tc.input)
		req, err := http.NewRequest(tc.method, "/word_count_per_sentence", bytes.NewReader(b))
		if err != nil {
			t.Fatal(err)
		}

		if tc.sendhdr {
			req.Header.Add("Content-Type", "application/json")
		}

		recorder := httptest.NewRecorder()
		h := NewHandler()
		h.ServeHTTP(recorder, req)

		assert.Equal(t, tc.expcode, recorder.Code,
			fmt.Sprintf("Invalid code for test [%s]", tc.name))
		assert.Equal(t, tc.expresp, recorder.Body.String(),
			fmt.Sprintf("Invalid response for test [%s]", tc.name))
	}
}

func TestLetterCounter(t *testing.T) {

	testCases := []struct {
		name    string
		method  string
		input   map[string]string
		expcode int
		expresp string
		sendhdr bool
	}{
		{
			name:    "there should be no GET for letter counter",
			method:  "GET",
			expcode: http.StatusNotFound,
			expresp: "404 page not found\n",
			sendhdr: true,
		},
		{
			name:    "POST should return something for letter counter",
			method:  "POST",
			input:   map[string]string{"somekey": "sentence1? sentence2."},
			expcode: http.StatusOK,
			expresp: "\"the text contains 2 Cs, 6 Es, 4 Ns, 2 Ss, 2 Ts, \"\n",
			sendhdr: true,
		},
		{
			name:    "POST with invalid header should return nothing",
			method:  "POST",
			input:   map[string]string{"somekey": "sentence1? sentence2."},
			expcode: http.StatusBadRequest,
			expresp: "{}\n",
			sendhdr: false,
		},
	}

	for _, tc := range testCases {

		b, _ := json.Marshal(&tc.input)
		req, err := http.NewRequest(tc.method, "/total_letter_count", bytes.NewReader(b))
		if err != nil {
			t.Fatal(err)
		}

		if tc.sendhdr {
			req.Header.Add("Content-Type", "application/json")
		}

		recorder := httptest.NewRecorder()
		h := NewHandler()
		h.ServeHTTP(recorder, req)

		assert.Equal(t, tc.expcode, recorder.Code, tc.name)
		assert.Equal(t, tc.expresp, recorder.Body.String(), tc.name)
	}
}
