package gobaseconfiguration

import (
	"gobase/internal/handler/errorhandler"

	"github.com/gofiber/fiber/v2"
)

func GetConfiguration() fiber.Config {
	return fiber.Config{
		ErrorHandler: errorhandler.ErrorHandler,
	}
}
