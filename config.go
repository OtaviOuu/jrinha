package main

import (
	"fmt"
	"os"
)

type Config struct {
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTSecret  string
}

func initConfig() *Config {
	return &Config{
		// 127.0.0.1:3306
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DP_PORT", "3306")),
		Port:       getEnv("PORT", "8080"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "password"),
		DBName:     getEnv("DB_NAME", "jirinha"),
		JWTSecret:  getEnv("JWT_SECRET", "jwtcontaprangm"),
	}

}

// Fallback = string padrão caso n ache no ambiente
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Variaveis de ambieten a serem injetadas
var Envs = initConfig()
