package password

import (
	"log/slog"

	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"golang.org/x/crypto/bcrypt"
)

const (
	defaultCost = bcrypt.DefaultCost
)

// Manager ...
type Manager interface {
	GenerateFromPassword(password string, cost ...int) (string, error)
	CompareHashAndPassword(hashedPassword, password string) error
}

// password ...
type password struct {
	logger eslog.Logger
}

// New ...
func New(logger eslog.Logger) Manager {
	return &password{
		logger: logger,
	}
}

// GenerateFromPassword ...
func (p *password) GenerateFromPassword(password string, cost ...int) (string, error) {
	op := "core.authorize.password.GenerateFromPassword"

	if len(cost) == 0 {
		cost = []int{defaultCost}
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cost[0])
	if err != nil {
		p.logger.Debug(op, slog.String("err", err.Error()))
		return "", core.ErrInternal
	}

	return string(hashedPassword), nil
}

// CompareHashAndPassword ...
func (p *password) CompareHashAndPassword(hashedPassword, password string) error {
	op := "core.authorize.password.CompareHashAndPassword"

	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		p.logger.Debug(op, slog.String("err", err.Error()))
		return core.ErrWrongLoginOrPassword
	}

	return nil
}
