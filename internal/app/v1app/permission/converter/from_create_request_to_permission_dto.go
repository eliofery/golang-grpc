package converter

import (
	"database/sql"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/dto"
	desc "github.com/eliofery/golang-grpc/pkg/api/permission/v1"
)

// FromCreateRequestToPermissionDTO ...
func FromCreateRequestToPermissionDTO(req *desc.CreateRequest) *dto.Permission {
	var description sql.NullString
	if req.Description != nil {
		description.String = req.Description.GetValue()
		description.Valid = true
	}

	return &dto.Permission{
		Name:        req.GetName(),
		Description: description,
	}
}
