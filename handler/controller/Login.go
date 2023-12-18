package controller

import (
	helper "github.com/aiteung/athelper/fiber"
	val "github.com/aiteung/athelper/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/rofinafiin/androidapi/handler/models"
	"github.com/rofinafiin/androidapi/handler/repository"
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

func (log *LoginHandler) Login(ctx *fiber.Ctx) (err error) {
	req := new(models.Login)
	if err = val.ParseAndValidatePlayGround(ctx.Body(), req); err != nil {
		return err
	}

	cihuy := log.Trx.Login(ctx.Context(), *req)
	if !cihuy {
		return fiber.NewError(fiber.StatusBadRequest, "Password salah")
	}

	err = helper.ReturnData[*models.Login]{
		Code:    fiber.StatusOK,
		Success: cihuy,
		Status:  "Berhasil Login",
		Data:    req,
	}.WriteResponseBody(ctx)
	return
}
