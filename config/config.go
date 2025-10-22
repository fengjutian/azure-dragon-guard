// package config

// import "os"

// var AppPort = getEnv("APP_PORT", "3000")

// func getEnv(key, defaultVal string) string {
// 	if value := os.Getenv(key); value != "" {
// 		return value
// 	}
// 	return defaultVal
// }

package config

import "os"

// 从环境变量读取数据库配置
type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}

func GetDBConfig() DBConfig {
	return DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "123456"),
		Name:     getEnv("DB_NAME", "fiberdb"),
	}
}

func getEnv(key, defaultVal string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultVal
}
