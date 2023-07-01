package main

import (
	"fmt"
)

type About struct {
	name        string
	version     string
	description string
	updateTime  string
}

func main() {
	about := About{
		name:        "SeimonCore",
		version:     "V0.0.1 Genesis Prototype",
		description: "The Genesis of SeimonCore",
		updateTime:  "2023-7-2T2:41:00",
	}
	fmt.Println(about)
}
