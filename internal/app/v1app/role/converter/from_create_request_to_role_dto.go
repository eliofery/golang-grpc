package converter

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/dto"
	desc "github.com/eliofery/golang-grpc/pkg/api/role/v1"
)

// FromCreateRequestToRoleDTO ...
func FromCreateRequestToRoleDTO(req *desc.CreateRequest) *dto.Role {
	return &dto.Role{
		Name: req.GetName(),
	}
}
