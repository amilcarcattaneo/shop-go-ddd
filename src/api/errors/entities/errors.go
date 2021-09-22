package entities

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Error string `json:"error"`
}

type ErrorType struct {
	Error  string `json:"error"`
	Status int    `json:"-"`
}

func (er *ErrorType) HandleError(w *http.ResponseWriter) {
	(*w).WriteHeader(er.Status)
	json.NewEncoder(*w).Encode(Error{
		Error: er.Error,
	})
}
