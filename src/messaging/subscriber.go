package main

type Subscriber interface {
	Listen(topic string) error
	StopListen(topic string) error
	Receive()
	Offset() int
}

//-----------------------------
type S1 struct {
	offset int
}

func (s *S1) Listen(topic string) error {

	return nil
}
func (s *S1) StopListen(topic string) error {
	return nil
}
func (s *S1) Receive() {
	s.offset++
	return
}
func (s *S1) Offset() int {
	return s.offset
}

//-----------------------------

type S2 struct {
	offset int
}

func (s *S2) Listen(topic string) error {

	return nil
}
func (s *S2) StopListen(topic string) error {
	return nil
}
func (s *S2) Receive() {
	return

}
func (s *S2) Offset() int {
	return s.offset
}
