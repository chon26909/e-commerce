package server

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"

	"github.com/chon26909/e-commerce/config"
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type server struct {
	app    *fiber.App
	db     *sqlx.DB
	config config.IConfig
}

type IServer interface {
	Start()
}

func NewServer(config config.IConfig, db *sqlx.DB) IServer {
	return &server{
		config: config,
		db:     db,
		app: fiber.New(fiber.Config{
			AppName:      config.App().Name(),
			BodyLimit:    config.App().BodyLimit(),
			ReadTimeout:  config.App().ReadTimeout(),
			WriteTimeout: config.App().WriteTimeout(),
			JSONEncoder:  json.Marshal,
			JSONDecoder:  json.Unmarshal,
		}),
	}
}

func (s *server) Start() {

	// Middleware
	middlewares := NewMiddleware(s)
	s.app.Use(middlewares.Cors())

	// Modules
	v1 := s.app.Group("v1")
	modules := NewModule(v1, s)

	modules.MonitorModule()

	s.app.Use(middlewares.RouterCheck())

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		log.Println("server is shutting down")
		_ = s.app.Shutdown()
	}()

	log.Printf("server is starting on %v", s.config.App().Url())

	s.app.Listen(s.config.App().Url())

}
