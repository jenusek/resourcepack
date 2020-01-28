package controllers

import (
	"github.com/jenusek/resourcepack/internal/durable"
	"github.com/gorilla/mux"
)

// Register is responsible for registration all systems endpoints
func Register(store durable.Datastore, router *mux.Router) {
	registerResourcesEndpoints(store, router)
	registerAuthenticationEndpoints(store, router)
}
