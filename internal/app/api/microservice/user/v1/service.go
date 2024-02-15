package v1

import pb "github.com/eliofery/golang-fullstack/pkg/microservice/user/v1"

// UserService ...
type UserService struct {
	pb.UnimplementedUserV1ServiceServer
}

// New ...
func New() *UserService {
	return &UserService{}
}
