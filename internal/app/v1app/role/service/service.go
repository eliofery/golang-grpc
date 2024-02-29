package service

import (
	"context"

	"github.com/eliofery/golang-grpc/internal/app/v1app/role/dto"
	"github.com/eliofery/golang-grpc/internal/app/v1app/role/model"
	rolev1 "github.com/eliofery/golang-grpc/internal/app/v1app/role/repository"
	"github.com/eliofery/golang-grpc/internal/core/database/postgres"
	"github.com/eliofery/golang-grpc/internal/core/pagination"
	"github.com/eliofery/golang-grpc/pkg/eslog"
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

	roleRepository rolev1.Repository
}

// New ...
func New(
	logger *eslog.Logger,
	txManager postgres.TxManager,
	pagination *pagination.Pagination,

	roleRepository rolev1.Repository,
) Service {
	return &service{
		logger:     logger,
		txManager:  txManager,
		pagination: pagination,

		roleRepository: roleRepository,
	}
}
