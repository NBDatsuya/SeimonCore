package main

import (
	"fmt"
	"net"
)

/*
type About struct {
	name        string
	version     string
	description string
	updateTime  string
	since string
}*/

func main() {
	/*
		about := About{
			name:        "SeimonCore",
			version:     "V0.0.1 Genesis Prototype",
			description: "The Genesis of SeimonCore",
			updateTime:  "2023-7-6T2:22:00",
			since: "2023-7-2T2:41:00",
		}*/
	server := NewServer("127.0.0.1", 2121)
	server.Start()
}
