package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func safelySay(s string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovering from: ", r)
		}
	}()
	say(s)
}

func say(s string) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond * 100)
		if i == 2 {
			panic("Oh dear!")
		}
	}
}

func test() {
	defer fmt.Println("done")
	defer fmt.Println("are we done?")
	fmt.Println("Testing")
}

func main() {
	test()
	wg.Add(1)
	go safelySay("hello")
	wg.Add(1)
	go safelySay("there")
	wg.Wait()
}
