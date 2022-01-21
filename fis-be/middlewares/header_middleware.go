package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

func CommonMiddleWare(c *fiber.Ctx) error {
	c.Accepts("Content-type", "application/json")
	c.Accepts("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	c.Accepts("Access-Control-Allow-Headers", "Origin, Accept, Content-Type, Content-Length, "+
		"Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, "+
		"Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, "+
		"Cache-Control, X-header")
	return c.Next()
}
