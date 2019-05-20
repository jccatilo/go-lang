package api

import (
	"net/http"
	"strings"

	// Specify the username that you used inside github.com folder
	"github.com/jccatilo/inventory_rnd/models"
	"github.com/uadmin/uadmin"
)

// CustomListHandler !
func CustomListHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/custom_list")

	// Fetch Data from DB
	stock := []models.Stock{}

	// Assigns a map as a string of interface to store any types of values
	results := []map[string]interface{}{}

	// "id" - order the todo model by id
	// false - to sort in descending order
	// 0 - start at index 0
	// 5 - get five records
	// &todo - todo model to execute
	// "" - fetch the id of the model itself
	uadmin.AdminPage("id", false, 0, 2, &stock, "")

	// Loop to fetch the record of todo
	for i := range stock {
		// Accesses and fetches the record of the linking models in Todo
		uadmin.Preload(&stock[i])

		// Assigns the string of interface in each Todo fields
		results = append(results, map[string]interface{}{
			//"ID":          stock[i].ID,
			"Name":           stock[i].Name,
			"Specifications": stock[i].Specifications,
			// This returns only the name of the Category model, not the
			// other fields
			"Quantity": stock[i].Quantity,
			// This returns only the name of the Friend model, not the
			// other fields
			"UOM": stock[i].UOM.Name,
			// This returns only the name of the Item model, not the other
			// fields
			"Supplier":   stock[i].Supplier.Name,
			"Allocation": stock[i].Allocation.Name,
		})
	}

	// Prints the results in JSON format
	uadmin.ReturnJSON(w, r, results)
}
