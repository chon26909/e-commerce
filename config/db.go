package config

import "fmt"

type db struct {
	host           string
	port           int
	protocal       string
	username       string
	password       string
	database       string
	sslMode        string
	maxConnections int
}

type IDbConfig interface {
	Url() string
	MaxOpenConnections() int
}

func (c *config) Db() IDbConfig {
	return c.db
}

// Url implements IDbConfig.
func (d *db) Url() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		d.host, d.port, d.username, d.password, d.database, d.sslMode,
	)
}

// MaxOpenConnections implements IDbConfig.
func (d *db) MaxOpenConnections() int {
	return d.maxConnections
}
