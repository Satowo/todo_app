package server

type Config struct {
	DBConfig *DBConfig
	APIConfig *APIConfig
}

type DBConfig struct {
	DBName       string
	DBUser       string
	DBPass       string
	DBHost       string
	DBPort       string
}

type APIConfig struct {
	Port string
}
