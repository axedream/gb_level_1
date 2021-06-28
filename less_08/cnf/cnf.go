package cnf

import (
	"os"
)

type Config struct {
	Port string
	Url string
	User string
}

func NewConfig() *Config  {
	return &Config{
	}
}

/**
	Загрузка конфигурации из переменных окружения
 */
func (t *Config) LoadConfig() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8090"
	}

	url := os.Getenv("NAME")
	if url == "" {
		url = "/"
	}

	user := os.Getenv("USER")
	if user == "" {
		user = "Guest"
	}

	t.Port = port
	t.Url = url
	t.User = user
}

