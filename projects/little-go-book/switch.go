package main

import(
	"time"
	"fmt"
)

func swit() {
	today := time.Now().Weekday()
	switch today {
	case time.Sunday:
		fmt.Println("Sunday")
	case time.Monday:
		fmt.Println("Monday")
	}
}