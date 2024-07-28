package main

import (
	"encoding/json"
	"errors"
	"net/http"
)

type jsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// function to read json object
func (app *Config) readJSON(w http.ResponseWriter, r *http.Request, data any) error {

	maxBytes := 1048576 // 1 megabyte
	// making a limitation on the size of an uploaded json file. json file should be less than 1mb.
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	// creating a new decoder and passed r.body to read from.
	dec := json.NewDecoder(r.Body)
	// decoding json values to data
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	// ensuring there is only one json data is passed on the request
	// err = dec.Decode(&struct{}{})
	// if err != io.EOF {
	// 	return errors.New("body must have only a single JSON value")
	// }
	// this is little more convenient way to check
	if dec.More() {
		return errors.New("body must have only a single JSON value")
	}

	return nil
}

// function to write json
func (app *Config) writeJSON(w http.ResponseWriter, status int, data any, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}

	return nil
}

func (app *Config) errorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest

	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload jsonResponse
	payload.Error = true
	payload.Message = err.Error()

	return app.writeJSON(w, statusCode, payload)
}
