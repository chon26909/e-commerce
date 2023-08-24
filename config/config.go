package config

import (
	"os"
	"time"
)

func envPath() string {
	if len(os.Args) > 0 {
		return ".env"
	} else {
		return os.Args[1]
	}
}

type IConfig interface {
	App()
	DB()
	JWT()
}

type config struct {
	app *app
	db  *db
	jwt *jwt
}

type app struct {
	host         string
	port         uint
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int
	fileLimit    int
	gcpBucket    string
}

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

type jwt struct {
	adminKey         string
	secertKey        string
	apiKey           string
	accessExpiresAt  int
	refreshExpiresAt int
}
