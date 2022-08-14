package models

type RegisterWithEmail struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
	DisplayName string `json:"displayName" validate:"required"`
	Age         int16  `json:"age" validate:"required,gt=0,lte=99"`
	DateOfBirth string `json:"dob"`
	BaseSchemas
}

type LoginWithEmail struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}
