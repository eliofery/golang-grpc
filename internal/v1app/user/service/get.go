package service

import (
	"context"

	desc "github.com/eliofery/golang-fullstack/pkg/api/user/v1"
)

// Get ...
func (s service) Get(_ context.Context, _ *desc.GetRequest) (*desc.GetResponse, error) {
	return nil, nil
}
