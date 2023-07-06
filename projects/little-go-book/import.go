package main

import (
	"fmt"
	"os"
	"reflect"
)

func sub_main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("it's over", os.Args[1], "num of args: ", len(os.Args))
	fmt.Println(reflect.TypeOf(os.Args))
	fmt.Println(os.Args[0])
}