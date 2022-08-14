package authentication

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/oldrock-api/db"
	middlewares "github.com/oldrock-api/handlers"
	"github.com/oldrock-api/models"
	validators "github.com/oldrock-api/validators"
)

var client = db.Connect()

func Routes(routers *mux.Router) *mux.Router {
	router := routers.PathPrefix("/auth").Subrouter()
	router.HandleFunc("/login-with-email", loginWithEmail()).Methods("POST")
	return router
}

func loginWithEmail() http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		var login models.LoginWithEmail
		err := json.NewDecoder(request.Body).Decode(&login)
		if err != nil {
			middlewares.ErrorResponse(err.Error(), response)
			return
		}

		if ok, errors := validators.ValidateInputs(login); !ok {
			middlewares.ValidationResponse(errors, response)
			return
		}

		// middlewares.ErrorResponse("test", response)
	})
}
