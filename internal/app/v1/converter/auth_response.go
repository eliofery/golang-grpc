package converter

import (
	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	desc "github.com/eliofery/golang-grpc/pkg/api/app/v1/auth"
	rpc "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
)

// TokenToSignInResponse ...
func TokenToSignInResponse(token *model.Token) *desc.SignInResponse {
	return &desc.SignInResponse{
		Token: &desc.Token{
			Refresh: token.Refresh,
			Access:  token.Access,
		},
		Status: &rpc.Status{
			Code:    int32(codes.OK),
			Message: "user authorized successfully",
		},
	}
}

// TokenToSignUpResponse ...
func TokenToSignUpResponse(token *model.Token) *desc.SignUpResponse {
	return &desc.SignUpResponse{
		Token: &desc.Token{
			Refresh: token.Refresh,
			Access:  token.Access,
		},
		Status: &rpc.Status{
			Code:    int32(codes.OK),
			Message: "user created successfully",
		},
	}
}
