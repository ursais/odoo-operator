package controller

import (
	"github.com/ursais/odoo-operator/pkg/controller/odoo"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, odoo.Add)
}
