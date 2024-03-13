package converter

import (
	"database/sql"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/dto"
	desc "github.com/eliofery/golang-grpc/pkg/api/permission/v1"
)

// FromUpdateRequestToUpdateDTO ...
func FromUpdateRequestToUpdateDTO(req *desc.UpdateRequest) *dto.Update {
	var name sql.NullString
	if req.Name != nil {
		name.String = req.Name.GetValue()
		name.Valid = true
	}

	var description sql.NullString
	if req.Description != nil {
		description.String = req.Description.GetValue()
		description.Valid = true
	}

	return &dto.Update{
		ID:          req.Id,
		Name:        name,
		Description: description,
	}
}
