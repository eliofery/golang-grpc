package converter

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
	desc "github.com/eliofery/golang-grpc/pkg/api/permission/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// FromPermissionModelToGetAllResponse ...
func FromPermissionModelToGetAllResponse(permissions []model.Permission) *desc.GetAllResponse {
	permissionsResp := make([]*desc.GetAllResponse_Permission, 0, len(permissions))
	for _, permission := range permissions {
		permissionsResp = append(permissionsResp, FromPermissionModelToGetAllResponseUser(permission))
	}

	return &desc.GetAllResponse{
		Permissions: permissionsResp,
	}
}

// FromPermissionModelToGetAllResponseUser ...
func FromPermissionModelToGetAllResponseUser(permission model.Permission) *desc.GetAllResponse_Permission {
	var description *wrapperspb.StringValue
	if permission.Description.Valid {
		description = &wrapperspb.StringValue{Value: permission.Description.String}
	}

	return &desc.GetAllResponse_Permission{
		Id:          permission.ID,
		Name:        permission.Name,
		Description: description,
	}
}
