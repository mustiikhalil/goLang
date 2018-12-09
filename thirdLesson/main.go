package main

import "fmt"

const usixteen float64 = 65535
const kmh_multiple float64 = 1.60934

type car struct {
	gasPedal       uint16
	breakPedal     uint16
	streeringWheel int16
	topSpeedKM     float64
}

func (c *car) kmh() float64 {
	return float64(c.gasPedal) * (c.topSpeedKM / usixteen)
}

func (c *car) setStreeringWheel(num int16) {
	c.streeringWheel = num
}

func (c *car) mph() float64 {
	return float64(c.gasPedal) * (c.topSpeedKM / usixteen / kmh_multiple)
}

func main() {
	aCar := car{gasPedal: 22341,
		breakPedal:     0,
		streeringWheel: 20,
		topSpeedKM:     225.0}

	fmt.Println("a car")
	fmt.Println(aCar)
	fmt.Println(aCar.kmh())

	bCar := car{gasPedal: 65000,
		breakPedal:     0,
		streeringWheel: 20,
		topSpeedKM:     225.0}

	fmt.Println("b car")
	fmt.Println(bCar)
	fmt.Println(bCar.kmh())
	fmt.Println(bCar.mph())
	fmt.Println(bCar.streeringWheel)
	bCar.setStreeringWheel(80)
	fmt.Println(bCar.streeringWheel)
}
