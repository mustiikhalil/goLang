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

var mainCL = make(chan int)
var doneCL = make(chan struct{})

func selectStatment() {
	for {
		select {
		case number := <-mainCL:
			fmt.Println("number: ", number)
		case <-doneCL:
			fmt.Println("done")
			break
		}
	}
}

func testChannels() {
	ch := make(chan int)
	wg.Add(2)
	go func(ch <-chan int) {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		i := 42
		ch <- i
		ch <- 30
		close(ch)
		wg.Done()
	}(ch)

	wg.Wait()
}

func main() {
	// test()
	// testChannels()
	go selectStatment()
	mainCL <- 40
	mainCL <- 20
	doneCL <- struct{}{}
}
