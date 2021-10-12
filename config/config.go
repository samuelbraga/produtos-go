package config

import (
	"os"
	"produtos/exception"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configImpl struct {
}

func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

func New(filenames ...string) Config {
	godotenv.Load(filenames...)
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &configImpl{}
}
