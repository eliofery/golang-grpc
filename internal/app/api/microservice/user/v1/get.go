package v1

import (
	"context"

	"github.com/golang/protobuf/ptypes/wrappers"

	desc "github.com/eliofery/golang-fullstack/pkg/microservice/user/v1"
)

// Get ...
func (i *UserService) Get(_ context.Context, _ *desc.GetRequest) (*desc.GetResponse, error) {
	return &desc.GetResponse{
		Result: &wrappers.StringValue{Value: "success"},
	}, nil
}
