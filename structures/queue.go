package structures

import "fmt"

type Queue[T any] struct {
	arr []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{arr: []T{}}
}

func (q Queue[T]) String() string {
	return fmt.Sprintf("%v", q.arr)
}

func (q *Queue[T]) Enqueue(val T) {
	q.arr = append(q.arr, val)
}

func (q *Queue[T]) Dequeue() (out T, err bool) {
	if q.IsEmpty() {
		err = true
		return
	}

	out = q.arr[0]
	err = false
	q.arr = q.arr[1:]
	return
}
func (q Queue[T]) Size() int {
	return len(q.arr)
}

func (q Queue[T]) IsEmpty() bool {
	return len(q.arr) == 0
}
