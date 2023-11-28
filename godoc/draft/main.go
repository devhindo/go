package main

import (
	"fmt"
)

var (
	a int
	b int 
)

func main() {
	asign()
	reassign()
	fmt.Println(a, b)
}

func asign() {
	a = 4
	b = 5
}

func reassign() {
	a = 3
	b = 6
}