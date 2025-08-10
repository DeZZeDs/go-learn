package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type AppConfig struct {
	SmtpConfig SMPTConfig
}

type SMPTConfig struct {
	Email    string
	Password string
	Address  string
	Port     string
}

func BuildConfig() *AppConfig {
	config := &AppConfig{
		SmtpConfig: BuildSMTPConfig(),
	}
	return config
}

func BuildSMTPConfig() SMPTConfig {
	smpt := SMPTConfig{
		Address:  os.Getenv("SMTP_HOST"),
		Email:    os.Getenv("SMTP_USER"),
		Password: os.Getenv("SMTP_PASSWORD"),
		Port:     os.Getenv("SMPT_PORT"),
	}
	return smpt
}
