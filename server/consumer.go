package server

import "net"

type consumer struct {
	conn net.Conn
}

func newConsumer(conn net.Conn) *consumer {
	return &consumer{conn}
}

func (c *consumer) sendMessage(data []byte) {
	c.conn.Write(data)
}
