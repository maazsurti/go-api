package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rssagg/internal/database"
	"time"

	"github.com/google/uuid"
)

func (configuration *ApiConfiguration) handlerCreateUser (response http.ResponseWriter, request *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(request.Body)

	params := parameters{}
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(response, 400, fmt.Sprintln("Error parsing JSON:", err))
		return 
	}

	user, err := configuration.Database.CreateUser(request.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})

	if err != nil {
		respondWithError(response, 400, fmt.Sprintln("Couldn't create user:", err))
		return 
	}


	respondWithJSON(response, 200, user)
}