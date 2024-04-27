package main

import "net/http"

func handlerError(response http.ResponseWriter, request *http.Request){
	respondWithError(response, 400, "Something went wrong")
}