package converter

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
	desc "github.com/eliofery/golang-grpc/pkg/api/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// FromUserModelToGetByIDResponse ...
func FromUserModelToGetByIDResponse(user *model.User) *desc.GetByIDResponse {
	var firstName *wrapperspb.StringValue
	if user.FirstName.Valid {
		firstName = &wrapperspb.StringValue{Value: user.FirstName.String}
	}

	var lastName *wrapperspb.StringValue
	if user.LastName.Valid {
		lastName = &wrapperspb.StringValue{Value: user.LastName.String}
	}

	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.GetByIDResponse{
		Id:        user.ID,
		FirstName: firstName,
		LastName:  lastName,
		Email:     user.Email,
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
