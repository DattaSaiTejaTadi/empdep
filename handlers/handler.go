package handlers

import (
	"github.com/LetsFocus/account-service/empdep/services"
)

type handler struct {
	app services.App
}

func New(a services.App) *handler {
	return &handler{
		app: a,
	}
}
