package converter

import (
	"github.com/eliofery/golang-grpc/internal/app/v1app/user/model"
	desc "github.com/eliofery/golang-grpc/pkg/api/user/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// FromUsersModelToGetAllResponse ...
func FromUsersModelToGetAllResponse(users []model.User) *desc.GetAllResponse {
	usersResp := make([]*desc.GetAllResponse_User, 0, len(users))
	for _, user := range users {
		usersResp = append(usersResp, FromUserModelToGetAllResponseUser(user))
	}

	return &desc.GetAllResponse{
		Users: usersResp,
	}
}

// FromUserModelToGetAllResponseUser ...
func FromUserModelToGetAllResponseUser(user model.User) *desc.GetAllResponse_User {
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

	return &desc.GetAllResponse_User{
		Id:        user.ID,
		FirstName: firstName,
		LastName:  lastName,
		Email:     user.Email,
		Role: &desc.GetAllResponse_Role{
			Id:   user.Role.ID,
			Name: user.Role.Name,
		},
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
