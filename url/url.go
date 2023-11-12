package url

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rofinafiin/androidapi/config"
	"github.com/rofinafiin/androidapi/handler/controller"
	"github.com/rofinafiin/androidapi/handler/repository"
)

func Web(page *fiber.App) {
	Mysqlconn := config.CreateMariaGormConnection(config.Stringmaria)
	mtk := repository.NewMatakuliahTable(Mysqlconn)
	Handle := controller.MatkulHandler{mtk}

	grp := page.Group("/mst")

	grp.Get("/data", Handle.GetDataMatakuliah)
	grp.Post("/insert", Handle.InsertDataMatakuliah)
	grp.Put("/update", Handle.UpdateDataMatakuliah)
	grp.Delete("/delete", Handle.DeleteDataMatakuliah)
}
