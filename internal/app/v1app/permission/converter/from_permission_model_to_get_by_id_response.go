package converter

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
	desc "github.com/eliofery/golang-grpc/pkg/api/permission/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// FromPermissionModelToGetByIDResponse ...
func FromPermissionModelToGetByIDResponse(permission *model.Permission) *desc.GetByIDResponse {
	var description *wrapperspb.StringValue
	if permission.Description.Valid {
		description = &wrapperspb.StringValue{Value: permission.Description.String}
	}

	return &desc.GetByIDResponse{
		Id:          permission.ID,
		Name:        permission.Name,
		Description: description,
	}
}
