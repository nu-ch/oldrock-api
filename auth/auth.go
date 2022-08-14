package auth

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oldrock-api/db"
)

var client = db.Connect()

func Routes(routers *mux.Router) *mux.Router {
	router := routers.PathPrefix("/auth").Subrouter()

	router.HandleFunc("/login-with-email", loginWithEmail).Methods("POST")
	return router
}

var loginWithEmail = http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
	json.NewEncoder(response).Encode(map[string]bool{"ok": true})
})
