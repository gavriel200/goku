package server

import (
	"fmt"
	"net"
)

type consumer struct {
	conn net.Conn
}

func newConsumer(conn net.Conn) *consumer {
	return &consumer{conn}
}

func (c *consumer) sendMessage(data []byte) error {
	val, err := c.conn.Write(data) // when consumer disconects error accurs only after second message
	fmt.Println(val)
	if err != nil {
		return err
	}
	return nil
}
