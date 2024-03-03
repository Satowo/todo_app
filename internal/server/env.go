package server

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// 環境変数を取得し、Configにセット
func NewConfig() *Config {
	config := &Config{
		DBConfig: &DBConfig{
			DBName: getEnv("MYSQL_DB"),
			DBUser: getEnv("MYSQL_USER"),
			DBPass: getEnv("MYSQL_PASSWORD"),
			DBHost: getEnv("MYSQL_HOST"),
			DBPort: getEnv("MYSQL_PORT"),
		},
		APIConfig: &APIConfig{
			Port: getEnv("API_PORT"),
		},
	}

	return config
}

// DBの環境変数を取得
func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("could not load .env")
	} 
	
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Fatal("No env: " + key)
	}

	return value
}