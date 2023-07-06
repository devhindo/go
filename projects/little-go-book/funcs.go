package main

func log(message string) {
	println(message)
}

func add(a, b int) int {
	return a+b
}

func power(name string) (int,bool) {
	return 5,true
}

func funcs() {
	_, exists := power("ahmed")
	if exists == false {

	}
}