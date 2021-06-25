package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/pranaySinghDev/ProfessorChaos/util"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		msg, err := util.CallHTTP(os.Getenv("url"))
		if err != nil {
			log.Print("Error in service 3", err)
			c.Status(500).SendString(err.Error())
		}
		return c.SendString(fmt.Sprintf("sevice3 --> %s", msg))
	})
	app.Listen(fmt.Sprintf(":%d", 7002))
}
