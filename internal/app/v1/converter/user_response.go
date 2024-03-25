package converter

import (
	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	descUser "github.com/eliofery/golang-grpc/pkg/api/app/v1/user"
	rpc "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
)

// UsersToGetAllResponse ...
func UsersToGetAllResponse(users []model.User) *descUser.GetUsersResponse {
	usersResp := make([]*descUser.User, 0, len(users))
	for _, user := range users {
		copyUser := user
		usersResp = append(usersResp, UserToDesc(&copyUser))
	}

	return &descUser.GetUsersResponse{
		Users: usersResp,
		Status: &rpc.Status{
			Code:    int32(codes.OK),
			Message: "users fetched successfully",
		},
	}
}

// UserToGetUserByIDResponse ...
func UserToGetUserByIDResponse(user *model.User) *descUser.GetUserByIDResponse {
	return &descUser.GetUserByIDResponse{
		User: UserToDesc(user),
		Status: &rpc.Status{
			Code:    int32(codes.OK),
			Message: "user fetched successfully",
		},
	}
}

// UserToUpdateUserResponse ...
func UserToUpdateUserResponse(user *model.User) *descUser.UpdateUserResponse {
	return &descUser.UpdateUserResponse{
		User: UserToDesc(user),
		Status: &rpc.Status{
			Code:    int32(codes.OK),
			Message: "user updated successfully",
		},
	}
}
