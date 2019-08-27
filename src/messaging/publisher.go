package main

import "fmt"

type publisher interface {
	Publish(topic string, msg interface{}) error
}

type NewsPublisher struct {
}

func (p *NewsPublisher) Publish(topic string, msg interface{}) error {
	sv := NewServer()

	fmt.Printf("Writing:  %v  %v\n", topic, msg)
	sv.Write(topic, msg)

	return nil
}
