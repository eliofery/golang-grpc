package apiuserv2

import (
	desc "github.com/eliofery/golang-fullstack/pkg/api/user/v2"
)

// API ...
type API struct {
	desc.UnimplementedUserV2ServiceServer
}

// New ...
func New() *API {
	return &API{}
}
