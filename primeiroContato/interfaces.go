package main

import "fmt"

type Pessoa struct {
	nome   string
	idade  int
	altura float64
}

type Pet interface {
	call()
}

type Cachorro struct {
	raca  string
	porte string
}

func (cachorro Cachorro) call() {
	fmt.Println("Au au")
}

type Gato struct {
	miado         string
	comportamento string
}

func (gato Gato) call() {
	fmt.Println("Meow")
}

func main2() {
	// nums := [...]int{1, 2, 3, 4}

	defer (func() {
		fmt.Println("hehe")
	})()

	mapaIdades := make(map[string]int)

	nomePatrick := "Patrick"
	idadePatrick := 17

	mapaIdades[nomePatrick] = idadePatrick

	pessoa := Pessoa{nomePatrick, idadePatrick, 1.9}

	fmt.Println(pessoa.idadeVezesAltura())
}

func (pessoa *Pessoa) idadeVezesAltura() float64 {
	return float64(pessoa.idade) * pessoa.altura
}

func addThemUp(nums []float64) float64 {
	sum := .0

	for _, val := range nums {
		sum += val
	}

	return sum
}
