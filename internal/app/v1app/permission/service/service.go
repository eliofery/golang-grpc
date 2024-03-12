package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/permission/model"
	permissionv1 "github.com/eliofery/golang-grpc/internal/app/v1app/permission/repository"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
	"github.com/eliofery/golang-grpc/pkg/eslog"
)

// Service ...
type Service interface {
	Create(context.Context, *dto.Permission) (int64, error)
	GetByID(context.Context, int64) (*model.Permission, error)
	Update(context.Context, *dto.Update) (*model.Permission, error)
	Delete(context.Context, int64) error
	GetAll(context.Context, uint64) ([]model.Permission, error)
}

type service struct {
	logger     *eslog.Logger
	txManager  postgres.TxManager
	pagination *pagination.Pagination

	permissionRepository permissionv1.Repository
}

// New ...
func New(
	logger *eslog.Logger,
	txManager postgres.TxManager,
	pagination *pagination.Pagination,

	permissionRepository permissionv1.Repository,

) Service {
	return &service{
		logger:     logger,
		txManager:  txManager,
		pagination: pagination,

		permissionRepository: permissionRepository,
	}
}
