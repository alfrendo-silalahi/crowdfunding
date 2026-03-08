package config

import "service/helper"

type DbConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func GetDbConfig() DbConfig {
	return DbConfig{
		Host:     helper.GetEnv("NEON_POSTGRES_HOST", "localhost"),
		Port:     helper.GetEnv("NEON_POSTGRES_PORT", "5432"),
		User:     helper.GetEnv("NEON_POSTGRES_USER", "postgres"),
		Password: helper.GetEnv("NEON_POSTGRES_PASSWORD", "password"),
		Name:     helper.GetEnv("NEON_POSTGRES_NAME", "crowdfunding"),
	}
}
