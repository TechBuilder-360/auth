package auth_server_sdk

type Config struct {
	URL       string
	SecretKey string
}

func New(config Config) *Config {
	return &config
}
