package handler

import (
	"strconv"

	"github.com/alfariiizi/go-echo-fx-template/internal/core/domain"
	serviceport "github.com/alfariiizi/go-echo-fx-template/internal/core/service/port"
	httpport "github.com/alfariiizi/go-echo-fx-template/internal/delivery/port/http"
)

func GetAllUsersHandler(ctx httpport.HttpContext, userService serviceport.UserService) error {
	users, err := userService.ListUsers()
	if err != nil {
		return ctx.SendErrorResponse(500, "Failed to retrieve users", nil)
	}

	return ctx.SendSuccessResponse(200, users)
}

func GetUserHandler(ctx httpport.HttpContext, userService serviceport.UserService) error {
	id, err := ctx.GetParam("id", true)
	if err != nil {
		return err
	}
	// Convert id to uint
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return ctx.SendErrorResponse(400, "Invalid user ID", nil)
	}

	user, err := userService.GetUser(uint(idUint))
	if err != nil {
		return ctx.SendErrorResponse(404, "User not found", nil)
	}

	return ctx.SendSuccessResponse(200, user)
}

func CreateUserHandler(ctx httpport.HttpContext, userService serviceport.UserService) error {
	var userRequest domain.UserRequest

	if err := ctx.BindBody(&userRequest); err != nil {
		return err
	}

	user, err := userService.CreateUser(userRequest)
	if err != nil {
		return ctx.SendErrorResponse(500, "Failed to create user", nil)
	}

	return ctx.SendSuccessResponse(201, user)
}

func UpdateUserHandler(ctx httpport.HttpContext, userService serviceport.UserService) error {
	id, err := ctx.GetParam("id", true)
	if err != nil {
		return err
	}

	var userRequest domain.UserRequest
	if err := ctx.BindBody(&userRequest); err != nil {
		return err
	}

	// Convert id to uint
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return ctx.SendErrorResponse(400, "Invalid user ID", nil)
	}

	user, err := userService.UpdateUser(uint(idUint), userRequest)
	if err != nil {
		return ctx.SendErrorResponse(404, "User not found", nil)
	}

	return ctx.SendSuccessResponse(200, user)
}

func DeleteUserHandler(ctx httpport.HttpContext, userService serviceport.UserService) error {
	id, err := ctx.GetParam("id", true)
	if err != nil {
		return err
	}

	// Convert id to uint
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return ctx.SendErrorResponse(400, "Invalid user ID", nil)
	}

	err = userService.DeleteUser(uint(idUint))
	if err != nil {
		return ctx.SendErrorResponse(404, "User not found", nil)
	}

	return ctx.SendSuccessResponse(200, nil)
}
