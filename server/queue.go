package server

import "fmt"

type queue struct {
	ch        chan []byte
	buffer    [][]byte
	consumers []*consumer
}

func newQueue() *queue {
	return &queue{
		ch:        make(chan []byte),
		buffer:    make([][]byte, 10),
		consumers: make([]*consumer, 10),
	}
}

func (q *queue) start() {
	fmt.Println("started consuming on queue")
	for data := range q.ch {
		fmt.Println(data)
	}
}
