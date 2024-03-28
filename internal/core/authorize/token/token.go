package token

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"log/slog"
	"math/big"
	"strconv"
	"time"

	"github.com/eliofery/golang-grpc/internal/core"
	"github.com/eliofery/golang-grpc/pkg/eslog"
	"github.com/golang-jwt/jwt/v5"
)

// Manager ...
type Manager interface {
	AccessToken(userID int64) (string, error)
	AccessTokenParse(token string) (int64, error)
	AccessTokenExpires() time.Time
	RefreshToken() (string, error)
	RefreshTokenExpires() time.Time
	GenerateFingerprint(ipAddress, userAgent string) string
	CompareHashFingerprint(hashedFingerprint, ipAddress, userAgent string) error
}

// token ...
type token struct {
	config *Config
	logger eslog.Logger
}

// New ...
func New(config *Config, logger eslog.Logger) Manager {
	return &token{
		config: config,
		logger: logger,
	}
}

// AccessToken ...
func (t *token) AccessToken(userID int64) (string, error) {
	op := "core.authorize.token.AccessToken"

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": strconv.FormatInt(userID, 10),
		"exp": t.AccessTokenExpires().Unix(),
	})

	jwtToken, err := claims.SignedString([]byte(t.config.Secret))
	if err != nil {
		t.logger.Debug(op, slog.String("err", err.Error()))
		return "", core.ErrInternal
	}

	return jwtToken, nil
}

// AccessTokenParse ...
func (t *token) AccessTokenParse(token string) (int64, error) {
	op := "core.authorize.token.AccessTokenParse"

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			t.logger.Debug(op, slog.String("err", "token signing method must be: *jwt.SigningMethodHMAC"))
			return nil, core.ErrInternal
		}

		return []byte(t.config.Secret), nil
	})

	if err != nil {
		t.logger.Debug(op, slog.String("err", err.Error()))
		return 0, core.ErrInternal
	}

	if !parsedToken.Valid {
		t.logger.Debug(op, slog.String("err", "token is invalid"))
		return 0, core.ErrInternal
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		t.logger.Debug(op, slog.String("err", "token claims must be: jwt.MapClaims"))
		return 0, core.ErrInternal
	}

	id, err := claims.GetSubject()
	if err != nil {
		t.logger.Debug(op, slog.String("err", err.Error()))
		return 0, core.ErrInternal
	}

	return strconv.ParseInt(id, 10, 64)
}

// AccessTokenExpires ...
func (t *token) AccessTokenExpires() time.Time {
	return time.Now().Add(t.config.Expires.AccessToken)
}

// RefreshToken ...
func (t *token) RefreshToken() (string, error) {
	op := "core.authorize.token.RefreshToken"

	available := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_-"
	length := 32
	refreshToken := make([]byte, length)
	for i := 0; i < length; i++ {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(available))))
		if err != nil {
			t.logger.Debug(op, slog.String("err", err.Error()))
			return "", core.ErrInternal
		}
		refreshToken[i] = available[index.Int64()]
	}

	return string(refreshToken), nil
}

// RefreshTokenExpires ...
func (t *token) RefreshTokenExpires() time.Time {
	return time.Now().Add(t.config.Expires.RefreshToken)
}

// GenerateFingerprint ...
func (t *token) GenerateFingerprint(ipAddress, userAgent string) string {
	fingerprint := []byte(ipAddress + userAgent)
	hash := sha256.Sum256(fingerprint)

	return hex.EncodeToString(hash[:])
}

// CompareHashFingerprint ...
func (t *token) CompareHashFingerprint(hashedFingerprint, ipAddress, userAgent string) error {
	if hashedFingerprint != t.GenerateFingerprint(ipAddress, userAgent) {
		return core.ErrInternal
	}

	return nil
}
