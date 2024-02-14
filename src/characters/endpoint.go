package characters

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	Character struct {
		Name string `json:"name"`
	}
	ResquestError struct {
		Error string `json:"error"`
	}
)

func CreateRouter(path string, router *mux.Router) {
	router.HandleFunc(path, getAllCharacters).Methods("GET")
	router.HandleFunc(path, createCharacter).Methods("POST")
}

func createCharacter(response http.ResponseWriter, request *http.Request) {
	var character Character

	err := json.NewDecoder(request.Body).Decode(&character)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(ResquestError{Error: "Invalid parameters"})
		return
	}
	if character.Name == "" {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(ResquestError{Error: "Name is requerid"})
		return
	}

	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(character)
}

func getAllCharacters(response http.ResponseWriter, request *http.Request) {
	var characters []Character

	goku := Character{Name: "Goku"}
	vegeta := Character{Name: "Vegeta"}

	characters = append(characters, goku, vegeta)

	json.NewEncoder(response).Encode(characters)
}
