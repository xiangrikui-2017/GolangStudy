package config

import "time"

type DBConfig struct {
	DriverName string
	Host       string
	Port       string
	Username   string
	Password   string
	Database   string
	Charset    string
}

type LogConfig struct {
	LogDir   string
	LogLevel string
}

type JwtConfig struct {
	Secret      string        `mapstructure:"Secret"`
	TokenExpire time.Duration `mapstructure:"TokenExpire"`
	Issuer      string        `mapstructure:"Issuer"`
}
