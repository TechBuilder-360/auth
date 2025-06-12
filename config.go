package auth_server_sdk

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type Config struct {
	URL    string
	APIKey string
}

type IAuth interface {
	Registration(ctx context.Context, payload RegistrationRequest) *Response
	ActivateEmail(ctx context.Context, token string) *Response
	Authenticate(ctx context.Context, emailAddress string) *Response
	Login(ctx context.Context, emailAddress, otp string) *Response
	RefreshToken(ctx context.Context, token, refreshToken string) *Response
	Logout(ctx context.Context, token string) *Response
	ValidateToken(ctx context.Context, token string) bool
	GetUser(ctx context.Context, id string) *Response
	GetUserByEmail(ctx context.Context, email string) *Response
}

func New(config Config) IAuth {
	return &config
}

var client *resty.Client

func init() {
	client = resty.New()
}
