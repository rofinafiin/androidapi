package controller

import (
	fibhelp "github.com/aiteung/athelper/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/rofinafiin/androidapi/handler/repository"
	"strconv"
)

// Complete Code about skeleton github.com/rofinafiin

type TransaksiHandler struct {
	Trx *repository.TransaksiTable
}

// Read Data
func (trx *TransaksiHandler) GetDataTransaksi(ctx *fiber.Ctx) (err error) {
	data, err := trx.Trx.GetTransaksiData(ctx.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Data Not Found "+err.Error())
	}

	err = fibhelp.ReturnData[[]repository.TransaksiPembelian]{
		fiber.StatusOK,
		true,
		"Data berhasil diambil",
		data,
	}.WriteResponseBody(ctx)

	return
}

// Create Data
func (trx *TransaksiHandler) InsertTransaksi(ctx *fiber.Ctx) (err error) {
	req := new(repository.TransaksiPembelian)
	err = ctx.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot retrieve the body")
	}
	err = trx.Trx.InsertTransaksi(ctx.Context(), req)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Gagal Insert data "+err.Error())
	}
	err = fibhelp.ReturnData[string]{
		fiber.StatusOK,
		true,
		"Data berhasil diinput",
		"data diinputkan dengan nomor faktur " + strconv.Itoa(req.NomorFaktur),
	}.WriteResponseBody(ctx)
	return
}

// Update Data
func (trx *TransaksiHandler) UpdateTransaksi(ctx *fiber.Ctx) (err error) {
	req := new(repository.TransaksiPembelian)
	err = ctx.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot retrieve the body")
	}
	err = trx.Trx.UpdateTransaksi(ctx.Context(), req)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Gagal Insert data "+err.Error())
	}
	err = fibhelp.ReturnData[string]{
		fiber.StatusOK,
		true,
		"Data berhasil Diupdate",
		"data Diupdate dengan Nomor Faktur " + strconv.Itoa(req.NomorFaktur),
	}.WriteResponseBody(ctx)
	return
}

// Delete Data
func (trx *TransaksiHandler) DeleteTransaksi(ctx *fiber.Ctx) (err error) {
	req := new(repository.RequestTrx)
	err = ctx.BodyParser(req)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Cannot retrieve the body")
	}
	deleted, err := trx.Trx.DeleteTransaksi(ctx.Context(), req.NomorFaktur)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Gagal Hapus data "+err.Error())
	}
	err = fibhelp.ReturnData[repository.TransaksiPembelian]{
		fiber.StatusOK,
		true,
		"Data berhasil Dihapus",
		deleted,
	}.WriteResponseBody(ctx)
	return
}
