package main

import (
	"github.com/julienschmidt/httprouter"
)

func spawnRoutes() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", handleJSON)
	router.GET("/buscacep/:cep", handleBuscaCep)

	return router
}
