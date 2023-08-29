package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chon26909/e-commerce/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type ILogger interface {
	Print() ILogger
	Save()
	SetQuery(c *fiber.Ctx)
	SetBody(c *fiber.Ctx)
	SetResponse(c any)
}

type logger struct {
	Time       string `json:"time"`
	Ip         string `json:"ip"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

func NewLogger(c *fiber.Ctx, res any) ILogger {
	return &logger{
		Time:       time.Now().Local().Format("2006-01-02 15:04:05"),
		Ip:         c.IP(),
		Method:     c.Method(),
		StatusCode: c.Response().StatusCode(),
		Path:       c.Path(),
	}
}

// Print implements ILogger.
func (l *logger) Print() ILogger {
	utils.Debug(l)
	return l
}

// Save implements ILogger.
func (l *logger) Save() {

	data := utils.Output(l)

	filename := fmt.Sprintf("./assets/logs/logger_%v.txt", strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", ""))

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	file.WriteString(string(data) + "")

}

// SetBody implements ILogger.
func (l *logger) SetBody(c *fiber.Ctx) {
	var body any
	if err := c.BodyParser(&body); err != nil {
		log.Printf("bodyParser error: %v", err)
	}

	switch l.Path {
	case "v1/users/signup":
		l.Body = "[masked]"
	default:
		l.Body = body
	}
}

// SetQuery implements ILogger.
func (l *logger) SetQuery(c *fiber.Ctx) {
	var query any
	if err := c.QueryParser(&query); err != nil {
		log.Printf("bodyParser error: %v", err)
	}

	l.Query = query
}

// SetResponse implements ILogger.
func (l *logger) SetResponse(res any) {
	l.Response = res
}
