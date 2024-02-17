package routers

import (
	"log"
	"net/http"

	"github.com/JoseGaldamez/kamehamehapi/corshandler"
	"github.com/JoseGaldamez/kamehamehapi/src/characters"
	"github.com/JoseGaldamez/kamehamehapi/src/jwttoken"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(mongoClient *mongo.Client, logger *log.Logger) http.Handler {

	// Create routers
	router := mux.NewRouter().StrictSlash(true)

	jwttoken.CreateRouter("/token", router)
	characters.CreateRouter("/characters", router, mongoClient, logger)

	handler := corshandler.GetHandleWithCors(router)

	return handler

}
