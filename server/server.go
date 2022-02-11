package server

import (
	"fmt"
	"net"
)

const (
	CONSUMER = iota
	SENDER
)

type Server struct {
	queues map[string]*queue
}

func NewServer(Name string) *Server {
	return &Server{map[string]*queue{}}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Printf("unable to start server: %s", err.Error())
		return
	}
	fmt.Println("starting server on port 8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("failed to accept connection: %s", err.Error())
			continue
		}
		fmt.Println("new connection")

		Ctype := make([]byte, 1)
		conn.Read(Ctype)
		Qsize := make([]byte, 1)
		conn.Read(Qsize)
		if Qsize[0] == byte(0) {
			fmt.Println("bad queue name size")
			conn.Close()
		}
		QName := make([]byte, uint8(Qsize[0]))
		conn.Read(QName)
		fmt.Println(Ctype)
		fmt.Println(Qsize)
		fmt.Println(QName)
		fmt.Println(string(QName))

		// when receiving a connection, checking if its calling a existing queue and creates if needed
		_, ok := s.queues[string(QName)]
		fmt.Println(ok)
		if !ok {
			s.queues[string(QName)] = newQueue()
			go s.queues[string(QName)].start()
		}
		fmt.Println(s.queues)

		if Ctype[0] == CONSUMER {
			c := newConsumer(conn)
			s.queues[string(QName)].consumers = append(s.queues[string(QName)].consumers, c)
		} else if Ctype[0] == SENDER {
			ch := s.queues[string(QName)].ch
			s := newSender(conn, ch)
			go s.listen()

		}
	}
}
