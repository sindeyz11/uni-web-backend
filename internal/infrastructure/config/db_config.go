package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type DatabaseConfig struct {
	Driver   string
	Username string
	Password string
	Host     string
	Port     string
	Database string
}

type ParseConfig interface {
	MySqlConfig() *DatabaseConfig
	PostgresConfig() *DatabaseConfig
}

type Config struct{}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) MySqlConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Driver:   os.Getenv("MYSQL_DRIVER"),
		Username: os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		Database: os.Getenv("MYSQL_DATABASE"),
	}
}

func (c *Config) PostgresConfig() *DatabaseConfig {
	return &DatabaseConfig{
		Driver:   os.Getenv("PS_DRIVER"),
		Username: os.Getenv("PS_USER"),
		Password: os.Getenv("PS_PASSWORD"),
		Port:     os.Getenv("PS_PORT"),
		Database: os.Getenv("PS_DB"),
	}
}

func NewMysqlConn(c *DatabaseConfig) *sql.DB {
	conn, err := sql.Open(c.Driver, fmt.Sprintf(
		"%s:%s@tcp(mysql_db:%s)/%s",
		c.Username,
		c.Password,
		c.Port,
		c.Database))
	if err != nil {
		log.Fatalf("Can't open database connection, %v", err)
		return nil
	}
	if err := conn.Ping(); err != nil {
		log.Fatalf("Can't open database connection, %v", err)
		return nil
	}
	return conn
}
