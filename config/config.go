package config

import (
	"github.com/hvchien216/golang-ecommerce/internal/pkg/env"
	"os"
)

type Config struct {
	DBUrl string  // mandatory
	Env   env.Env // mandatory
	Port  string  // mandatory

}

func NewConfigFromEnv() Config {
	return Config{
		DBUrl: os.Getenv("DB_URL"),
		Env:   env.Env(os.Getenv("APP_ENV")),
		Port:  os.Getenv("PORT"),
	}
}

//dsn := os.Getenv("DBConnectionStr")
//
//s3BucketName := os.Getenv("S3BucketName")
//s3Region := os.Getenv("S3Region")
//s3APIKey := os.Getenv("S3APIKey")
//s3SecretKey := os.Getenv("S3SecretKey")
//s3Domain := os.Getenv("S3Domain")
//secretKey := os.Getenv("SYSTEM_SECRET")
