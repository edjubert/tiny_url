package config

import "github.com/spf13/viper"

type Configuration struct {
	Server   ServerConfig
	Database DatabaseConfig
}

type ServerConfig struct {
	Port     string
	Mode     string
	LogLevel string
}

type DatabaseConfig struct {
	Driver   string
	Dbname   string
	Username string
	Password string
	Host     string
	Port     string

	MaxIdleConns int
	MaxOpenConns int
}

func New(configPath string) (*Configuration, error) {
	viper.SetConfigFile(configPath)
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var configuration *Configuration
	if err := viper.Unmarshal(&configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}
