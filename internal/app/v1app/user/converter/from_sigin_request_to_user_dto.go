package converter

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/dto"
	desc "github.com/eliofery/golang-grpc/pkg/api/user/v1"
)

// FromSignInRequestToUserDTO ...
func FromSignInRequestToUserDTO(req *desc.SignInRequest) *dto.User {
	return &dto.User{
		Email:    req.Email,
		Password: req.Password,
	}
}
