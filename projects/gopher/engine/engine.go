package engine

import (
	"fmt"
	"gopher/sub"
)

func Run() {
	Config()
	fmt.Println("engine running...")
	sub.PrintSub()
}