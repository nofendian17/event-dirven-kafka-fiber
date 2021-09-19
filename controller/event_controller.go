package controller

import (
	"github.com/gofiber/fiber/v2"
	"kafka-fibre/exception"
	"kafka-fibre/model"
	"kafka-fibre/service"
	"kafka-fibre/validation"
)

type EventController struct {
	EventService service.ProduceService
}

func NewEventController(eventService *service.ProduceService) EventController {
	return EventController{
		EventService: *eventService,
	}
}

func (controller *EventController) Route(app *fiber.App) {
	app.Post("/api/event", controller.Create)
}

func (controller *EventController) Create(c *fiber.Ctx) error {

	requestBody := model.EventRequest{}

	err := c.BodyParser(&requestBody)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	validation.ValidateCreateEvent(requestBody)

	err = controller.EventService.Create(requestBody)
	if err != nil {
		return exception.ErrorHandler(c, err)
	}

	c.Status(201)
	return nil
}
