package apiuserv2

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"

	desc "github.com/eliofery/golang-fullstack/pkg/api/user/v2"
)

// Get ...
func (s *API) Get(_ context.Context, _ *desc.GetRequest) (*desc.GetResponse, error) {
	return &desc.GetResponse{
		Result: &wrappers.StringValue{Value: "success"},
	}, nil
}
