package models

import (
	"github.com/uadmin/uadmin"
)

type Uom struct {
	uadmin.Model
	Name string `uadmin:"list_exclude"`
}
