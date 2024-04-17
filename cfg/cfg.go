package cfg

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	AppAddr = "APP_ADDR"
	AppPort = "APP_PORT"

	PsqlLogin = "PSQL_LOGIN"
	PsqlPass  = "PSQL_PASS"
	PsqlDB    = "PSQL_DB"
	PsqlAddr  = "PSQL_ADDR"
	PsqlPort  = "PSQL_PORT"

	SecretKey = "SECRET_KEY"
)

type AppConfig struct {
	AppAddr string
	AppPort string
	AppMode string

	PsqlLogin string
	PsqlPass  string
	PsqlDB    string
	PsqlAddr  string
	PsqlPort  string

	SecretKey string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found!")
	}
}

var appConfig *AppConfig

func GetAppConfig() *AppConfig {
	if appConfig == nil {
		appConfig = &AppConfig{
			AppAddr: getEnv(AppAddr, "127.0.0.1"),
			AppPort: getEnv(AppPort, "8080"),

			PsqlLogin: getEnv(PsqlLogin, "idon"),
			PsqlPass:  getEnv(PsqlPass, "test_db_password"),
			PsqlDB:    getEnv(PsqlDB, "network"),
			PsqlAddr:  getEnv(PsqlAddr, "127.0.0.1"),
			PsqlPort:  getEnv(PsqlPort, "5432"),

			SecretKey: getEnv(SecretKey, "very_very_secret_key"),
		}
	}

	return appConfig
}

func getEnv(key string, defaultVal string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}

func getEnvAsUint16(name string, defaultVal uint16) uint16 {
	valueStr := getEnv(name, "")
	if value, err := strconv.ParseUint(valueStr, 10, 16); err == nil {
		return uint16(value)
	}

	return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
	valStr := getEnv(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}

	return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
	valStr := getEnv(name, "")

	if valStr == "" {
		return defaultVal
	}

	val := strings.Split(valStr, sep)

	return val
}
