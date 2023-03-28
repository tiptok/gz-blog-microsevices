package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/tiptok/gocomm/pkg/log"
	"github.com/tiptok/gz-blog-microsevices/pkg/config"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

func NewManager(logger log.Log, conf *config.Config) *Manager {
	return &Manager{
		secret:  conf.JWT.Secret,
		expires: conf.JWT.Expires,
		logger:  logger,
	}
}

type Manager struct {
	secret  string
	expires time.Duration
	logger  log.Log
}

type UserClaims struct {
	ID uint64 `json:"id"`
	jwt.StandardClaims
}

func (manager *Manager) Generate(id uint64) (string, error) {
	claims := UserClaims{
		ID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(manager.expires).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(manager.secret))
}

func (manager *Manager) Validate(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			_, ok := token.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				return nil, status.Error(codes.Unauthenticated, "invalid token")
			}

			return []byte(manager.secret), nil
		},
	)

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}

	return claims, nil
}
