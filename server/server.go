package server

import (
	"fmt"
	"net"
)

const (
	CONSUMER uint8 = iota
	SENDER
)

type Server struct {
	queues map[string]*queue
}

func NewServer() *Server {
	return &Server{map[string]*queue{}}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("unable to start server: ", err)
		return
	}
	fmt.Println("starting server on port 8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("failed to accept connection: ", err)
			continue
		}
		go s.handleNewConnection(conn)
	}
}

func (s *Server) handleNewConnection(conn net.Conn) {
	connData := make([]byte, 2)
	conn.Read(connData)

	connType := connData[0]
	queueSize := connData[1]
	if queueSize == byte(0) {
		fmt.Println("bad queue name size")
		conn.Close()
	}
	queueNameBytes := make([]byte, uint8(queueSize))
	conn.Read(queueNameBytes)
	queueName := string(queueNameBytes)

	// check if queue already exists
	_, ok := s.queues[queueName]
	if !ok {
		s.queues[queueName] = newQueue()
		s.queues[queueName].start()
	}

	if connType == CONSUMER {
		fmt.Println("new connection CONSUMER")
		s.queues[queueName].addConsumer(newConsumer(conn))
	} else if connType == SENDER {
		fmt.Println("new connection SENDER")
		ch := s.queues[queueName].ch
		s := newSender(conn, ch)
		go s.listen()
	}
}
