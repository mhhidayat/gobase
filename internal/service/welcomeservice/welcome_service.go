package welcomeservice

import (
	"gobase/internal/serviceinterface/welcomeserviceinterface"
	"log"
)

type WelcomeService struct {
}

func NewWelcomeService() welcomeserviceinterface.WelcomeServiceInterface {
	return &WelcomeService{}
}

func (ws *WelcomeService) Welcome() {
	log.Println("Welcome Service")
}
