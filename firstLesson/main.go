package main

import (
	"fmt"
	"math/rand"
	"time"
)

const pi = 3.14

//Add This function will be accessable throught the entire projects since it started with a capital letter
func Add(x, y float64) float64 {
	return x + y
}

//This function returns two strings and is only visible here
func multiple(a, b string) (string, string) {
	return a, b
}

func main() {

	a, b := 9.4, 6.5
	var string1, string2 string = "hey", "there"

	fmt.Println(multiple(string1, string2))
	fmt.Println(Add(a, b))

	rand.Seed(time.Now().UnixNano())
	fmt.Println("random number between 0 - 100: ", rand.Intn(100))

	x := &a // refrence to the memory address
	fmt.Println(x)
	fmt.Println(*x)
}
