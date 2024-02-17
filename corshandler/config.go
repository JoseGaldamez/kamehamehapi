package corshandler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func GetHandleWithCors(router *mux.Router) http.Handler {
	corsConfig := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})

	handler := corsConfig.Handler(router)
	return handler
}
