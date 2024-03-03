package handler

import (
	"refresh/internal/user/service"
	"refresh/internal/user/dto/request"
	"refresh/pkg/auth"
	"refresh/pkg/helper"

	"strings"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	service service.UserServiceInterface
	auth   auth.Token
}

func NewUserhandler(service service.UserServiceInterface,auth   auth.Token) *handler {
	return &handler{service,auth}
}

func (user *handler) Register(c *fiber.Ctx) error {
	input := request.UserRequest{}

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(err.Error()))
	}

	err = user.service.InsertUser(input)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(err.Error()))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(helper.ErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusCreated).JSON(helper.SuccessResponse("success create user"))
}

func (user *handler) Login(c *fiber.Ctx) error {
	input := request.UserRequest{}

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(err.Error()))
	}

	response, err := user.service.Login(input)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(err.Error()))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(helper.ErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(helper.SuccessWithDataResponse("success login", response))
}

func (user *handler) RefreshToken(c *fiber.Ctx) error {
	input := request.RequestRefresh{}

	err := c.BodyParser(&input)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(err.Error()))
	}

	response, err := user.service.RefreshToken(input.RefreshToken)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(err.Error()))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(helper.ErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(helper.SuccessWithDataResponse("success refresh token", response))
}

func (user *handler) GetProfile(c *fiber.Ctx) error {
	id, err := user.auth.ExtractToken(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(err.Error()))
	}

	response, err := user.service.SelectUserById(id)
	if err != nil {
		if strings.Contains(err.Error(), "error") {
			return c.Status(fiber.StatusBadRequest).JSON(helper.ErrorResponse(err.Error()))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(helper.ErrorResponse(err.Error()))
	}

	return c.Status(fiber.StatusOK).JSON(helper.SuccessWithDataResponse("success get profile", response))
}
