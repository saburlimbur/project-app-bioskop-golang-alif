package utils

import (
	"errors"
	"time"

	"github.com/spf13/viper"
)

type Configuration struct {
	AppName    string
	Port       string
	PathLogg   string
	Limit      string
	Debug      bool
	DB         DatabaseCofig
	Email      EmailConfig
	JWTSecret  string
	JWTExpired time.Duration
}

type EmailConfig struct {
	RequestURL string
	APIKey     string
}

type DatabaseCofig struct {
	Name     string
	Username string
	Password string
	Host     string
	Port     string
	MaxConn  int32
}

func ReadConfiguration() (Configuration, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return Configuration{}, errors.New("err loaded env files")
	}

	viper.AutomaticEnv()

	return Configuration{
		AppName:    viper.GetString("APP_NAME"),
		Port:       viper.GetString("PORT"),
		JWTSecret:  viper.GetString("JWT_SECRET"),
		JWTExpired: viper.GetDuration("JWT_EXPIRED"),
		PathLogg:   viper.GetString("PATH_LOGGING"),
		Limit:      viper.GetString("LIMIT"),
		Debug:      viper.GetBool("DEBUG"),
		DB: DatabaseCofig{
			Name:     viper.GetString("DATABASE_NAME"),
			Username: viper.GetString("DATABASE_USERNAME"),
			Password: viper.GetString("DATABASE_PASSWORD"),
			Host:     viper.GetString("DATABASE_HOST"),
			Port:     viper.GetString("DATABASE_PORT"),
			MaxConn:  viper.GetInt32("DATABASE_MAX_CONN"),
		},
		Email: EmailConfig{
			RequestURL: viper.GetString("REQUEST_EMAIL"),
			APIKey:     viper.GetString("TOKEN_VALUE"),
		},
	}, nil
}
