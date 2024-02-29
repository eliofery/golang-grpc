package converter

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	desc "github.com/eliofery/golang-grpc/pkg/api/role/v1"
)

// FromRoleModelToUpdateResponse ...
func FromRoleModelToUpdateResponse(role *model.Role) *desc.UpdateResponse {
	return &desc.UpdateResponse{
		Id:   role.ID,
		Name: role.Name,
	}
}
