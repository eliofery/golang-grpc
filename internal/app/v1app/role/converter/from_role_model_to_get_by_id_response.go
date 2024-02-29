package converter

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	desc "github.com/eliofery/golang-grpc/pkg/api/role/v1"
)

// FromRoleModelToGetByIDResponse ...
func FromRoleModelToGetByIDResponse(role *model.Role) *desc.GetByIDResponse {
	return &desc.GetByIDResponse{
		Id:   role.ID,
		Name: role.Name,
	}
}
