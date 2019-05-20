package models

import (
	"github.com/uadmin/uadmin"
)

type Stock struct {
	uadmin.Model
	Name           string `uadmin:"list_exclude"`
	Specifications string
	Quantity       string
	UOM            Uom
	UOMID          uint
	Supplier       Supplier
	SupplierID     uint
	Allocation     Allocation
	AllocationID   uint
}
