package config

import "github.com/gofiber/fiber/v2"

var Iteung = fiber.Config{
	Prefork:       true,
	CaseSensitive: true,
	StrictRouting: true,
	ServerHeader:  "Iteung",
	AppName:       "Message Router",
}

var Stringmaria = "root:@tcp(127.0.0.1:3306)/matakuliah_api"
