package converter

import (
	"database/sql"
	"time"

	"github.com/eliofery/golang-grpc/internal/app/v1/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1/model"
	descPermission "github.com/eliofery/golang-grpc/pkg/api/app/v1/permission"
	descRole "github.com/eliofery/golang-grpc/pkg/api/app/v1/role"
	descUser "github.com/eliofery/golang-grpc/pkg/api/app/v1/user"
	rpc "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// IDToCreateUserRequest ...
func IDToCreateUserRequest(id int64) *descUser.CreateUserResponse {
	return &descUser.CreateUserResponse{
		Id: id,
		Status: &rpc.Status{
			Code:    int32(codes.OK),
			Message: "user created successfully",
		},
	}
}

// CreateUserRequestToUser ...
func CreateUserRequestToUser(req *descUser.CreateUserRequest) *dto.User {
	var firstName sql.NullString
	if req.FirstName != nil {
		firstName.String = req.FirstName.GetValue()
		firstName.Valid = true
	}

	var lastName sql.NullString
	if req.LastName != nil {
		lastName.String = req.LastName.GetValue()
		lastName.Valid = true
	}

	return &dto.User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     req.Email,
		Password:  req.Password,
	}
}

// UpdateUserRequestToUser ...
func UpdateUserRequestToUser(req *descUser.UpdateUserRequest) *dto.UserUpdate {
	var firstName sql.NullString
	if req.FirstName != nil {
		firstName.String = req.FirstName.GetValue()
		firstName.Valid = true
	}

	var lastName sql.NullString
	if req.LastName != nil {
		lastName.String = req.LastName.GetValue()
		lastName.Valid = true
	}

	var email sql.NullString
	if req.Email != nil {
		email.String = req.Email.GetValue()
		email.Valid = true
	}

	var oldPassword sql.NullString
	if req.OldPassword != nil {
		oldPassword.String = req.OldPassword.GetValue()
		oldPassword.Valid = true
	}

	var newPassword sql.NullString
	if req.NewPassword != nil {
		newPassword.String = req.NewPassword.GetValue()
		newPassword.Valid = true
	}

	var roleID sql.NullInt64
	if req.RoleId != nil {
		roleID.Int64 = req.RoleId.GetValue()
		roleID.Valid = true
	}

	return &dto.UserUpdate{
		ID:          req.Id,
		FirstName:   firstName,
		LastName:    lastName,
		Email:       email,
		OldPassword: oldPassword,
		NewPassword: newPassword,
		RoleID:      roleID,
		UpdatedAt:   time.Now(),
	}
}

// UserPermissionsToDesc ...
func UserPermissionsToDesc(permissions []model.Permission) []*descPermission.Permission {
	permissionsResp := make([]*descPermission.Permission, 0, len(permissions))
	for _, permission := range permissions {
		var description *wrapperspb.StringValue
		if permission.Description.Valid {
			description = &wrapperspb.StringValue{Value: permission.Description.String}
		}

		permissionsResp = append(permissionsResp, &descPermission.Permission{
			Id:          permission.ID,
			Name:        permission.Name,
			Description: description,
		})
	}

	return permissionsResp
}

// UserToDesc ...
func UserToDesc(user *model.User) *descUser.User {
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

	return &descUser.User{
		Id:        user.ID,
		FirstName: firstName,
		LastName:  lastName,
		Email:     user.Email,
		Password:  user.Password,
		Role: &descRole.Role{
			Id:          user.Role.ID,
			Name:        user.Role.Name,
			Permissions: UserPermissionsToDesc(user.Role.Permissions),
		},
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}
