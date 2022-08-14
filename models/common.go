package models

type ErrorModel struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

type Credential struct {
	AccountID string `json:"accountId bson:"accountId,omitempty"`
}
