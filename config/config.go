package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AWSConfig struct {
	AccessKeyID     string
	SecretAccessKey string
	Region          string
}

type ConfigLoader interface {
	LoadConfig() (*AWSConfig, error)
}

type EnvConfigLoader struct{}

func NewEnvConfigLoader() ConfigLoader {
	return &EnvConfigLoader{}
}

func (e *EnvConfigLoader) LoadConfig() (*AWSConfig, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	config := &AWSConfig{
		AccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
		SecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
		Region:          os.Getenv("AWS_REGION"),
	}

	if config.Region == "" {
		config.Region = "ap-northeast-2"
	}

	return config, nil
}
