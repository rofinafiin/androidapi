package url

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rofinafiin/androidapi/config"
	"github.com/rofinafiin/androidapi/handler/controller"
	"github.com/rofinafiin/androidapi/handler/repository"
)

func Web(page *fiber.App) {
	Mysqlconn := config.CreateMariaGormConnection(config.Stringmaria)
	trx := repository.NewTransaksiTable(Mysqlconn)
	Handle := controller.TransaksiHandler{Trx: trx}

	grp := page.Group("/mst")

	grp.Get("/data", Handle.GetDataTransaksi)
	grp.Post("/insert", Handle.InsertTransaksi)
	grp.Put("/update", Handle.UpdateTransaksi)
	grp.Delete("/delete", Handle.DeleteTransaksi)
}
