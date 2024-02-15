package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JoseGaldamez/kamehamehapi/src/characters"
	"github.com/JoseGaldamez/kamehamehapi/src/jwttoken"
	"github.com/JoseGaldamez/kamehamehapi/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	hostname, _ := os.Hostname()
	_ = godotenv.Load()

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	logger.Println("Server staring")
	logger.Println(hostname)

	router := mux.NewRouter().StrictSlash(true)

	jwttoken.CreateRouter("/token", router)
	characters.CreateRouter("/characters", router)

	server := &http.Server{
		Handler: router,
		Addr:    utils.ServerAddress,
	}

	log.Println("====> Listening on: " + utils.ServerAddress)

	log.Fatal(server.ListenAndServe())

}
