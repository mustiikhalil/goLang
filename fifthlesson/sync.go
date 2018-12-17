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
	wg.Add(1)
	go safelySay("hello")
	wg.Add(1)
	go safelySay("there")
	wg.Wait()
}

func testChannels() {
	ch := make(chan int)
	go func() {
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}()
	go func() {
		i := 42
		ch <- i
		i = 30
		fmt.Println(i)
		wg.Done()
	}()

	wg.Wait()
}

func main() {
	// test()
	testChannels()
}
