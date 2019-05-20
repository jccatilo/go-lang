package models

import (
	"github.com/uadmin/uadmin"
)

type Supplier struct {
	uadmin.Model
	Name string `uadmin:"list_exclude"`
}
