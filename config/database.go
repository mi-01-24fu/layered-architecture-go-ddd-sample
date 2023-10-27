package config

import (
	"database/sql"
	"fmt"
	"github.com/friendsofgo/errors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
)

type dbConfig struct {
	dbms         string
	databaseName string
	databaseHost string
	databasePort string
	username     string
	password     string
	protocol     string
}

func NewDB() (*sql.DB, error) {
	config, err := loadEnv()
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", config.username, config.password, config.protocol, config.databaseHost, config.databasePort, config.databaseName)
	db, err := sql.Open(config.dbms, connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func loadEnv() (dbConfig, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return dbConfig{}, errors.Errorf("nCan't read environment variables: %v", err)
	}
	dbConfig := dbConfig{
		dbms:         os.Getenv("DBMS"),
		databaseName: os.Getenv("DATABASE_NAME"),
		databaseHost: os.Getenv("DATABASE_HOST"),
		databasePort: os.Getenv("DATABASE_PORT"),
		username:     os.Getenv("DATABASE_USERNAME"),
		password:     os.Getenv("DATABASE_PASSWORD"),
		protocol:     os.Getenv("DATABASE_PROTOCOL"),
	}
	return dbConfig, nil
}
