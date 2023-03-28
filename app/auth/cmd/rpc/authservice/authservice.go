// Code generated by goctl. DO NOT EDIT!
// Source: auth.proto

package authservice

import (
	"context"

	"github.com/tiptok/gz-blog-microsevices/api/protobuf/auth/v1"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GenerateTokenRequest  = v1.GenerateTokenRequest
	GenerateTokenResponse = v1.GenerateTokenResponse
	RefreshTokenRequest   = v1.RefreshTokenRequest
	RefreshTokenResponse  = v1.RefreshTokenResponse
	ValidateTokenRequest  = v1.ValidateTokenRequest
	ValidateTokenResponse = v1.ValidateTokenResponse

	AuthService interface {
		GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error)
		ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error)
		RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	}

	defaultAuthService struct {
		cli zrpc.Client
	}
)

func NewAuthService(cli zrpc.Client) AuthService {
	return &defaultAuthService{
		cli: cli,
	}
}

func (m *defaultAuthService) GenerateToken(ctx context.Context, in *GenerateTokenRequest, opts ...grpc.CallOption) (*GenerateTokenResponse, error) {
	client := v1.NewAuthServiceClient(m.cli.Conn())
	return client.GenerateToken(ctx, in, opts...)
}

func (m *defaultAuthService) ValidateToken(ctx context.Context, in *ValidateTokenRequest, opts ...grpc.CallOption) (*ValidateTokenResponse, error) {
	client := v1.NewAuthServiceClient(m.cli.Conn())
	return client.ValidateToken(ctx, in, opts...)
}

func (m *defaultAuthService) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	client := v1.NewAuthServiceClient(m.cli.Conn())
	return client.RefreshToken(ctx, in, opts...)
}
