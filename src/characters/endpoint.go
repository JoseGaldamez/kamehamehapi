package characters

import (
	"log"

	"github.com/JoseGaldamez/kamehamehapi/models"
	"github.com/JoseGaldamez/kamehamehapi/utils"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	ResquestError struct {
		Error string `json:"error"`
	}
)

func MakeEnpoint(s Service) models.Endpoints {
	return models.Endpoints{
		Create: MakeCreateCharactersController(s),
		Get:    MakeGetCharactersController(s),
		GetAll: MakeGetAllCharactersController(s),
		Update: MakeUpdateCharacterController(s),
		Delete: MakeDeleteCharacterController(s),
	}
}

func CreateRouter(path string, router *mux.Router, clientDB *mongo.Client, logger *log.Logger) {

	service := NewCharacterService(clientDB, logger)
	endpoints := MakeEnpoint(service)

	router.HandleFunc(utils.ApiUrlPrefix+path, endpoints.GetAll).Methods("GET")
	router.HandleFunc(utils.ApiUrlPrefix+path, endpoints.Create).Methods("POST")
}
