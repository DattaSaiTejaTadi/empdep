package services

import (
	"github.com/LetsFocus/account-service/empdep/store"
)

type app struct {
	store store.Store
}

func New(s store.Store) *app {
	return &app{
		store: s,
	}
}
