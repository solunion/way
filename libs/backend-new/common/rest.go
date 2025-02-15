package common

import (
	"github.com/gofiber/fiber/v3"
)

type Rest[T any] interface {
	Create(ctx fiber.Ctx) error
	GetAll(ctx fiber.Ctx) error
	//GetById(ctx fiber.Ctx) error
}
