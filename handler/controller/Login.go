package controller

import (
	helper "github.com/aiteung/athelper/fiber"
	val "github.com/aiteung/athelper/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/rofinafiin/androidapi/handler/models"
	"github.com/rofinafiin/androidapi/handler/repository"
	helper2 "github.com/rofinafiin/androidapi/helper"
)

type LoginHandler struct {
	Trx *repository.LoginTable
}

func (log *LoginHandler) Register(ctx *fiber.Ctx) (err error) {
	req := new(models.Login)
	if err = val.ParseAndValidatePlayGround(ctx.Body(), req); err != nil {
		return err
	}

	err = log.Trx.Register(ctx.Context(), *req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Gagal Register data")
	}

	err = helper.ReturnData[*models.Login]{
		Code:    fiber.StatusOK,
		Success: true,
		Status:  "Berhasil Registrasi",
		Data:    req,
	}.WriteResponseBody(ctx)
	return
}

func (log *LoginHandler) GetDataUser(ctx *fiber.Ctx) (err error) {
	data, err := log.Trx.GetUser(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Gagal Register data")
	}

	err = helper.ReturnData[[]models.Login]{
		Code:    fiber.StatusOK,
		Success: true,
		Status:  "Berhasil Registrasi",
		Data:    data,
	}.WriteResponseBody(ctx)
	return
}

func (log *LoginHandler) Login(ctx *fiber.Ctx) (err error) {
	req := new(models.Login)
	if err = val.ParseAndValidatePlayGround(ctx.Body(), req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.LoginResponse{
			Token:   "-",
			Message: err.Error(),
			Status:  fiber.StatusBadRequest,
		})
	}

	cihuy := log.Trx.Login(ctx.Context(), *req)
	if !cihuy {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.LoginResponse{
			Token:   "-",
			Message: "password salah",
			Status:  fiber.StatusBadRequest,
		})
	}

	randstring := helper2.GenerateRandomString(15)

	err = ctx.Status(fiber.StatusOK).JSON(models.LoginResponse{
		Token:   randstring,
		Message: "Behasil Login",
		Status:  fiber.StatusOK,
	})
	return
}
