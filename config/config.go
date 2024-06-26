package config

import (
	"os"
	"strconv"
)

type HTTPConfig struct {
	Port         string `mapstructure:"PORT"`
	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

type DatabaseConfig struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
}

type RedisConfig struct {
	REDIS_HOST string `mapstructure:"REDIS_HOST"`
	REDIS_PORT string `mapstructure:"REDIS_PORT"`
	DB         int    `mapstructure:"REDIS_DB"`
}

type Config struct {
	HTTPConfig
	DatabaseConfig
	RedisConfig
}

func LoadConfig() *Config {
	redisDB, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
	return &Config{
		HTTPConfig: HTTPConfig{
			Port:         os.Getenv("PORT"),
			ClientOrigin: os.Getenv("CLIENT_ORIGIN"),
		},
		DatabaseConfig: DatabaseConfig{
			DBHost:         os.Getenv("POSTGRES_HOST"),
			DBUserName:     os.Getenv("POSTGRES_USER"),
			DBUserPassword: os.Getenv("POSTGRES_PASSWORD"),
			DBName:         os.Getenv("POSTGRES_DB"),
			DBPort:         os.Getenv("POSTGRES_PORT"),
		},
		RedisConfig: RedisConfig{
			REDIS_HOST: os.Getenv("REDIS_HOST"),
			REDIS_PORT: os.Getenv("REDIS_PORT"),
			DB:         redisDB,
		},
	}
}
