package converter

import (
	"database/sql"
	"time"

	"github.com/eliofery/golang-grpc/internal/app/v1app/user/dto"
	desc "github.com/eliofery/golang-grpc/pkg/api/user/v1"
)

// FromUpdateRequestToUpdateDTO ...
func FromUpdateRequestToUpdateDTO(req *desc.UpdateRequest) *dto.Update {
	var firstName sql.NullString
	if req.FirstName != nil {
		firstName.String = req.FirstName.GetValue()
		firstName.Valid = true
	}

	var lastName sql.NullString
	if req.FirstName != nil {
		lastName.String = req.LastName.GetValue()
		lastName.Valid = true
	}

	var email sql.NullString
	if req.Email != nil {
		email.String = req.Email.GetValue()
		email.Valid = true
	}

	var oldPassword sql.NullString
	if req.OldPassword != nil {
		oldPassword.String = req.OldPassword.GetValue()
		oldPassword.Valid = true
	}

	var newPassword sql.NullString
	if req.OldPassword != nil && req.NewPassword != nil {
		newPassword.String = req.NewPassword.GetValue()
		newPassword.Valid = true
	}

	return &dto.Update{
		ID:          req.Id,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		OldPassword: oldPassword,
		NewPassword: newPassword,
		UpdatedAt:   time.Now(),
	}
}
