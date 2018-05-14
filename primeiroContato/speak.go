package main

import (
	"fmt"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}

func say(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str)
		time.Sleep(time.Millisecond * 100)
	}
	waitGroup.Done()
}

func speak() {
	waitGroup.Add(2)
	go say("There")
	go say("Hey")
	waitGroup.Wait()
}
