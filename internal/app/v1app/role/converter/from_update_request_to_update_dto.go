package converter

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/dto"
	desc "github.com/eliofery/golang-grpc/pkg/api/role/v1"
)

// FromUpdateRequestToRoleDTO ...
func FromUpdateRequestToRoleDTO(req *desc.UpdateRequest) *dto.Role {
	return &dto.Role{
		ID:   req.Id,
		Name: req.Name,
	}
}
