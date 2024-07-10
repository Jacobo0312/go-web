package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddr   string `json:"server_addr"`
	DBConnString string `json:"db_conn_string"`
}

func Load() (*Config, error) {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		ServerAddr:   os.Getenv("SERVER_ADDR"),
		DBConnString: os.Getenv("DB_CONN_STRING"),
	}, nil

}
