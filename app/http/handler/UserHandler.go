package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gofiber-orm/domain/service"
)

type UserHandler struct {
	userSvc service.UserService
	logger  fiber.Handler
}

func NewUserHandler(userSvc service.UserService, logger fiber.Handler) *UserHandler {
	return &UserHandler{
		userSvc: userSvc,
		logger:  logger,
	}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	userId, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user ID")
	}

	user, err := h.userSvc.GetUser(userId)
	if err != nil {

		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return c.Render("user", user)
}

//resp := model.UserResponse{
//		ID:           user.ID,
//		Username:     user.Username,
//		Name:         user.Name,
//		Nickname:     user.Nickname,
//		AvatarID:     user.AvatarID,
//		PhoneNumber:  user.PhoneNumber,
//		StudentYear:  user.StudentYear,
//		StudentGroup: user.StudentGroup,
//		MajorCode:    user.MajorCode,
//	}
//
//	return c.JSON(resp)
