package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	role "github.com/eliofery/golang-grpc/internal/app/v1app/role/repository"
	rolePermission "github.com/eliofery/golang-grpc/internal/app/v1app/role_permission/repository"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
	"github.com/eliofery/golang-grpc/pkg/eslog"
)

const (
	createPermission = "create_roles"
	readPermission   = "read_roles"
	updatePermission = "update_roles"
	deletePermission = "delete_roles"
)

// Service ...
type Service interface {
	Create(context.Context, *dto.Role) (int64, error)
	GetByID(context.Context, int64) (*model.Role, error)
	Update(context.Context, *dto.Role) (*model.Role, error)
	Delete(context.Context, int64) error
	GetAll(context.Context, uint64) ([]model.Role, error)
}

type service struct {
	logger     *eslog.Logger
	txManager  postgres.TxManager
	pagination *pagination.Pagination

	roleRepository           role.Repository
	rolePermissionRepository rolePermission.Repository
}

// New ...
func New(
	logger *eslog.Logger,
	txManager postgres.TxManager,
	pagination *pagination.Pagination,

	roleRepository role.Repository,
	rolePermissionRepository rolePermission.Repository,
) Service {
	return &service{
		logger:     logger,
		txManager:  txManager,
		pagination: pagination,

		roleRepository:           roleRepository,
		rolePermissionRepository: rolePermissionRepository,
	}
}
