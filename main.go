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

func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:   ip,
		Port: port,
	}
}

func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp",
		fmt.Sprintf("%s:%d", this.Ip, this.Port))

	if err != nil {
		fmt.Println("[ERROR] net.Listen err: ", err)
		return
	}

	// close listen socket
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[ERROR] listener accept err: ", err)
			continue
		}

		// do handler
		go this.Handler(conn)
	}
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Println("[INFO] connection established: ",
		conn.RemoteAddr())
}
