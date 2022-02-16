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
		buffer:    [][]byte{},
		consumers: []*consumer{},
	}
}

func (q *queue) start() {
	for data := range q.ch {
		for index, consumer := range q.consumers {
			if consumer.available {
				err := consumer.sendMessage(data)
				if err != nil {
					fmt.Println("consumer disconected")
					consumer.conn.Close()
					q.consumers = append(q.consumers[:index], q.consumers[index+1:]...)
					continue
				} else {
					break
				}
			}
		}
		fmt.Println(q.consumers)
		// handel no available consumer / no consumers
	}
}
