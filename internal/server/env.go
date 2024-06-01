package server

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var config *Config

// 環境変数を取得し、configにセット
func NewConfig() {
	config = &Config{
		DBConfig: &DBConfig{
			DBName: getEnv("MYSQL_DB"),
			DBUser: getEnv("MYSQL_USER"),
			DBPass: getEnv("MYSQL_PASSWORD"),
			DBHost: getEnv("MYSQL_HOST"),
			DBPort: getEnv("MYSQL_PORT"),
		},
		APIConfig: &APIConfig{
			APIPort: getEnv("API_PORT"),
		},
		WebConfig : &WebConfig{
			WebURL: getEnv("WEB_URL"),
		},
	}
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