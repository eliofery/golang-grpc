package core

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Codes error
var (
	ErrInternal             = status.Error(codes.Internal, "internal error")
	ErrAccessDenied         = status.Error(codes.PermissionDenied, "access denied")
	ErrTokenNotValid        = status.Error(codes.InvalidArgument, "token not valid")
	ErrWrongLoginOrPassword = status.Error(codes.InvalidArgument, "wrong login or password")
)
