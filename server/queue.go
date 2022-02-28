package server

import (
	"fmt"
	"sync"
)

type queue struct {
	ch            chan []byte
	buffer        [][]byte
	bufferLock    sync.RWMutex
	consumers     []*consumer
	consumersLock sync.RWMutex
	availableAck  chan struct{}
}

func newQueue() *queue {
	return &queue{
		ch:            make(chan []byte),
		buffer:        [][]byte{},
		bufferLock:    sync.RWMutex{},
		consumers:     []*consumer{},
		consumersLock: sync.RWMutex{},
		availableAck:  make(chan struct{}),
	}
}

func (q *queue) start() {
	go q.listenForChannel()
	go q.listenForAvailableAck()
}

func (q *queue) listenForChannel() {
	for data := range q.ch {
		go q.handleNewData(data)
	}
}

func (q *queue) listenForAvailableAck() {
	for range q.availableAck {
		if len(q.buffer) == 0 {
			continue
		}
		data := q.getFromBuffer()
		go q.handleNewData(data)
	}
}

func (q *queue) handleNewData(data []byte) {
	for index, consumer := range q.consumers {
		if consumer.available {
			err := consumer.sendMessage(data)
			if err != nil {
				fmt.Println("consumer disconected")
				q.removeConsumer(consumer, index)
				continue
			} else {
				q.availableAck <- struct{}{}
				return
			}
		}
	}
	// handle no available consumer
	q.addToBuffer(data)
	fmt.Println(q.buffer)
}

func (q *queue) addConsumer(c *consumer) {
	q.consumers = append(q.consumers, c)
	q.availableAck <- struct{}{}

}

func (q *queue) removeConsumer(c *consumer, index int) {
	c.conn.Close()
	q.consumers = append(q.consumers[:index], q.consumers[index+1:]...)
}

func (q *queue) addToBuffer(data []byte) {
	q.buffer = append(q.buffer, data)
}

func (q *queue) getFromBuffer() []byte {
	data := q.buffer[0]
	q.buffer = q.buffer[1:]
	return data
}
