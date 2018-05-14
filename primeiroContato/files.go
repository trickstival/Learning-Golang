package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func files() {
	fmt.Println("Hello world!")

	file, err := os.Create("file.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Faaala galerinha do youtube, beleza?")
	file.Close()

	stream, err := ioutil.ReadFile("file.txt")

	if err != nil {
		log.Fatal(err)
	}

	readString := string(stream)

	fmt.Println(readString)

}
