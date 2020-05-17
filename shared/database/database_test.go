package database

import (
	"testing"

	"github.com/xmlking/micro-starter-kit/shared/config"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/xmlking/micro-starter-kit/shared/logger"
)

func TestDatabase(t *testing.T) {
	logConf := config.LogConfiguration{Level: "debug", Runtime: "development"}
	logger.InitLogger(logConf)
	_, err := GetDatabaseConnection(config.DatabaseConfiguration{
		Dialect:  "sqlite3",
		Host:     "127.0.0.1",
		Port:     3306,
		Username: "root",
		Password: "123456",
		Database: ":memory:",
	}, logConf)
	if err != nil {
		t.Fatalf("Database connection failed, %v!", err)
	}
}
