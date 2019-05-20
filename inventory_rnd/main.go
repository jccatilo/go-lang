package main

import (
	"net/http"

	"github.com/jccatilo/inventory_rnd/api"
	"github.com/jccatilo/inventory_rnd/models"
	"github.com/uadmin/uadmin"
)

func main() {
	uadmin.RootURL = "/admin/"
	uadmin.Register(
		models.Stock{},
		models.Uom{},
		models.Supplier{},
		models.Allocation{},
	)

	uadmin.Port = 8888
	//api handler
	http.HandleFunc("/api/", api.APIHandler)
	uadmin.StartServer()

}
