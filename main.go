package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oldrock-api/auth"
	middlewares "github.com/oldrock-api/handlers"
	"github.com/rs/cors"
)

func main() {
	port := middlewares.DotEnvVariable("PORT")

	middlewares.Logger("Server running on : " + port)

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	router := mux.NewRouter()
	auth.Routes(router)

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})

	handler := c.Handler(router)

	http.ListenAndServe(":"+port, middlewares.LogRequest(handler))
}
