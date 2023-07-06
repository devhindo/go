package main

import "fmt"

func variables() {
	var power int // equals 0 by default
	power = 9000
	var energy int = power * 2
	dis := 3
	fmt.Println(energy, dis)
	fmt.Printf("it's over %d\n", power)
	a := returnHund()
	fmt.Println(a)
	x,y,z := 1,"ahmed",5.5
	println(x,y,z)
}

func returnHund() int {
	return 1000
}