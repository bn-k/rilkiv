package config

import (
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Auth             Auth
	Db               Db
	Email            Email
	WebClientAddress string
}

type Auth struct {
	Secret     string
	ExpireTime time.Duration
	Encoder string
}

type Email struct {
	Address  string
	Username string
	Password string
	SMTP     string
	Port     int
}

type Db struct {
	Port     int
	Host     string
	Password string
	Database string
	User     string
	Verb     bool
}

func Provide() (Config, error) {
	cf := Config{}
	err := viper.Unmarshal(&cf)
	if err != nil {
		return cf, err
	}

	return cf, nil
}
