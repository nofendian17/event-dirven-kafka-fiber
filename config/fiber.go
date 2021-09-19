package config

import (
	"github.com/gofiber/fiber/v2"
	"kafka-fibre/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}
