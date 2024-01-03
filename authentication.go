package auth_server_sdk

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
)

func (c *Config) Registration(ctx context.Context, payload RegistrationRequest) *Response {
	result := new(Response)

	// Create a Resty Client
	client := resty.New()

	_, err := client.R().
		EnableTrace().
		SetHeader("x-auth", c.SecretKey).
		SetContext(ctx).
		SetBody(&payload).
		SetResult(result).
		SetError(result).
		Post(fmt.Sprintf("%s/auth/register", c.URL))

	if err != nil {
		return &Response{
			Status:  false,
			Message: "request failed",
			Error:   err.Error(),
		}
	}

	return result
}

func (c *Config) ActivateEmail(ctx context.Context, token string) *Response {
	result := new(Response)

	// Create a Resty Client
	client := resty.New()

	_, err := client.R().
		EnableTrace().
		SetHeader("x-auth", c.SecretKey).
		SetContext(ctx).
		SetResult(result).
		SetError(result).
		SetPathParam("token", token).
		Get(fmt.Sprintf("%s/auth/activate/{token}", c.URL))

	if err != nil {
		return &Response{
			Status:  false,
			Message: "request failed",
			Error:   err.Error(),
		}
	}

	return result
}

func (c *Config) Authenticate(ctx context.Context, emailAddress string) *Response {
	result := new(Response)

	payload := emailRequest{
		EmailAddress: emailAddress,
	}

	// Create a Resty Client
	client := resty.New()

	_, err := client.R().
		EnableTrace().
		SetHeader("x-auth", c.SecretKey).
		SetContext(ctx).
		SetBody(&payload).
		SetResult(result).
		SetError(result).
		Post(fmt.Sprintf("%s/auth/authentication", c.URL))

	if err != nil {
		return &Response{
			Status:  false,
			Message: "request failed",
			Error:   err.Error(),
		}
	}

	return result
}

func (c *Config) Login(ctx context.Context, emailAddress, otp string) *Response {
	result := new(Response)

	payload := login{
		EmailAddress: emailAddress,
		Otp:          otp,
	}

	// Create a Resty Client
	client := resty.New()

	_, err := client.R().
		EnableTrace().
		SetHeader("x-auth", c.SecretKey).
		SetContext(ctx).
		SetBody(&payload).
		SetResult(result).
		SetError(result).
		Post(fmt.Sprintf("%s/auth/login", c.URL))

	if err != nil {
		return &Response{
			Status:  false,
			Message: "request failed",
			Error:   err.Error(),
		}
	}

	return result
}

func (c *Config) RefreshToken(ctx context.Context, token, refreshToken string) *Response {
	result := new(Response)

	payload := refreshTokenRequest{
		Token:        token,
		RefreshToken: refreshToken,
	}

	// Create a Resty Client
	client := resty.New()

	_, err := client.R().
		EnableTrace().
		SetHeader("x-auth", c.SecretKey).
		SetContext(ctx).
		SetBody(&payload).
		SetResult(result).
		SetError(result).
		Post(fmt.Sprintf("%s/auth/refresh", c.URL))

	if err != nil {
		return &Response{
			Status:  false,
			Message: "request failed",
			Error:   err.Error(),
		}
	}

	return result
}

func (c *Config) Logout(ctx context.Context, token string) *Response {
	result := new(Response)

	// Create a Resty Client
	client := resty.New()

	_, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		SetHeader("x-auth", c.SecretKey).
		SetContext(ctx).
		SetResult(result).
		SetError(result).
		Put(fmt.Sprintf("%s/auth/logout", c.URL))

	if err != nil {
		return &Response{
			Status:  false,
			Message: "request failed",
			Error:   err.Error(),
		}
	}

	return result
}

func (c *Config) ValidateToken(ctx context.Context, token string) bool {
	result := new(Response)

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		SetAuthToken(token).
		SetHeader("x-auth", c.SecretKey).
		SetContext(ctx).
		SetResult(result).
		SetError(result).
		Put(fmt.Sprintf("%s/auth/validate-token", c.URL))

	if err != nil || resp.IsError() {
		return false
	}

	return true
}
