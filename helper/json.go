package helper

import (
	"encoding/json"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&result)
	if err != nil {
		panic(r.Body)
	}
	PanicIfError(err)
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	encoder := json.NewEncoder(w)
	err := encoder.Encode(response)
	PanicIfError(err)
}
