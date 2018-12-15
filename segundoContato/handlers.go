package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
)

func handleJson (w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// profile := map[string]string{"email": "Alex", "senha": "teste de senha"}
	profile := Profile{ "eai", "issae" }
	js, err := json.Marshal(profile)

	fmt.Println("Chamou", string(js))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}