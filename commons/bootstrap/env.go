package bootstrap

import (
	"github.com/spf13/viper"
	"log/slog"
)

type Env struct {
	ServerAddr     string `mapstructure:"SERVER_ADDRESS"`
	ContextTimeout int    `mapstructure:"CONTEXT_TIMEOUT"`
	AccessTokenKey string `mapstructure:"ACCESS_TOKEN_KEY"`
	AccessTokenAge int    `mapstructure:"ACCESS_TOKEN_AGE"`
	DBUser         string `mapstructure:"DBUSER"`
	DBPassword     string `mapstructure:"DBPASSWORD"`
	DBHost         string `mapstructure:"DBHOST"`
	DBPort         string `mapstructure:"DBPORT"`
	DBDatabase     string `mapstructure:"DBDATABASE"`
}

func NewEnv() *Env {
	env := Env{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		slog.Error("Error reading config file", slog.String("error", err.Error()))
		return nil
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		slog.Error("unable to decode into struct", slog.String("error", err.Error()))
		return nil
	}
	return &env
}
