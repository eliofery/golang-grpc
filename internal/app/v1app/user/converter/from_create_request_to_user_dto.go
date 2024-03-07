package converter

import (
	"database/sql"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/dto"
	desc "github.com/eliofery/golang-grpc/pkg/api/user/v1"
)

// FromCreateRequestToUserDTO ...
func FromCreateRequestToUserDTO(req *desc.CreateRequest) *dto.User {
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
