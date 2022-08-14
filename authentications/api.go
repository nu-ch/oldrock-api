package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/oldrock-api/database"
	middlewares "github.com/oldrock-api/handlers"
	"github.com/oldrock-api/models"
	validators "github.com/oldrock-api/validators"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db = database.Connect().Database(middlewares.DotEnvVariable("DB_NAME"))

func Routes(routers *mux.Router) *mux.Router {
	router := routers.PathPrefix("/auth").Subrouter()
	router.HandleFunc("/register-with-email", registerWithEmail()).Methods("POST")
	router.HandleFunc("/login-with-email", loginWithEmail()).Methods("POST")

	return router
}

func registerWithEmail() http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		var register models.RegisterWithEmail

		err := json.NewDecoder(request.Body).Decode(&register)
		if err != nil {
			middlewares.ErrorResponse(err.Error(), response)
			middlewares.ErrorLogger(errors.New(err.Error()))
			return
		}

		if ok, errors := validators.ValidateInputs(register); !ok {
			middlewares.ValidationResponse(errors, response)
			return
		}

		collection := db.Collection("account")

		_, errIndex := collection.Indexes().CreateOne(
			context.Background(),
			mongo.IndexModel{
				Keys: bson.M{
					"email": 1,
				},
				Options: options.Index().SetUnique(true),
			},
		)
		if err != nil {
			middlewares.ErrorLogger(errIndex)
			return
		}

		password, err := middlewares.HashPassword(register.Password)

		if err != nil {
			middlewares.ErrorResponse(err.Error(), response)
			middlewares.ErrorLogger(errors.New(err.Error()))
			return
		}

		docs := bson.D{
			primitive.E{Key: "email", Value: register.Email},
			primitive.E{Key: "password", Value: password},
			primitive.E{Key: "displayName", Value: register.DisplayName},
			primitive.E{Key: "age", Value: register.Age},
			primitive.E{Key: "createdAt", Value: time.Now()},
			primitive.E{Key: "updatedAt", Value: time.Now()},
		}

		_, err = collection.InsertOne(context.TODO(), docs)
		if err != nil {
			mongoException := err.(mongo.WriteException)
			if mongoException.WriteErrors[0].Code == 11000 {
				middlewares.ErrorResponse("Email is existing, Please try again", response)
			} else {
				middlewares.ErrorResponse(err.Error(), response)
			}
			middlewares.ErrorLogger(errors.New(err.Error()))
			return
		}

		var account models.Account

		findOptions := options.FindOne().SetProjection(bson.D{primitive.E{Key: "email", Value: 1}})

		err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: register.Email}}, findOptions).Decode(&account)
		if err != nil {
			middlewares.ErrorResponse("Account does not exist", response)
			return
		}

		accessToken, err := middlewares.GenerateJWT(account)

		if err != nil {
			middlewares.ErrorResponse("Invalid authentication token", response)
			return
		}

		middlewares.SuccessResponse(map[string]interface{}{"accessToken": accessToken}, response)

	})
}

func loginWithEmail() http.HandlerFunc {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {

		var login models.LoginWithEmail

		err := json.NewDecoder(request.Body).Decode(&login)
		if err != nil {
			middlewares.ErrorResponse(err.Error(), response)
			middlewares.ErrorLogger(errors.New(err.Error()))
			return
		}

		if ok, errors := validators.ValidateInputs(login); !ok {
			middlewares.ValidationResponse(errors, response)
			return
		}

		var account models.Account

		collection := db.Collection("account")
		err = collection.FindOne(context.TODO(), bson.D{primitive.E{Key: "email", Value: login.Email}}).Decode(&account)
		if err != nil {
			middlewares.ErrorResponse("Account does not exist", response)
			return
		}

		if isMatch := middlewares.CheckPasswordHash(login.Password, account.Password); !isMatch {
			middlewares.ErrorResponse("Password mismatch. Please try again", response)
			return
		}

		accessToken, err := middlewares.GenerateJWT(account)

		if err != nil {
			middlewares.ErrorResponse("Invalid authentication token", response)
			return
		}

		middlewares.SuccessResponse(map[string]interface{}{"accessToken": accessToken}, response)

	})
}
