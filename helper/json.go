package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReadFromRequestBody(r *http.Request, result interface{}) {
	decoder := json.NewDecoder(r.Body)
	log.Println("Decoding json")
	err := decoder.Decode(&result)

	if err != nil {
		log.Println("error when decoding json", result)
		panic(r.Body)
	} else {
		log.Println("decoding json successfull")
	}
}

func WriteToResponseBody(w http.ResponseWriter, response interface{}, code int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if code != 204 {
		encoder := json.NewEncoder(w)
		err := encoder.Encode(response)
		PanicIfError(err)
	}
}
