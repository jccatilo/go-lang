package api

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/jccatilo/inventory_rnd/models"
	"github.com/uadmin/uadmin"
)

// AddFriendHandler !
func AddFriendHandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/add_friend")
	res := map[string]interface{}{}

	// Fetch data from Friend DB
	stock := models.Stock{}
	//uom := models.Uom{}
	//supplier := models.Supplier{}
	//allocation := models.Allocation{}

	// Set the parameters of Name, Email, and Password
	stockName := r.FormValue("name")
	stockSpecifications := r.FormValue("specification")
	stockQuantity := r.FormValue("quantity")

	stockUom := r.FormValue("uom")
	stockSupplier := r.FormValue("supplier")
	stockAllocation := r.FormValue("allocation")

	// Validate if the friendName variable is empty.
	if stockName == "" {
		res["status"] = "ERROR"
		res["err_msg"] = "Item Name is required."
		uadmin.ReturnJSON(w, r, res)
		return
	}
	// Store input into the Name, Email, and Password fields
	stock.Name = stockName
	stock.Specifications = stockSpecifications
	stock.Quantity = stockQuantity
	uomid, _ := strconv.ParseInt(stockUom, 10, 64)
	stock.UOMID = uint(uomid)
	supplierid, _ := strconv.ParseInt(stockSupplier, 10, 64)
	stock.SupplierID = uint(supplierid)
	stockallocationid, _ := strconv.ParseInt(stockAllocation, 10, 64)
	stock.AllocationID = uint(stockallocationid)

	// Store input in the Friend model
	uadmin.Save(&stock)

	res["status"] = "ok"
	uadmin.ReturnJSON(w, r, res)
}
