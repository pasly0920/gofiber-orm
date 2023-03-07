package logger

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"os"
	"time"
)

func NewLogger() fiber.Handler {
	l := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		if err != nil {
			c.Status(http.StatusInternalServerError)
			c.SendString(err.Error())
			return err
		}
		logMsg := fmt.Sprintf("%s - %s %s - %d", c.IP(), c.Method(), c.Path(), c.Response().StatusCode())
		l.Printf("%s - %v\n", logMsg, time.Since(start))
		return nil
	}
}
