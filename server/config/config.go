package config

// Config is a default config file for server
type Config struct {
	Port string
}

// NewConfig allow you to create a instance of config
func NewConfig() *Config {
	return &Config{
		Port: ":8080",
	}
}
