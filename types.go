package auth_server_sdk

// Registration ...
type RegistrationRequest struct {
	EmailAddress string  `json:"email_address" validate:"required,email"`
	Avatar       *string `json:"avatar" validate:"url"`
	FirstName    string  `json:"first_name" validate:"required"`
	LastName     string  `json:"last_name" validate:"required"`
	DisplayName  *string `json:"display_name"`
	PhoneNumber  *string `json:"phone_number" validate:"e164"`
}

// EmailRequest ...
type emailRequest struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
}

type refreshTokenRequest struct {
	Token        string `json:"token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// AuthRequest ...
type login struct {
	EmailAddress string `json:"email_address" validate:"required,email"`
	Otp          string `json:"otp" validate:"required,len=6"`
}

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}
