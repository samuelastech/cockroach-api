package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type Config struct {
	Server *Server
}

type Server struct {
	Port string
}

var (
	once   sync.Once
	config *Config
)

func GetConfig() *Config {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println(fmt.Sprintf("[err: %s]", err.Error()))
		}

		port := os.Getenv("SERVER_PORT")
		if port == "" {
			port = "8080"
		}

		config = &Config{
			Server: &Server{port},
		}
	})

	return config
}
