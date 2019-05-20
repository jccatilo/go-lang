package api

import (
	"net/http"
	"strings"

	"github.com/jccatilo/inventory_rnd/models"
	"github.com/uadmin/uadmin"
)

func TodoListHandler(w http.ResponseWriter, r *http.Request) {
	// r.URL.Path creates a new path called /todo_list
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/todolist")

	// Fetches all object in the database
	stock := []models.Stock{}
	uadmin.All(&stock)

	// Accesses and fetches data from another model
	for t := range stock {
		uadmin.Preload(&stock[t])
	}

	// Prints the todo in JSON format
	uadmin.ReturnJSON(w, r, stock)
}
