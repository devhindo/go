package engine

import "fmt"

const (
	appID   int    = 1
	appName string = "gopher"
)

func Config() {

	fmt.Println(appID,appName)

}