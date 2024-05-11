package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
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
		Driver:   os.Getenv("PGSQL_DRIVER"),
		Username: os.Getenv("PGSQL_USER"),
		Password: os.Getenv("PGSQL_PASSWORD"),
		Port:     os.Getenv("PGSQL_PORT"),
		Database: os.Getenv("PGSQL_DATABASE"),
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

func NewPostgresConn(c *DatabaseConfig) *sql.DB {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@database/%s?sslmode=disable",
		c.Username,
		c.Password,
		c.Database,
	)

	conn, err := sql.Open(c.Driver, connStr)

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
