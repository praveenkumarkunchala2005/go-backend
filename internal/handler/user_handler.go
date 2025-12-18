package handler

import (
	"time"

	"backend-task/internal/models"
	"backend-task/internal/service"

	"github.com/go-playground/validator/v10"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	service  *service.UserService
	validate *validator.Validate
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{
		service:  s,
		validate: validator.New(),
	}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.service.CreateUser(c.Context(), req.Name, dob)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(201).JSON(user)
}

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id <= 0 {
		return fiber.ErrBadRequest
	}

	user, err := h.service.GetUserByID(c.Context(), int64(id))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(200).JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id <= 0 {
		return fiber.ErrBadRequest
	}

	if err := h.service.DeleteUser(c.Context(), int64(id)); err != nil {
		return fiber.ErrInternalServerError
	}

	return c.SendStatus(204)
}

func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.service.ListUsers(c.Context())
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(200).JSON(users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil || id <= 0 {
		return fiber.ErrBadRequest
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	req.ID = int64(id)

	if err := h.validate.Struct(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	dob, err := time.Parse("2006-01-02", req.DOB)
	if err != nil {
		return fiber.ErrBadRequest
	}

	user, err := h.service.UpdateUser(c.Context(), req.ID, req.Name, dob)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.Status(200).JSON(user)
}

func (h *UserHandler) ListUsersPaginated(c *fiber.Ctx) error {
	page := c.QueryInt("page", 1)
	limit := c.QueryInt("limit", 10)

	users, total, err := h.service.ListUsersPaginated(
		c.Context(),
		int64(page),
		int64(limit),
	)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if users == nil {
		users = make([]models.User, 0)
	}

	return c.JSON(fiber.Map{
		"page":  page,
		"limit": limit,
		"total": total,
		"data":  users,
	})
}
