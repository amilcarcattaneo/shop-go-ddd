package http

import (
	"encoding/json"
	"net/http"
)

func (handler *CheckHandler) Check(w http.ResponseWriter, r *http.Request) {
	ok, err := handler.usecase.Check()
	if err != nil {
		err.HandleError(&w)
		return
	}

	if ok {
		json.NewEncoder(w).Encode(map[string]string{"health": "ok"})
	}
}
