package converter

import (
	"database/sql"

	"github.com/eliofery/golang-fullstack/internal/app/v1app/user/model"
	desc "github.com/eliofery/golang-fullstack/pkg/api/user/v1"
)

// ToUserInfoFromDesc ...
func ToUserInfoFromDesc(signup *desc.SignUpRequest) *model.UserInfo {
	var firstName sql.NullString
	if signup.FirstName != nil {
		firstName = sql.NullString{
			String: signup.FirstName.GetValue(),
			Valid:  true,
		}
	}

	var lastName sql.NullString
	if signup.LastName != nil {
		lastName = sql.NullString{
			String: signup.LastName.GetValue(),
			Valid:  true,
		}
	}

	return &model.UserInfo{
		FirstName: firstName,
		LastName:  lastName,
		Email:     signup.Email,
		Password:  signup.Password,
	}
}
