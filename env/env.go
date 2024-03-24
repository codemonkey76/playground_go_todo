package env

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func getVarOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetConnectionString() string {
	godotenv.Load()
	db_name := getVarOrDefault("DB_NAME", "todo")
	db_user := getVarOrDefault("DB_USER", "todo")
	db_pass := getVarOrDefault("DB_PASS", "secret")
	db_host := getVarOrDefault("DB_HOST", "localhost")
	db_port := getVarOrDefault("DB_PORT", "5432")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", db_user, db_pass, db_host, db_port, db_name)

	return connStr
}
