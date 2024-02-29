package converter

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	desc "github.com/eliofery/golang-grpc/pkg/api/role/v1"
)

// FromRolesModelToGetAllResponse ...
func FromRolesModelToGetAllResponse(roles []model.Role) *desc.GetAllResponse {
	rolesResp := make([]*desc.GetAllResponse_Role, 0, len(roles))
	for _, role := range roles {
		rolesResp = append(rolesResp, FromRoleModelToGetAllResponseUser(role))
	}

	return &desc.GetAllResponse{
		Roles: rolesResp,
	}
}

// FromRoleModelToGetAllResponseUser ...
func FromRoleModelToGetAllResponseUser(role model.Role) *desc.GetAllResponse_Role {
	return &desc.GetAllResponse_Role{
		Id:   role.ID,
		Name: role.Name,
	}
}
