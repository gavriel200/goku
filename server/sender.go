package server

import "net"

type sender struct {
	conn net.Conn
	ch   chan []byte
}

func newSender(conn net.Conn, ch chan []byte) *sender {
	return &sender{conn, ch}
}

func (s *sender) listen() {
	for {
		data := make([]byte, 1)
		s.conn.Read(data)
		s.ch <- data
	}
}
