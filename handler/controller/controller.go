package controller

import (
	fibhelp "github.com/aiteung/athelper/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/rofinafiin/androidapi/handler/models"
	"github.com/rofinafiin/androidapi/handler/repository"
)

type MatkulHandler struct {
	Mtk *repository.MatkulTable
}

func (m *MatkulHandler) GetDataMatakuliah(ctx *fiber.Ctx) (err error) {
	datamatkul, err := m.Mtk.GetMatkuldata(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Data tidak ada "+err.Error())
	}
	err = fibhelp.ReturnData[[]models.Matakuliah]{
		fiber.StatusOK,
		true,
		"Data berhasil diambil",
		datamatkul,
	}.WriteResponseBody(ctx)
	return
}

func (m *MatkulHandler) InsertDataMatakuliah(ctx *fiber.Ctx) (err error) {
	req := new(models.Matakuliah)
	err = ctx.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot retrieve the body")
	}
	err = m.Mtk.InsertMatkuldata(ctx.Context(), req)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Gagal Insert data "+err.Error())
	}
	err = fibhelp.ReturnData[string]{
		fiber.StatusOK,
		true,
		"Data berhasil diinput",
		"data diinputkan dengan kode mk " + req.KodeMK,
	}.WriteResponseBody(ctx)
	return
}

func (m *MatkulHandler) UpdateDataMatakuliah(ctx *fiber.Ctx) (err error) {
	req := new(models.Matakuliah)
	err = ctx.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot retrieve the body")
	}
	err = m.Mtk.UpdateMatkuldata(ctx.Context(), req)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Gagal Insert data "+err.Error())
	}
	err = fibhelp.ReturnData[string]{
		fiber.StatusOK,
		true,
		"Data berhasil Diupdate",
		"data Diupdate dengan kode mk " + req.KodeMK,
	}.WriteResponseBody(ctx)
	return
}

func (m *MatkulHandler) DeleteDataMatakuliah(ctx *fiber.Ctx) (err error) {
	req := new(models.Request)
	err = ctx.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot retrieve the body")
	}
	deleted, err := m.Mtk.DeleteMatakuliah(ctx.Context(), req.KodeMK)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Gagal Insert data "+err.Error())
	}
	err = fibhelp.ReturnData[models.Matakuliah]{
		fiber.StatusOK,
		true,
		"Data berhasil Dihapus",
		deleted,
	}.WriteResponseBody(ctx)
	return
}
