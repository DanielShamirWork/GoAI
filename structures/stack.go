package structures

import "fmt"

type Stack[T any] struct {
	arr []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{arr: []T{}}
}

func (s Stack[T]) String() string {
	return fmt.Sprintf("%v", s.arr)
}

func (s *Stack[T]) Push(val T) {
	s.arr = append(s.arr, val)
}

func (s *Stack[T]) Peek() (out T, err bool) {
	if s.IsEmpty() {
		err = true
		return
	}

	out = s.arr[len(s.arr)-1]
	err = false
	return
}
func (s *Stack[T]) Pop() (out T, err bool) {
	out, err = s.Peek()
	if !err {
		s.arr = s.arr[:len(s.arr)-1]
	}
	return
}

func (s *Stack[T]) Size() int {
	return len(s.arr)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.arr) == 0
}
