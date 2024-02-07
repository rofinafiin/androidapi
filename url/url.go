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
	log := repository.NewLogTable(Mysqlconn)
	Handle := controller.TransaksiHandler{Trx: trx}
	loginHandler := controller.LoginHandler{Trx: log}

	grp := page.Group("/mst")

	grp.Get("/data", Handle.GetDataTransaksi)
	grp.Post("/insert", Handle.InsertTransaksi)
	grp.Get("/byid/:nomorfaktur", Handle.GetById)
	grp.Put("/update", Handle.UpdateTransaksi)
	grp.Delete("/delete/:nomorfaktur", Handle.DeleteTransaksi)

	grplg := page.Group("/")
	grplg.Post("register", loginHandler.Register)
	grplg.Post("login", loginHandler.Login)
	grplg.Get("data", loginHandler.GetDataUser)
}
