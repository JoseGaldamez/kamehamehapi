package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JoseGaldamez/kamehamehapi/src/characters"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	hostname, _ := os.Hostname()
	_ = godotenv.Load()

	serverAddress := "127.0.0.1:8000"

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	logger.Println("Server staring")
	logger.Println(hostname)

	router := mux.NewRouter()

	characters.CreateRouter("/characters", router)

	server := &http.Server{
		Handler: router,
		Addr:    serverAddress,
	}

	log.Println("====> Listening on: " + serverAddress)

	log.Fatal(server.ListenAndServe())

}
