package main

import (
	"fmt"
	"strconv"
	"time"
)

var pizzaNum = 0
var pizzaName = ""

func makeDough(channel chan string) {
	pizzaNum++
	pizzaName = "Pizza #" + strconv.Itoa(pizzaNum)

	fmt.Println("Fazendo massa da", pizzaName, "e enviando para adicionar recheio")
	channel <- pizzaName

	time.Sleep(time.Millisecond * 10)
}

func addSauce(channel chan string) {
	pizza := <-channel

	fmt.Println("Fazendo recheio e enviando a", pizza, "para adicionar borda")

	channel <- pizza

	time.Sleep(time.Millisecond * 10)
}

func addToppings(channel chan string) {
	pizza := <-channel

	fmt.Println("Adicionando borda para a", pizza, "e finalmente entregando")

	time.Sleep(time.Millisecond * 10)
}

func pizza() {
	stringChan := make(chan string)

	for i := 0; i < 3; i++ {
		go makeDough(stringChan)
		go addSauce(stringChan)
		go addToppings(stringChan)

		time.Sleep(time.Millisecond * 5000)
	}
}
