package server

import (
	"fmt"
	"net"
)

const ping byte = 0xFF

type consumer struct {
	conn      net.Conn
	available bool
}

func newConsumer(conn net.Conn) *consumer {
	available := true
	return &consumer{conn, available}
}

func (c *consumer) sendMessage(data []byte) error {
	// using ping because the write returns an error on the second Write since the disconnect
	c.conn.Write([]byte{ping})
	_, conError := c.conn.Write(data)
	if conError != nil {
		return conError
	}
	fmt.Println("sent: ", data)
	return nil
}
