package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func throwsError(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatal(err)
	}
}

func handleBuscaCep(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	getCepURL := func(cepCode string) string {
		const BuscaCepURL = "http://apps.widenet.com.br/busca-cep/api/cep/"
		return BuscaCepURL + cepCode + ".json"
	}

	cepToSearch := p.ByName("cep")
	fmt.Println("Handle json called with", cepToSearch)

	cepURL := getCepURL(cepToSearch)
	resp, err := http.Get(cepURL)

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println("Got response", string(body))
	throwsError(err, w)

	defer resp.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

func handleJSON(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// profile := map[string]string{"email": "Alex", "senha": "teste de senha"}
	profile := Profile{"eai", "issae"}
	js, err := json.Marshal(profile)

	fmt.Println("Handle json called with", string(js))
	throwsError(err, w)
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}
