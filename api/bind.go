package api

import (
	"github.com/driver005/gateway/sql"
	"github.com/driver005/gateway/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func Bind[T any](context fiber.Ctx, validator *validator.Validate) (*T, *utils.ApplictaionError) {
	var model T
	if err := context.Bind().Body(&model); err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid query parameters",
			nil,
		)
	}

	if err := validator.Struct(&model); err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			err.Error(),
			nil,
		)
	}

	return &model, nil
}

func BindGet(context fiber.Ctx, name string) (uuid.UUID, *sql.Options, *utils.ApplictaionError) {
	id, err := utils.ParseUUID(context.Params(name))
	if err != nil {
		return uuid.Nil, nil, err
	}

	var config sql.Options
	if err := context.Bind().Query(&config); err != nil {
		return uuid.Nil, nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid query parameters",
			nil,
		)
	}

	return id, &config, nil
}

func BindDelete(context fiber.Ctx, name string) (uuid.UUID, *utils.ApplictaionError) {
	id, err := utils.ParseUUID(context.Params(name))
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func BindList[T any](context fiber.Ctx) (*T, *sql.Options, *utils.ApplictaionError) {
	var model T
	if err := context.Bind().Query(&model); err != nil {
		return nil, nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid query parameters",
			nil,
		)
	}

	var config sql.Options
	if err := context.Bind().Query(&config); err != nil {
		return nil, nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid query parameters",
			nil,
		)
	}

	return &model, &config, nil
}

func BindCreate[T any](context fiber.Ctx, validator *validator.Validate) (*T, *utils.ApplictaionError) {
	var model T
	if err := context.Bind().Body(&model); err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid query parameters",
			nil,
		)
	}

	if err := validator.Struct(&model); err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			err.Error(),
			nil,
		)
	}

	return &model, nil
}

func BindUpdate[T any](context fiber.Ctx, name string, validator *validator.Validate) (*T, uuid.UUID, *utils.ApplictaionError) {
	id, err := utils.ParseUUID(context.Params(name))
	if err != nil {
		return nil, uuid.Nil, err
	}

	var model T
	if err := context.Bind().Body(&model); err != nil {
		return nil, uuid.Nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid query parameters",
			nil,
		)
	}

	if err := validator.Struct(&model); err != nil {
		return nil, uuid.Nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			err.Error(),
			nil,
		)
	}

	return &model, id, nil
}

func BindWithString[T any](context fiber.Ctx, name string, validator *validator.Validate) (*T, string, *utils.ApplictaionError) {
	value := context.Params(name)

	var model T
	if err := context.Bind().Body(&model); err != nil {
		return nil, "", utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid query parameters",
			nil,
		)
	}

	if err := validator.Struct(&model); err != nil {
		return nil, "", utils.NewApplictaionError(
			utils.INVALID_DATA,
			err.Error(),
			nil,
		)
	}

	return &model, value, nil
}

func BindAll[T any](context fiber.Ctx, name string, validator *validator.Validate) (*T, uuid.UUID, *sql.Options, *utils.ApplictaionError) {
	id, err := utils.ParseUUID(context.Params(name))
	if err != nil {
		return nil, uuid.Nil, nil, err
	}

	var model T
	if err := context.Bind().Query(&model); err != nil {
		return nil, uuid.Nil, nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid query parameters",
			nil,
		)
	}

	if err := validator.Struct(&model); err != nil {
		return nil, uuid.Nil, nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			err.Error(),
			nil,
		)
	}

	var config sql.Options
	if err := context.Bind().Query(&config); err != nil {
		return nil, uuid.Nil, nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid query parameters",
			nil,
		)
	}

	return &model, id, &config, nil
}

func BindWithUUID[T any](context fiber.Ctx, name string, validator *validator.Validate) (*T, uuid.UUID, *utils.ApplictaionError) {
	id, err := utils.ParseUUID(context.Params(name))
	if err != nil {
		return nil, uuid.Nil, err
	}

	var model T
	if err := context.Bind().Body(&model); err != nil {
		return nil, uuid.Nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid query parameters",
			nil,
		)
	}

	if err := validator.Struct(&model); err != nil {
		return nil, uuid.Nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			err.Error(),
			nil,
		)
	}

	return &model, id, nil
}

func BindConfig(context fiber.Ctx, validator *validator.Validate) (*sql.Options, *utils.ApplictaionError) {
	var config sql.Options
	if err := context.Bind().Query(&config); err != nil {
		return nil, utils.NewApplictaionError(
			utils.INVALID_DATA,
			"Invalid query parameters",
			nil,
		)
	}

	return &config, nil
}
