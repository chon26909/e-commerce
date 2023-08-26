package config

import (
	"fmt"
	"time"
)

type app struct {
	host         string
	port         int
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int
	fileLimit    int
	gcpBucket    string
}

type IAppConfig interface {
	Url() string
	Name() string
	Version() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
	FileLimit() int
	GCPBucket() string
}

func (c *config) App() IAppConfig {
	return c.app
}

// Url implements IAppConfig.
func (a *app) Url() string {
	return fmt.Sprintf("%s:%d", a.host, a.port)
}

// Name implements IAppConfig.
func (a *app) Name() string {
	return a.name
}

// Version implements IAppConfig.
func (a *app) Version() string {
	return a.version
}

// ReadTimeout implements IAppConfig.
func (a *app) ReadTimeout() time.Duration {
	return a.readTimeout
}

// WriteTimeout implements IAppConfig.
func (a *app) WriteTimeout() time.Duration {
	return a.writeTimeout
}

// BodyLimit implements IAppConfig.
func (a *app) BodyLimit() int {
	return a.bodyLimit
}

// FileLimit implements IAppConfig.
func (a *app) FileLimit() int {
	return a.fileLimit
}

// GCPBucket implements IAppConfig.
func (a *app) GCPBucket() string {
	return a.gcpBucket
}
