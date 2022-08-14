package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/oldrock-api/models"
)

func ErrorResponse(message string, writer http.ResponseWriter) {
	temp := &models.ErrorModel{StatusCode: http.StatusBadRequest, Message: message}
	writer.Header().Set("Content-Type", "application/json")
	ErrorLogger(errors.New(message))
	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(temp)
}

func SuccessResponse(fields, writer http.ResponseWriter) {
	_, err := json.Marshal(fields)

	if err != nil {
		ErrorResponse("Internal server error", writer)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(fields)
}

func ValidationResponse(fields map[string][]string, writer http.ResponseWriter) {
	response := make(map[string]interface{})
	response["errors"] = fields
	response["statusCode"] = http.StatusUnprocessableEntity
	response["message"] = "Validation error"

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(response)
}
