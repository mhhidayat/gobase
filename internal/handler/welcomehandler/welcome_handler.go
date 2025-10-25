package welcomehandler

import (
	"gobase/internal/serviceinterface/welcomeserviceinterface"
	"gobase/pkg/response"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type welcomeHandler struct {
	service welcomeserviceinterface.WelcomeServiceInterface
}

func NewWelcomeHandler(router fiber.Router, service welcomeserviceinterface.WelcomeServiceInterface) {
	wh := &welcomeHandler{
		service: service,
	}
	router.Get("/", wh.Welcome)
}

func (wh *welcomeHandler) Welcome(c *fiber.Ctx) error {
	wh.service.Welcome()
	return c.Status(http.StatusOK).JSON(
		response.SuccessResponse(
			"Welcome to GoBase",
			nil,
		),
	)
}
