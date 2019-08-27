package main

import (
	"fmt"
	"time"
)

func main() {
	pub := &NewsPublisher{}

	s1 := &S1{}
	s2 := &S2{}

	sv := NewServer()
	sv.AddSubsciber("TOPIC-0", s1)
	sv.AddSubsciber("TOPIC-0", s2)
	sv.AddSubsciber("TOPIC-1", s2)

	time.Sleep(1000)
	for i := 0; i < 10; i++ {
		pub.Publish(fmt.Sprintf("TOPIC-%d", i&1), i)
		time.Sleep(500)
	}
}
