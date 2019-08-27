package main

import (
	"fmt"
	"sync"
)

type Topic struct {
	Subscribers map[Subscriber]bool
	queue       []interface{}
	total       int
}

func newTopic() *Topic {
	t := new(Topic)
	t.Subscribers = make(map[Subscriber]bool)
	t.queue = make([]interface{}, 0)

	return t
}

type Server struct {
	topics map[string]*Topic
	lock   *sync.Mutex
}

var once sync.Once
var sv Server

func NewServer() Server {
	once.Do(func() { // atomic, does not allow repeating
		sv.topics = make(map[string]*Topic)
		sv.lock = new(sync.Mutex)
	})

	return sv
}

func (s *Server) Write(topic string, msg interface{}) error {
	s.lock.Lock()
	s.lock.Unlock()

	if t, ok := s.topics[topic]; ok {
		t.queue = append(t.queue, msg)
	} else {
		t := newTopic()
		s.topics[topic] = t
	}

	go s.StartConsumming(topic)

	return nil
}

func (s *Server) AddSubsciber(topic string, sub Subscriber) error {
	s.lock.Lock()
	s.lock.Unlock()

	if t, ok := s.topics[topic]; ok {
		t.Subscribers[sub] = true
	} else {
		t := newTopic()
		s.topics[topic] = t
		t.Subscribers[sub] = true
	}

	return nil
}

func (s *Server) StartConsumming(topic string) {
	t, _ := s.topics[topic]

	for sub, _ := range t.Subscribers {
		if sub.Offset() < len(t.queue) {
			fmt.Printf("Consumming:  %v  %v\n", topic, t.queue[sub.Offset()])
			sub.Receive()
		}
	}
}
