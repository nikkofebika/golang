package users

import (
	"learn/fiber/app/modules/users/dto"
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
	service UserServiceInterface
}

func NewUserController(service UserServiceInterface) *userController {
	return &userController{service}
}

func (controller *userController) Index(c *fiber.Ctx) error {
	users, err := controller.service.FindAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.JSON(users)
}

func (controller *userController) Show(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user, err := controller.service.FindById(uint(id))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.JSON(user)
}

func (controller *userController) Create(c *fiber.Ctx) error {
	var userRequest dto.UserCreateRequest
	if err := c.BodyParser(userRequest); err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{"error": err})
	}

	if err := controller.service.Create(userRequest); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.JSON(fiber.Map{"data": fiber.Map{"message": "Data created successfuly"}})
}
