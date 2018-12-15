package main

import (
	"net/http"
	"fmt"
)

func main () {
	fmt.Println("subiu")
	router := spawnRoutes()
	http.ListenAndServe(":8080", router)
}
