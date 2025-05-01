package handlers

import (
	"github.com/gofiber/fiber/v3"
)

type Rest[T any] interface {
	Create(ctx fiber.Ctx) error
	GetAll(ctx fiber.Ctx) error
	GetById(ctx fiber.Ctx) error
	Update(ctx fiber.Ctx) error
	Delete(ctx fiber.Ctx) error
}
