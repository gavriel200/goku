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
		buffer:    [][]byte{},    // make([][]byte, 10),
		consumers: []*consumer{}, // make([]*consumer, 0),
	}
}

func (q *queue) start() {
	fmt.Println("started consuming on queue")
	for data := range q.ch {
		fmt.Println(data)
		fmt.Println(q.consumers)
		err := q.consumers[0].sendMessage(data)
		if err != nil {
			fmt.Println("consumer disconected")
			q.consumers[0].conn.Close()
			q.consumers = []*consumer{}
		}
	}
}
