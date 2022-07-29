package settings

import (
	"os"
	"strconv"
)

type Config struct {
	AdminChatId int64
	TGKey       string
}

// New returns a new Config struct
func New() *Config {
	return &Config{
		AdminChatId: getEnvAsInt("AdminChatId"),
		TGKey:       getEnv("TGKey"),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return "0"
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string) int64 {
	valueStr := getEnv(name)
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}

	return 0
}

// Helper to read an environment variable into a bool or return default value
// func getEnvAsBool(name string, defaultVal bool) bool {
// 	valStr := getEnv(name, "")
// 	if val, err := strconv.ParseBool(valStr); err == nil {
// 		return val
// 	}

// 	return defaultVal
// }

// Helper to read an environment variable into a string slice or return default value
// func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
// 	valStr := getEnv(name, "")

// 	if valStr == "" {
// 		return defaultVal
// 	}

// 	val := strings.Split(valStr, sep)

// 	return val
// }
