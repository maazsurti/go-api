package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(response http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response: %v", err)
		response.WriteHeader(500)
		return 
	}

	response.Header().Add("Content-Type", "application/json")
	response.WriteHeader(code)
	response.Write(data)
}

func respondWithError(response http.ResponseWriter, code int,message string){
	if code >  499 {
		log.Println("Network error: ", message)
	}

	type  errorResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(response, code, errorResponse{Error: message})
}