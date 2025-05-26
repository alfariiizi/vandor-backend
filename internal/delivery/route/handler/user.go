package handler

import (
	httpctx "github.com/alfariiizi/go-service/internal/delivery/route/context"
	"github.com/alfariiizi/go-service/internal/model"
	"github.com/alfariiizi/go-service/internal/service/user"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) GetAllUsers(ctx httpctx.HttpContext) error {
	users, err := u.userService.ListUsers()
	if err != nil {
		return ctx.SendErrorResponse(500, "Failed get all users", err)
	}

	return ctx.SendSuccessResponse(200, users)
}

func (u *UserHandler) GetUserByID(ctx httpctx.HttpContext) error {
	id, err := ctx.GetParam("id", true)
	if err != nil {
		return ctx.SendErrorResponse(400, "Need ID", err)
	}
	if id == "" {
		return ctx.SendErrorResponse(400, "ID is required", nil)
	}

	user, err := u.userService.GetUser(id)
	if err != nil {
		return ctx.SendErrorResponse(500, "Failed get user by id", err)
	}

	return ctx.SendSuccessResponse(200, user)
}

func (u *UserHandler) CreateUser(ctx httpctx.HttpContext) error {
	var userRequest model.UserRequest
	if err := ctx.BindBody(&userRequest); err != nil {
		return ctx.SendErrorResponse(400, "Need Request Body", err)
	}

	user, err := u.userService.CreateUser(userRequest)
	if err != nil {
		return ctx.SendErrorResponse(500, "Failed create user", err)
	}

	return ctx.SendSuccessResponse(201, user)
}

func (u *UserHandler) UpdateUser(ctx httpctx.HttpContext) error {
	id, err := ctx.GetParam("id", true)
	if err != nil {
		return ctx.SendErrorResponse(400, "Need ID", err)
	}
	if id == "" {
		return ctx.SendErrorResponse(400, "ID is required", nil)
	}

	var userRequest model.UserRequest
	if err := ctx.BindBody(&userRequest); err != nil {
		return ctx.SendErrorResponse(400, "Need Request Body", err)
	}

	user, err := u.userService.UpdateUser(id, userRequest)
	if err != nil {
		return ctx.SendErrorResponse(500, "Failed update user", err)
	}

	return ctx.SendSuccessResponse(200, user)
}

func (u *UserHandler) DeleteUser(ctx httpctx.HttpContext) error {
	id, err := ctx.GetParam("id", true)
	if err != nil {
		return ctx.SendErrorResponse(400, "Need ID", err)
	}
	if id == "" {
		return ctx.SendErrorResponse(400, "ID is required", nil)
	}

	err = u.userService.DeleteUser(id)
	if err != nil {
		return ctx.SendErrorResponse(500, "Failed delete user", err)
	}

	return ctx.SendSuccessResponse(200, "User deleted successfully")
}
