package utils

import (
	"log"
	"os"
)

type AppConfig struct {
	UsersServiceUrl  string
	MoviesServiceUrl string
}

func LoadCnfFromEnv() *AppConfig {
	return &AppConfig{
		UsersServiceUrl:  loadEnvByKey("USERS_URL"),
		MoviesServiceUrl: loadEnvByKey("MOVIES_URL"),
	}
}

func loadEnvByKey(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Fatalf("can`t find env value by key = %s", key)
	}
	return val
}
				