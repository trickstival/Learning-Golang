package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func chann() {
	channel := make(chan string)
	wg.Add(2)

	go func() {
		str := <-channel
		fmt.Println(str)
		wg.Done()
	}()

	go func() {
		channel <- "eai po"
		wg.Done()
	}()

	wg.Wait()
}
