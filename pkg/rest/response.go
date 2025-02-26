package rest

import "github.com/gofiber/fiber/v2"

type SearchResponse struct {
	Entities   any   `json:"entities"`
	TotalCount int64 `json:"totalCount"`
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
	TotalPages int   `json:"totalPages"`
}

func Success(c *fiber.Ctx, data interface{}) error {
	return c.JSON(fiber.Map{
		"success": true,
		"data":    data,
	})
}

func Error(err error) error {
	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
}

func BadRequest(err error) error {
	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
}
