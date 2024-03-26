package converter

import (
	"database/sql"

	"github.com/eliofery/golang-grpc/internal/app/v1/dto"
	desc "github.com/eliofery/golang-grpc/pkg/api/app/v1/auth"
)

// SignInRequestToUser ...
func SignInRequestToUser(req *desc.SignInRequest) *dto.User {
	return &dto.User{
		Email:    req.Email,
		Password: req.Password,
	}
}

// SignUpRequestToUser ...
func SignUpRequestToUser(req *desc.SignUpRequest) *dto.User {
	var firstName sql.NullString
	if req.FirstName != nil {
		firstName.String = req.FirstName.GetValue()
		firstName.Valid = true
	}

	var lastName sql.NullString
	if req.LastName != nil {
		lastName.String = req.LastName.GetValue()
		lastName.Valid = true
	}

	return &dto.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     req.Email,
		Password:  req.Password,
	}
}
