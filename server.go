package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

type Server struct {
	Ip   string
	Port int

	OnlineMap map[string]*User
	mapLock   sync.RWMutex

	Message chan string
}

func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:   ip,
		Port: port,

		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
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

	go this.ListenMessager()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("[ERROR] Listener accept err: ", err)
			continue
		}

		// do handler
		go this.Handler(conn)
	}
}

func (this *Server) Handler(conn net.Conn) {
	fmt.Println("[INFO] Connection established: ",
		conn.RemoteAddr())

	user := NewUser(conn, this)
	user.Online()

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("[ERROR] Conn Read err: ", err)
				return
			}

			// remove the '\n' at the end of msg
			msg := string(buf[:n-1])

			// broadcast
			user.DoMessage(msg)
		}
	}()

	select {}
}

func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "](" +
		user.Name + "): " + msg

	this.Message <- sendMsg
}

func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message

		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}

		this.mapLock.Unlock()
	}
}
