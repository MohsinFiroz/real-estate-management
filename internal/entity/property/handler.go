package property

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Create(c *fiber.Ctx) error {
	// Handler implementation
	return nil
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	// Handler implementation
	return nil
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	// Handler implementation
	return nil
}
