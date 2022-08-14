package middlewares

import (
	"encoding/json"
	"net/http"

	"github.com/oldrock-api/models"
)

func ErrorResponse(message string, writer http.ResponseWriter) {
	temp := &models.Error{StatusCode: http.StatusBadRequest, Message: message}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(temp)
}

func SuccessResponse(fields interface{}, writer http.ResponseWriter) {
	_, err := json.Marshal(fields)

	if err != nil {
		ErrorResponse(http.StatusText(500), writer)
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(fields)
}

func ValidationResponse(fields map[string][]string, writer http.ResponseWriter) {

	temp := &models.Error{StatusCode: http.StatusBadRequest, Message: http.StatusText(400), Errors: fields}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(writer).Encode(temp)
}
