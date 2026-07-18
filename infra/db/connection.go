package db

import (
	"ecommerce/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnectionString(conf *config.DbConfig) string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.DbHost,
		conf.DbPort,
		conf.DbUser,
		conf.DbPassword,
		conf.DbName,
		conf.DbSSLMode,
	)
}

func NewConnectionString(conf *config.DbConfig) (*sqlx.DB, error) {
	dbSource := GetConnectionString(conf)
	dbConn, err := sqlx.Connect("postgres", dbSource)
	if err != nil {
		return nil, err
	}
	return dbConn, nil
}
