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
	}{
		{
			name:    "there should be no GET",
			method:  "GET",
			expcode: http.StatusNotFound,
			expresp: "404 page not found\n",
		},
		{
			name:    "POST should return something",
			method:  "POST",
			input:   map[string]string{"somekey": "sentence1? sentence2."},
			expcode: http.StatusOK,
			expresp: "{Hello world}",
		},
	}

	for _, tc := range testCases {

		b, _ := json.Marshal(&tc.input)
		req, err := http.NewRequest(tc.method, "/word_count_per_sentence", bytes.NewReader(b))
		if err != nil {
			t.Fatal(err)
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
		expcode int
		expresp string
	}{
		{
			name:    "there should be no GET",
			method:  "GET",
			expcode: http.StatusNotFound,
			expresp: "404 page not found\n",
		},
		{
			name:    "POST should return something",
			method:  "POST",
			expcode: http.StatusOK,
			expresp: "{Hello world}",
		},
	}

	for _, tc := range testCases {
		req, err := http.NewRequest(tc.method, "/total_letter_count", nil)
		if err != nil {
			t.Fatal(err)
		}
		recorder := httptest.NewRecorder()
		h := NewHandler()
		h.ServeHTTP(recorder, req)

		assert.Equal(t, tc.expcode, recorder.Code, tc.name)
		assert.Equal(t, tc.expresp, recorder.Body.String(), tc.name)
	}
}
