package Auxiliary

import (
	"fmt"
	"log"
)

type Stack[T any] struct {
	underLyingArray []T
}

func (s *Stack[T]) Push(item T) {
	s.underLyingArray = append(s.underLyingArray, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		return *new(T), fmt.Errorf("empty Array")
	}
	lastItem := s.underLyingArray[s.Length()-1]
	s.underLyingArray = s.underLyingArray[:s.Length()-1]
	return lastItem, nil
}

func (s *Stack[T]) Peck() T {
	if s.IsEmpty() {
		log.Fatalln("empty stack!")
	}
	return s.underLyingArray[s.Length()-1]
}

func (s *Stack[T]) IsEmpty() bool {
	return s.Length() == 0
}

func (s *Stack[T]) Length() int {
	return len(s.underLyingArray)
}
