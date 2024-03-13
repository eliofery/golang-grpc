package converter

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
	desc "github.com/eliofery/golang-grpc/pkg/api/permission/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// FromPermissionModelToUpdateResponse ...
func FromPermissionModelToUpdateResponse(permission *model.Permission) *desc.UpdateResponse {
	var description *wrapperspb.StringValue
	if permission.Description.Valid {
		description = &wrapperspb.StringValue{Value: permission.Description.String}
	}

	return &desc.UpdateResponse{
		Id:          permission.ID,
		Name:        permission.Name,
		Description: description,
	}
}
