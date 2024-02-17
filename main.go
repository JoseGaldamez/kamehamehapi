package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/JoseGaldamez/kamehamehapi/database"
	"github.com/JoseGaldamez/kamehamehapi/routers"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	logger.Println("[!] Starting Kamehamehapi")

	// Database connection
	mongoClient := database.GetDatabaseClient(logger)
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	// Create the router
	router := routers.Create(mongoClient, logger)

	// Start the servers
	log.Println("=> Listening on PORT: " + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))

}
