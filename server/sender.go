package server

import (
	"fmt"
	"net"
)

type sender struct {
	conn net.Conn
	ch   chan []byte
}

func newSender(conn net.Conn, ch chan []byte) *sender {
	return &sender{conn, ch}
}

func (s *sender) listen() {
	defer s.conn.Close()
	for {
		data := make([]byte, 1)
		_, err := s.conn.Read(data)
		if err != nil {
			fmt.Println("sender disconected")
			break
		}
		fmt.Println("received: ", data)
		s.ch <- data
	}
}
