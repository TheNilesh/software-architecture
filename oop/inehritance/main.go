package main

import "fmt"

func main() {
	s := NewTwiceIncrementer(8)
	fmt.Println(s.GetCount())
	s.DecrementCount()
	fmt.Println(s.GetCount())
	s.IncrementCount()
	s.IncrementCount()
	fmt.Println(s.GetCount())
}

type Incrementer interface {
	GetCount() int
	IncrementCount()
	DecrementCount()
}

// incrementer is a parent class
type incrementer struct {
	count int
}

func NewIncrementer(initialCount int) Incrementer {
	return &incrementer{
		count: initialCount,
	}
}

func (s *incrementer) GetCount() int {
	return s.count
}

func (s *incrementer) IncrementCount() {
	s.count++
}

func (s *incrementer) DecrementCount() {
	s.count--
}

func NewTwiceIncrementer(initialCount int) Incrementer {
	return &twiceIncrementer{
		incrementer: &incrementer{
			count: initialCount,
		},
	}
}

// twiceIncrementer is a child struct
type twiceIncrementer struct {
	*incrementer // this is composition
}

func (s *twiceIncrementer) GetCount() int {
	return s.incrementer.GetCount()
}

func (s *twiceIncrementer) IncrementCount() {
	s.incrementer.IncrementCount()
	s.incrementer.IncrementCount()
}

func (s *twiceIncrementer) DecrementCount() {
	s.incrementer.DecrementCount()
	s.incrementer.DecrementCount()
}
