package models

import "time"

type BaseSchemas struct {
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
	UpdateAt  time.Time `json:"updatedAt" bson:"updatedAt"`
}

type Error struct {
	StatusCode int                 `json:"statusCode"`
	Message    string              `json:"message"`
	Errors     map[string][]string `json:"errors"`
}

type Credential struct {
	AccountID string `json:"accountId bson:"accountId,omitempty"`
}

type PayloadResponse struct {
	Payload map[string]interface{} `json:"payload"`
}
