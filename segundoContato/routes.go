package main

import (
	"github.com/julienschmidt/httprouter"
)

func spawnRoutes () *httprouter.Router {
	router := httprouter.New()
	router.GET("/", handleJson)

	return router
}
