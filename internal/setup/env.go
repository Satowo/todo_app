package setup

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type DBConfig struct {
	DBName       string
	DBUser       string
	DBPass       string
	DBHost       string
	DBPort       string
}

// DBの環境変数を取得
func GetDBEnv() *DBConfig {
	err := loadEnv()
	if err != nil {
		return nil
	}

	return &DBConfig{
		DBName:       os.Getenv("MYSQL_DB"),
		DBUser:       os.Getenv("MYSQL_USER"),
		DBPass:       os.Getenv("MYSQL_PASSWORD"),
		DBHost:       os.Getenv("MYSQL_HOST"),
		DBPort:       os.Getenv("MYSQL_PORT"),
	}
}

// 環境変数の読み込み
func loadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}
	return nil
}