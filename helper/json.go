package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func ReadFromRequestBody(request *http.Request, result any) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	if err != nil {
		log.Fatal(err)
	}
}

func WriteToResponseBody(writer http.ResponseWriter, response any) {
	writer.Header().Add("content-type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	if err != nil {
		log.Fatal(err)
	}
}
