package main

import (
	"log"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rofinafiin/androidapi/config"
	"github.com/rofinafiin/androidapi/url"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2"
)

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	site.Use(logger.New(logger.Config{
		Format: "${status} - ${method} ${path}\n",
	}))
	url.Web(site)
	log.Fatal(site.Listen(":3000"))

}
