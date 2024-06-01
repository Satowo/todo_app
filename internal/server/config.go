package server

type Config struct {
	DBConfig *DBConfig
	APIConfig *APIConfig
	WebConfig *WebConfig
}

type DBConfig struct {
	DBName       string
	DBUser       string
	DBPass       string
	DBHost       string
	DBPort       string
}

type APIConfig struct {
	APIPort string
}

type WebConfig struct {
	WebURL string
}
