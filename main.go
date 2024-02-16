package main

import (
	"log"
	"net/http"
	"os"

	"github.com/JoseGaldamez/kamehamehapi/src/characters"
	"github.com/JoseGaldamez/kamehamehapi/src/jwttoken"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	_ = godotenv.Load()

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	logger.Println("==============> Starting Kamehamehapi <=============")

	router := mux.NewRouter().StrictSlash(true)

	jwttoken.CreateRouter("/token", router)
	characters.CreateRouter("/characters", router)

	log.Println("====> Listening on PORT: " + os.Getenv("PORT"))

	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := corsConfig.Handler(router)

	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), handler))

}
