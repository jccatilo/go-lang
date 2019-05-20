package models

import (
	"github.com/uadmin/uadmin"
)

type Allocation struct {
	uadmin.Model
	Name string `uadmin:"list_exclude"`
}
