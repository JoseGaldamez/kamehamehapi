package characters

import (
	"encoding/json"
	"net/http"

	"github.com/JoseGaldamez/kamehamehapi/models"
)

func MakeCreateCharactersController(service Service) models.Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var character Character
		err := json.NewDecoder(r.Body).Decode(&character)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result, err := service.Create(character)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(result)
	}
}

func MakeGetCharactersController(service Service) models.Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := r.URL.Query().Get("id")
		result, err := service.Get(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(result)
	}
}

func MakeGetAllCharactersController(service Service) models.Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		result, err := service.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(result)
	}
}

func MakeUpdateCharacterController(service Service) models.Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var character Character
		err := json.NewDecoder(r.Body).Decode(&character)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id := r.URL.Query().Get("id")
		result, err := service.Update(id, &character)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(result)
	}
}

func MakeDeleteCharacterController(service Service) models.Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := r.URL.Query().Get("id")
		err := service.Delete(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
